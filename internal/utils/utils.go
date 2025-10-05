package utils

import (
	"os"
	"path"
	"runtime"
)

func GetUserData() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	if runtime.GOOS == "windows" {
		home = path.Join(home, "AppData")
	}

	return home, nil
}
