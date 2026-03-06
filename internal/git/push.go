package git

import (
	"os"
)

func CekGit() (bool, error) {
	info, err := os.Stat(".git")
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}
