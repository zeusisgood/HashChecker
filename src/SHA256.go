package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func hash_sha256(readfile string) {
	file, err := os.Open(readfile)
	if err != nil {
		log.Fatal("File Open Error", err)
	}

	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}

	hashInBytes := hash.Sum(nil)
	fmt.Println("sha256", hex.EncodeToString(hashInBytes))
}
