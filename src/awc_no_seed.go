package main

import (
	"encoding/binary"
	"os"

	"github.com/tyeolrik/rngset"
)

const BUFFERSIZE = 10000

func main() {
	buf := make([]byte, BUFFERSIZE)
	awc := rngset.NewAWC_Recommend()
	for {
		for i := 0; i < (BUFFERSIZE / 4); i++ {
			binary.LittleEndian.PutUint32(buf[i*4:], uint32(awc.NextInt64()))
		}
		binary.Write(os.Stdout, binary.LittleEndian, buf)
	}
}
