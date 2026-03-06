package main

import (
	"Manual_bill/internal/config"
	"Manual_bill/internal/git"
	"bufio"
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

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error menjalankan git:", err)
	}
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	isGit, err := git.CekGit()
	if err != nil {
		fmt.Println("error sistem:", err)
		return
	}

	if isGit {

		fmt.Println("Repo git ditemukan")

		fmt.Print("Masukan pesan commit: ")
		commitMsg, _ := reader.ReadString('\n')
		commitMsg = strings.TrimSpace(commitMsg)

		gitCommand("add", ".")
		gitCommand("commit", "-m", commitMsg)

		// supaya tidak kena error fetch first
		gitCommand("pull", "origin", "main")
		gitCommand("push")

		fmt.Println("Update repo selesai 🚀")

	} else {

		fmt.Println("folder .git tidak ada, membuat repo baru...")

		fmt.Print("Masukan nama repo: ")
		repoName, _ := reader.ReadString('\n')
		repoName = strings.TrimSpace(repoName)

		gitCommand("init")

		fmt.Print("Masukan Token Github: ")
		token, _ := reader.ReadString('\n')
		token = strings.TrimSpace(token)

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

		fmt.Println("Repo berhasil dibuat 🎉 :", repoName)
	}
}
