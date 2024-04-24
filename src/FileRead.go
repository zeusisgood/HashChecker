package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/sqweek/dialog"
)

func FileRead() {
	filename, err := dialog.File().Load()
	if err != nil {
		log.Fatal("File read Error", err)
		return
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(file)
	defer file.Close()
	bf := bufio.NewReader(file)

}
