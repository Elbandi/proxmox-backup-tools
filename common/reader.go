package common

import (
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"io"
	"os"
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

func ReadHashFromFile(filename string) ([]*chainhash.Hash, error) {
	if strings.HasSuffix(filename, "fidx") {
		return readHashFromFidx(filename)
	}
	return readHashFromDidx(filename)
}
