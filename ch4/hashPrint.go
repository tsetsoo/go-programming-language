package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	params := os.Args[1:]
	if len(params) == 1 {
		fmt.Println(sha256.Sum256([]byte(params[0])))
		fmt.Println(hashDifferentBytes(sha256.Sum256([]byte(params[0])), sha256.Sum256([]byte("X"))))
	} else if params[1] == "384" {
		fmt.Println(sha512.Sum384([]byte(params[0])))
	} else if params[1] == "512" {
		fmt.Println(sha512.Sum512([]byte(params[0])))
	}
}

func hashDifferentBytes(a, b [32]byte) byte {
	var result byte
	for i, ax := range a {
		result += popCountShift8(ax, b[i])
	}
	return result
}

func popCountShift8(x, y byte) byte {
	var result byte
	for i := 0; i < 8; i++ {
		a := int((x >> i) & 1)
		b := int((y >> i) & 1)
		if a != b {
			result++
		}
	}
	return result
}
