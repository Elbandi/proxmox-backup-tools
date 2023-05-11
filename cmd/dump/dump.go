package dump

import (
	"encoding/hex"
	"fmt"
	"github.com/elbandi/proxmox-backup-tools/common"
	"github.com/urfave/cli/v2"
	"sort"
)

var DumpHashCommand = cli.Command{
	Name:  "dump-hash",
	Usage: "Dump hash",
	Description: `lofasz
	`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:     "with-dir",
			Required: false,
			Usage:    "Print hash with dir",
		},
	},
	Action: cmdDumpHash,
}

func cmdDumpHash(ctx *cli.Context) error {
	var outputFormat string
	if ctx.Bool("with-dir") {
		outputFormat = "%02[1]s/%[2]s\n"
	} else {
		outputFormat = "%[2]s\n"
	}
	args := ctx.Args()
	for i, k := range args.Slice() {
		hashtable, err := common.ReadHashFromFile(k)
		common.CheckErr(err, "hash read on %s", k)
		sort.Sort(common.HashSorter(hashtable))
		for _, h := range hashtable {
			hashStr := hex.EncodeToString(h[:])
			fmt.Printf(outputFormat, hashStr[:2], hashStr)
		}
		_ = i
		//		hashes[i] = hashtable
	}
	return nil
}
