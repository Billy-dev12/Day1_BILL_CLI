package main

import (
	"Manual_bill/internal/config"
	"Manual_bill/internal/git"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func validasiToken(token string) bool {
	return strings.HasPrefix(token, "ghp_") || strings.HasPrefix(token, "github_pat_")
}

func gitCommand(args ...string) {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func main() {

	isGit, err := git.CekGit()
	if err != nil {
		fmt.Print("error sistem:", err)
		return
	}

	if isGit {

		fmt.Println("Repo git ditemukan")

		// input commit message
		var commitMsg string
		fmt.Print("Masukan pesan commit: ")
		fmt.Scanln(&commitMsg)

		gitCommand("add", ".")
		gitCommand("commit", "-m", commitMsg)
		gitCommand("push")

		fmt.Println("Update repo selesai 🚀")

	} else {

		fmt.Println("folder .git tidak ada, membuat repo baru...")

		// input nama repo
		var repoName string
		fmt.Print("Masukan nama repo: ")
		fmt.Scanln(&repoName)

		repoName = fmt.Sprintf("\"%s\"", repoName)

		gitCommand("init")

		// input token github
		var token string
		fmt.Print("Masukan Token Github: ")
		fmt.Scanln(&token)

		// validasi token
		if !validasiToken(token) {
			fmt.Println("⚠ ini bukan token github")
			return
		}

		err := config.SaveToken(token)
		if err != nil {
			fmt.Println("Gagal simpan config:", err)
			return
		}

		fmt.Println("Token berhasil disimpan di config.json")

		gitCommand("add", ".")
		gitCommand("commit", "-m", "first commit")

		fmt.Println("Repo berhasil dibuat 🎉")
	}
}
