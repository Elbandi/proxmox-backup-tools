package dump

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/elbandi/proxmox-backup-tools/common"
	"github.com/urfave/cli/v2"
	"io"
	"os"
	"sort"
	"strings"
)

func readHashFromFidx(filename string) ([]*chainhash.Hash, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Skip the first 4096 bytes
	if _, err := file.Seek(4096, 0); err != nil {
		return nil, err
	}

	var data []*chainhash.Hash
	buffer := make([]byte, 32)

	// Read the rest of the file in 32-byte pieces and store them in an array
	for {
		bytesRead, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if bytesRead < 32 {
			break
		}
		h, err := chainhash.NewHash(buffer)
		if err != nil {
			return nil, err
		}
		data = append(data, h)
	}

	return data, nil
}

func readHashFromDidx(filename string) ([]*chainhash.Hash, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Skip the first 4096 bytes
	if _, err := file.Seek(4096, 0); err != nil {
		return nil, err
	}

	var data []*chainhash.Hash
	buffer := make([]byte, 40) // 8+32

	// Read the rest of the file in 40-byte pieces and store them in an array
	for {
		bytesRead, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if bytesRead < 40 {
			break
		}
		h, err := chainhash.NewHash(buffer[8:])
		if err != nil {
			return nil, err
		}
		data = append(data, h)
	}

	return data, nil
}

func readHashFromFile(filename string) ([]*chainhash.Hash, error) {
	if strings.HasSuffix(filename, "fidx") {
		return readHashFromFidx(filename)
	}
	return readHashFromDidx(filename)
}

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
		hashtable, err := readHashFromFile(k)
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
