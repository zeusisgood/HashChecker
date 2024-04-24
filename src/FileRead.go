package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/sqweek/dialog"
)

func FileRead() (readfile string, err error) {
	targetfile, err := dialog.File().Load()
	if err != nil {
		log.Fatal("File Read Error", err)
	}
	file, err := os.Open(targetfile)
	if err != nil {
		log.Fatal("File Open Error", err)
	}
	defer file.Close()
	bf := bufio.NewReader(file)
	fmt.Println(bf)
	return bf, nil
}
