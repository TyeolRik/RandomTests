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

	zx81 := rngset.NewZX81(seed) // Maximum is 16 bit unsigned integer
	var myRand uint16
	for {
		for i := 0; i < (BUFFERSIZE / 2); i++ {
			myRand = uint16(zx81.Generate())
			binary.LittleEndian.PutUint16(buf[i*2:], myRand)
		}
		binary.Write(os.Stdout, binary.LittleEndian, buf)
	}
}
