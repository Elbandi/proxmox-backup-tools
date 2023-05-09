package main

import (
	"github.com/elbandi/proxmox-backup-tools/cmd/checksum"
	"github.com/elbandi/proxmox-backup-tools/cmd/dump"
	"github.com/elbandi/proxmox-backup-tools/cmd/reencrypt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "print-version",
		Aliases: []string{"V"},
		Usage:   "print only the version",
	}

	app := &cli.App{
		Name:        "proxmox-backup-tools",
		Version:     "v1.0",
		Description: "Proxmox Backup tools",
		Commands: []*cli.Command{
			&checksum.ChecksumCommand,
			&dump.DumpHashCommand,
			&reencrypt.ReencryptCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
