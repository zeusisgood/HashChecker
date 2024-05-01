package main

import (
	"fmt"

	"github.com/sqweek/dialog"
)

func FileRead() (selectedFileSlice []string) {
	for {
		fmt.Println("ファイルを一つずつ選択してください")
		temp, err := dialog.File().Load()
		if err != nil {
			fmt.Println("ファイル選択end")
			break
		}
		selectedFileSlice = append(selectedFileSlice, temp)
		fmt.Println(temp)
	}
	return
}
