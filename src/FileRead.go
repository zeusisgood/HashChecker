package main

import (
	"io"
	"log"
	"os"

	"github.com/sqweek/dialog"
)

func FileRead() (readfile []byte) {
	targetfile, err := dialog.File().Load()
	if err != nil {
		log.Fatal("File Read Error", err)
	}
	file, err := os.Open(targetfile)
	if err != nil {
		log.Fatal("File Open Error", err)
	}

	defer file.Close()

	readfile, err = io.ReadAll(file)
	if err != nil {
		log.Fatal("File Read Error", err)
	}
	return
}
