package utils

import (
	"io/fs"
	"os"
)

const (
	DirPerms  fs.FileMode = 0755
	FilePerms fs.FileMode = 0644
)

func CreateDir(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return os.MkdirAll(dirPath, DirPerms)
	}
	return nil
}

func CreateFile(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return os.WriteFile(filePath, []byte(""), FilePerms)
	}
	return nil
}

func CopyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	return os.WriteFile(dst, data, FilePerms)
}
