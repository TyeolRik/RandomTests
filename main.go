package main

import (
	"encoding/binary"
	"os"
	"strconv"

	"github.com/tyeolrik/rngset"
)

const BUFFERSIZE = 10000

func main() {
	seed, _ := strconv.ParseUint(os.Args[1], 10, 64)
	buf := make([]byte, BUFFERSIZE)
	mt19937_64 := rngset.NewMT19937_64(seed)
	for {
		for i := 0; i < (BUFFERSIZE / 8); i++ {
			binary.LittleEndian.PutUint64(buf[i*8:], mt19937_64.NextUint64())
		}
		binary.Write(os.Stdout, binary.LittleEndian, buf)
	}
}
