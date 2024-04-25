package main

import (
	"crypto/sha256"
	"fmt"
)

func hash_sha256(readfile []byte) {
	hash := sha256.Sum256(readfile)
	fmt.Println("sha256: " + fmt.Sprintf("%x", hash))
}
