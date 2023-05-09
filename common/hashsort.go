package common

import (
	"bytes"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

type HashSorter []*chainhash.Hash

func (s HashSorter) Len() int { return len(s) }

func (s HashSorter) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s HashSorter) Less(i, j int) bool {
	return bytes.Compare(s[i][:], s[j][:]) < 0
}
