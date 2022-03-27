package main

import (
	"encoding/binary"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/crypto"
)

const BUFFERSIZE = 32000

var temp []byte

func main() {
	seed, _ := strconv.ParseUint(os.Args[1], 10, 64)
	byteArray := make([]byte, 8)
	binary.LittleEndian.PutUint64(byteArray, seed)

	buf := make([]byte, BUFFERSIZE)
	temp = crypto.Keccak256(byteArray)
	for {
		for i := 0; i < (BUFFERSIZE / 32); i++ {
			temp = crypto.Keccak256(temp)
			copy(buf[i*32:], temp)
		}
		binary.Write(os.Stdout, binary.LittleEndian, buf)
	}
}
