package main

import (
	"crypto/rand"
	"encoding/binary"
	"os"
)

const BUFFERSIZE = 10000

func main() {
	buf := make([]byte, BUFFERSIZE)
	for {
		rand.Read(buf)
		binary.Write(os.Stdout, binary.LittleEndian, buf)
	}
}
