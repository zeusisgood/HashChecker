package main

import "fmt"

func main() {
	var input string
	readfile := FileRead()
	hash_sha256(readfile)
	fmt.Println("endと入力すると終了します")
	fmt.Scan(&input)
	if input == "end" {
		return
	}
}
