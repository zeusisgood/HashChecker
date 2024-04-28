package main

import "fmt"

func main() {
	readfile := FileRead()
	hash_sha256(readfile)
	fmt.Println("Enterで終了")
	fmt.Scanln()
}
