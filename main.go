package main

import (
	"github.com/elbandi/proxmox-backup-tools/cmd/checksum"
	"github.com/elbandi/proxmox-backup-tools/cmd/dump"
	"github.com/elbandi/proxmox-backup-tools/cmd/reencrypt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "lncli"
	app.Version = "1.0"
	app.Usage = "Proxmox Backup tools"
	app.Commands = []cli.Command{
		checksum.ChecksumCommand,
		dump.DumpHashCommand,
		reencrypt.ReencryptCommand,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
