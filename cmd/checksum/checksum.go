package checksum

import (
	"crypto/sha256"
	"errors"
	"fmt"
	pbs "github.com/Elbandi/go-proxmox-backup-client"
	"github.com/elbandi/proxmox-backup-tools/common"
	"github.com/urfave/cli/v2"
	"io"
	"strings"
)

var ChecksumCommand = cli.Command{
	Name:  "checksum",
	Usage: "Calculate backup images checksum",
	Description: `
	`,
	Flags: []cli.Flag{
		&common.SrcRepo,
		&common.SrcNamespace,
		&common.SrcFingerprint,
		&common.SrcPassword,
		&common.SrcKeyFile,
		&common.SrcKeyPassword,
		&common.BackupType,
		&common.BackupId,
		&common.BackupTime,
	},
	Action: cmdChecksum,
}

func cmdChecksum(ctx *cli.Context) error {
	srcClient, err := pbs.NewRestore(
		ctx.String(common.SrcRepo.Name),
		ctx.String(common.SrcNamespace.Name),
		ctx.String(common.BackupType.Name),
		ctx.String(common.BackupId.Name),
		ctx.Generic(common.BackupTime.Name).(*common.ProxmoxDateTimeValue).Value(),
		ctx.String(common.SrcPassword.Name),
		ctx.String(common.SrcFingerprint.Name),
		ctx.String(common.SrcKeyFile.Name),
		ctx.String(common.SrcKeyPassword.Name),
	)

	common.CheckErr(err, "open restore")
	defer srcClient.Close()

	files, err := srcClient.ListFiles()
	common.CheckErr(err, "list manifest files")

	for _, file := range files {
		if strings.HasSuffix(file, ".conf.blob") || strings.HasSuffix(file, ".log.blob_") {
			buf := make([]byte, pbs.GetDefaultChunkSize())
			readSize, err := srcClient.RestoreBlob(file, buf)
			common.CheckErr(err, "restore blob %s", file)
			hash := sha256.Sum256(buf[:readSize])
			fmt.Printf("%s: %x\n", file, hash)
		} else if strings.HasSuffix(file, ".img.fidx") {
			hash := sha256.New()
			buf := make([]byte, pbs.GetDefaultChunkSize())
			srcImage, err := srcClient.OpenImage(file)
			common.CheckErr(err, "open image %s", file)
			imgSize, err := srcImage.Size()
			common.CheckErr(err, "get image %s size", file)
			reader := common.NewOffsetReadSeeker(srcImage, 0)
			bar := common.SetupProgressbar(imgSize, file)
			written, err := io.CopyBuffer(io.MultiWriter(hash, bar), reader, buf)
			common.CheckErr(err, "copy image %s", file)
			if uint64(reader.Position()) != imgSize {
				common.CheckErr(errors.New("not fully read"), "read image %s", file)
			}
			if imgSize != uint64(written) {
				common.CheckErr(errors.New("not fully written"), "read image %s", file)
			}
			fmt.Printf("%s: %x\n", file, hash.Sum(nil))
		}
	}
	return nil
}
