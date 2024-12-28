package compare

import (
	"bytes"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/elbandi/proxmox-backup-tools/common"
	"github.com/urfave/cli/v2"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"path/filepath"
	"sort"
	"strings"
)

var CompareCommand = cli.Command{
	Name:  "compare",
	Usage: "Compare backups",
	Description: `Compare backups
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
	Action: cmdCompareBackups,
}

func countCommonElements(arr1, arr2 []*chainhash.Hash) int {
	var count int
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		cmp := bytes.Compare(arr1[i][:], arr2[j][:])
		if cmp < 0 {
			i++
		} else if cmp > 0 {
			j++
		} else {
			count++
			i++
			j++
		}
	}

	return count
}

func cmdCompareBackups(ctx *cli.Context) error {
	files := make([]string, ctx.NArg())
	hashes := make([][]*chainhash.Hash, ctx.NArg())
	args := ctx.Args()
	for i, k := range args.Slice() {
		h, err := common.ReadHashFromFile(k)
		common.CheckErr(err, "hash read on %s", k)
		sort.Sort(common.HashSorter(h))
		files[i] = strings.Split(filepath.Base(filepath.Dir(k)), "T")[0]
		hashes[i] = h
	}

	outputFormat := ctx.Uint64("output")
	p := message.NewPrinter(ctx.Generic("locale").(*common.LanguageValue).Value())
	fmt.Printf("X;%s\n", strings.Join(files, ";"))
	for i, f := range files {
		fmt.Printf("%s;", f)
		h1 := hashes[i]
		for j, h2 := range hashes {
			if j < i {
				p.Printf("0;")
				continue
			} else if j == i {
				switch outputFormat {
				case 2:
					p.Printf("%d;", len(h1))
				default:
					p.Printf("0;")
				}
				continue
			}
			count := countCommonElements(h1, h2)
			switch outputFormat {
			case 1:
				p.Printf("%d;", len(h1)+len(h2)-(count*2))
			case 2:
				p.Printf("%d;", count)
			default:
				p.Printf("%.2f;", 100-float64(count*2*100)/float64(len(h1)+len(h2)))
			}
		}
		fmt.Printf("\n")
	}
	return nil
}
