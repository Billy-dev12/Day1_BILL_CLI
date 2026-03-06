package main

import (
	"Manual_bill/internal/git"
	"fmt"
)

func main() {
	isGit, err := git.CekGit()
	if err != nil {
		fmt.Print("error sistem:", err)
		return
	}
	if isGit {
		fmt.Print("Gascik folder .git sudah ada")
	} else {
		fmt.Print("folder .git tidak ada ey")
	}
}
