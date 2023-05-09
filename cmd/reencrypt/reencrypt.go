package reencrypt

import (
	"fmt"
	"github.com/elbandi/proxmox-backup-tools/common"
	"github.com/urfave/cli/v2"
	"os"
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
		&common.BackupId,
		&common.BackupTime,
	},
	Action: cmdReencrypt,
}

func cmdReencrypt(ctx *cli.Context) error {
	// ctxc := getContext()
	backupFiles := ctx.String("Backup")
	fmt.Println(backupFiles)
	file, err := os.Open(backupFiles)
	common.CheckErr(err, "open file")
	defer common.DeferClose(file, "close file")
	return nil
}
