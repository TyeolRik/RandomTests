package main

import (
	"encoding/binary"
	"os"
	"strconv"

	"github.com/tyeolrik/rngset"
)

const BUFFERSIZE = 10000

func main() {
	_seed := os.Args[1]
	// seed, _ := strconv.ParseInt(_seed, 10, 64)
	seed, _ := strconv.ParseUint(_seed, 10, 64)
	buf := make([]byte, BUFFERSIZE)

	kiss := rngset.NewKISS(seed, seed)
	for {
		for i := 0; i < (BUFFERSIZE / 8); i++ {
			binary.LittleEndian.PutUint64(buf[i*8:], kiss.NextUInt64())
		}
		binary.Write(os.Stdout, binary.LittleEndian, buf)
	}
}
