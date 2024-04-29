package main

import (
	"log"

	"github.com/sqweek/dialog"
)

func FileRead() (readfile string) {
	selectedFile, err := dialog.File().Load()
	if err != nil {
		log.Fatal("File Read Error", err)
	}
	return selectedFile
}
