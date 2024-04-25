package main

func main() {
	readfile := FileRead()
	hash_sha256(readfile)
	// fmt.Println(readfile)
}
