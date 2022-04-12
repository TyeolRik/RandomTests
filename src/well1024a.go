package main

import (
	"encoding/binary"
	"os"

	"github.com/tyeolrik/rngset"
)

var temp []byte

func main() {
	temp = make([]byte, 320000)
	seeds := [32]uint32{}
	for i := range seeds {
		seeds[i] = uint32(i)
	}
	r := rngset.NewWELL1024a(seeds)
	for {
		for i := 0; i < 320000; i = i + 4 {
			binary.LittleEndian.PutUint32(temp[i:i+4], r.NextUint32())
		}
		binary.Write(os.Stdout, binary.LittleEndian, temp)
	}
}
