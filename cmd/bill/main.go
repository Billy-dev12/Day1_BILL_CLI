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

func runGit(args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func gitAdd() {
	runGit("add", ".")
}

func gitCommit(msg string) {
	err := runGit("commit", "-m", msg)
	if err != nil {
		fmt.Println("Tidak ada perubahan untuk di commit")
	}
}

func gitPull() {
	err := runGit("pull", "--rebase", "origin", "main")
	if err != nil {
		fmt.Println("⚠ pull gagal, mencoba merge biasa...")
		runGit("pull", "--no-rebase", "origin", "main")
	}
}

func gitPush() {
	err := runGit("push")
	if err != nil {
		fmt.Println("⚠ push gagal setelah pull")
	}
}

func gitInit() {
	runGit("init")
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

		gitAdd()
		gitCommit(commitMsg)

		gitPull()
		gitPush()

		fmt.Println("Update repo selesai 🚀")

	} else {

		fmt.Println("folder .git tidak ada, membuat repo baru...")

		fmt.Print("Masukan nama repo: ")
		repoName, _ := reader.ReadString('\n')
		repoName = strings.TrimSpace(repoName)

		gitInit()

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

		gitAdd()
		gitCommit("first commit")

		fmt.Println("Repo berhasil dibuat 🎉 :", repoName)
	}
}
