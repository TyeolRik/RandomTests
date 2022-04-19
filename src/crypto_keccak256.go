package main

import (
	"crypto/rand"
	"encoding/binary"
	"hash"
	"os"

	"golang.org/x/crypto/sha3"
)

func XOR_32bytes_arr(a []byte, b []byte) []byte {
	var ret []byte = make([]byte, 32)
	for i := 0; i < 32; i++ {
		ret[i] = a[i] ^ b[i]
	}
	return ret
}

func XOR_32bytes_6arr(a []byte, b []byte, c []byte, d []byte, e []byte, f []byte) []byte {
	var ret []byte = make([]byte, 32)
	for i := 0; i < 32; i++ {
		ret[i] = a[i] ^ b[i] ^ c[i] ^ d[i] ^ e[i] ^ f[i]
	}
	return ret
}

var keccak256 hash.Hash

func main() {
	for {
		input := make([]byte, 64*6*1000) // 6블록
		rand.Read(input)
		afterCollectBlock := make([]byte, 32*6*1000)
		for i := range afterCollectBlock {
			afterCollectBlock[i] = input[i] ^ input[i+32]
		}

		// Need Keccak?

		merge6Blocks := make([]byte, 0, 32*1000)
		for i := 0; i < 1000; i++ {
			merge6Blocks = append(merge6Blocks, XOR_32bytes_6arr(afterCollectBlock[192*i:192*i+32], afterCollectBlock[192*i+32:192*i+64], afterCollectBlock[192*i+64:192*i+96], afterCollectBlock[192*i+96:192*i+128], afterCollectBlock[192*i+128:192*i+160], afterCollectBlock[192*i+160:192*i+192])...)
		}
		for i := 0; i < 1000; i++ {
			keccak256 = sha3.NewLegacyKeccak256()
			keccak256.Write(merge6Blocks[i*32 : i*32+32])
			copy(merge6Blocks[i*32:i*32+32], keccak256.Sum(nil))
		}
		binary.Write(os.Stdout, binary.LittleEndian, merge6Blocks)
	}

}
