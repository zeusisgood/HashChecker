package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func hash_sha256(readfile []string) {
	for i, v := range readfile { // TODO for内の処理追記
		file, err := os.Open(v)
		if err != nil {
			log.Fatal("File Open Error", err)
		}
		defer file.Close()
		hash := sha256.New()
		if _, err := io.Copy(hash, file); err != nil {
			log.Fatal(err)
		}
		hashInBytes := hash.Sum(nil)
		temp := hex.EncodeToString(hashInBytes)
		fmt.Println(v)
		fmt.Println(i+1, "sha256")
		fmt.Println(temp)
	}
}
