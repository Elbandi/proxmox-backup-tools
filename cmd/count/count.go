package count

import (
	"fmt"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/elbandi/proxmox-backup-tools/common"
	"github.com/urfave/cli/v2"
	"golang.org/x/text/language"
	"slices"
	"sort"
)

var CountChunksCommand = cli.Command{
	Name:  "count-chunks",
	Usage: "Count backup chunks",
	Description: `Count backup chunks
	`,
	Flags: []cli.Flag{
		&cli.Uint64Flag{
			Name:     "output",
			Required: false,
			Value:    0,
			Usage:    "Differences output format",
		},
		&cli.GenericFlag{
			Name:     "locale",
			Required: false,
			Value:    &common.LanguageValue{Default: language.English},
			Usage:    "Locale",
		},
	},
	Action: cmdCountchunks,
}

func cmdCountchunks(ctx *cli.Context) error {
	args := ctx.Args()
	hashes := make([]*chainhash.Hash, 0)
	for i, k := range args.Slice() {
		h, err := common.ReadHashFromFile(k)
		common.CheckErr(err, "hash read on %s", k)
		hashes = append(hashes, h...)
		sort.Sort(common.HashSorter(hashes))
		slices.Compact(hashes)
		_ = i
	}

	fmt.Printf("Full hash count: %d\n", len(hashes))
	return nil
}
