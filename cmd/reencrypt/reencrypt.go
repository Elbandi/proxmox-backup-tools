package reencrypt

import (
	"errors"
	"fmt"
	pbs "github.com/Elbandi/go-proxmox-backup-client"
	"github.com/elbandi/proxmox-backup-tools/common"
	"github.com/urfave/cli/v2"
	"io"
	"strings"
)

var ReencryptCommand = cli.Command{
	Name:  "reencrypt",
	Usage: "Re-encrypt backup",
	Description: `
	`,
	Flags: []cli.Flag{
		&common.SrcRepo,
		&common.SrcNamespace,
		&common.SrcFingerprint,
		&common.SrcPassword,
		&common.SrcKeyFile,
		&common.SrcKeyPassword,
		&common.DstRepo,
		&common.DstNamespace,
		&common.DstFingerprint,
		&common.DstPassword,
		&common.DstKeyFile,
		&common.DstKeyPassword,
		&common.BackupType,
		&common.BackupId,
		&common.BackupTime,
	},
	Action: cmdReencrypt,
}

func cmdReencrypt(ctx *cli.Context) error {
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

	dstClient, err := pbs.NewBackup(
		ctx.String(common.DstRepo.Name),
		ctx.String(common.DstNamespace.Name),
		ctx.String(common.BackupId.Name),
		ctx.Generic(common.BackupTime.Name).(*common.ProxmoxDateTimeValue).Value(),
		ctx.String(common.DstPassword.Name),
		ctx.String(common.DstFingerprint.Name),
		ctx.String(common.DstKeyFile.Name),
		ctx.String(common.DstKeyPassword.Name),
		true,
	)
	common.CheckErr(err, "create backup")
	defer dstClient.Close()

	for _, file := range files {
		fmt.Println(file)
		if strings.HasSuffix(file, ".conf.blob") || strings.HasSuffix(file, ".log.blob_") {
			buf := make([]byte, pbs.GetDefaultChunkSize())
			readSize, err := srcClient.RestoreBlob(file, buf)
			common.CheckErr(err, "restore blob %s", file)
			err = dstClient.AddConfig(strings.TrimSuffix(file, ".blob"), buf[:readSize])
			common.CheckErr(err, "write blob %s", file)
			//if writeSize != readSize {
			//	common.CheckErr(errors.New("not fully written"), "read image %s", file)
			//}
		} else if strings.HasSuffix(file, ".img.fidx") {
			func() { // csak hogy menjen a defer
				buf := make([]byte, pbs.GetDefaultChunkSize())
				srcImage, err := srcClient.OpenImage(file)
				common.CheckErr(err, "open image %s", file)
				//defer DeferClose(srcImage, "close image")
				imgSize, err := srcImage.Size()
				common.CheckErr(err, "get image %s size", file)
				dstImage, err := dstClient.RegisterImage(strings.TrimSuffix(file, ".img.fidx"), imgSize)
				common.CheckErr(err, "register image %s", file)
				defer common.DeferClose(dstImage, "close image", file)
				reader := common.NewOffsetReadSeeker(srcImage, 0)
				writer := common.NewOffsetWriter(dstImage, 0)
				bar := common.SetupProgressbar(imgSize, file)
				written, err := io.CopyBuffer(io.MultiWriter(writer, bar), reader, buf)
				common.CheckErr(err, "copy image %s", file)
				if uint64(reader.Position()) != imgSize {
					common.CheckErr(errors.New("not fully read"), "read image %s", file)
				}
				if imgSize != uint64(written) {
					common.CheckErr(errors.New("not fully written"), "read image %s", file)
				}
			}()
		}
	}
	err = dstClient.Finish()
	common.CheckErr(err, "finish backup")
	return nil
}
