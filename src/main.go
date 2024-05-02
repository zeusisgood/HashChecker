package main

import "fmt"

func main() {
	selectedFile := FileRead()
	hash_sha256(selectedFile)
	fmt.Println("Enterで終了")
	fmt.Scanln()
}
