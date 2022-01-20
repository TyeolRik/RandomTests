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
	seed, _ := strconv.ParseInt(_seed, 10, 64)
	buf := make([]byte, BUFFERSIZE)

	acorn := rngset.NewACORN(seed)
	for {
		for i := 0; i < (BUFFERSIZE / 4); i++ {
			binary.LittleEndian.PutUint32(buf[i*4:], uint32(acorn.NextInt64()))
		}
		binary.Write(os.Stdout, binary.LittleEndian, buf)
	}
}
