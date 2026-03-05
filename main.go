package main

import (
	"fmt"
	"os"
)

func main() {
	foldname := ".git"

	info, err := os.Stat(foldname)

	if os.IsNotExist(err) {
		fmt.Printf("folder '%s' tidak di temukan cik. \n", foldname)
	} else if err != nil {
		fmt.Printf("Terjadi eroor tidak terduga: %v\n", err)
	} else {
		if info.IsDir() {
			fmt.Printf("Folder '%s' ditemukan wak\n", foldname)
		} else {
			fmt.Printf("Path '%s' nya ada tapi bukan berbentuk folder. \n ", foldname)
		}
	}
}
