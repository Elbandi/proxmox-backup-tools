package checksum

import (
	"fmt"
	"github.com/elbandi/proxmox-backup-tools/common"
	"github.com/urfave/cli"
	"os"
)

var ChecksumCommand = cli.Command{
	Name:  "checksum",
	Usage: "Calculate backup images checksum",
	Description: `
	`,
	Flags: []cli.Flag{
		common.SrcRepo,
		common.SrcNamespace,
		common.SrcFingerprint,
		common.SrcPassword,
		common.SrcKeyFile,
		common.SrcKeyPassword,
		common.BackupId,
		common.BackupTime,
	},
	Action: cmdChecksum,
}

func cmdChecksum(ctx *cli.Context) error {
	// ctxc := getContext()
	backupFiles := ctx.String("Backup")
	fmt.Println(backupFiles)
	file, err := os.Open(backupFiles)
	common.CheckErr(err, "open file")
	defer common.DeferClose(file, "close file")
	return nil
}
