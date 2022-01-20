package main

import (
	"encoding/binary"
	"os"
	"strconv"

	mathrand "math/rand"
)

const BUFFERSIZE = 10000

func main() {
	_seed := os.Args[1]
	seed, _ := strconv.ParseInt(_seed, 10, 64)
	mathrand.Seed(seed)
	buf := make([]byte, BUFFERSIZE)
	var myRand uint64
	for {
		for i := 0; i < (BUFFERSIZE / 8); i++ {
			myRand = mathrand.Uint64()
			binary.LittleEndian.PutUint64(buf[i*8:], myRand)
		}
		binary.Write(os.Stdout, binary.LittleEndian, buf)
	}
}
