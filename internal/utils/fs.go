package utils

import (
	"io/fs"
	"os"

	"github.com/devkcud/dotfile-tool/internal/tool"
)

func CreateDir(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return os.MkdirAll(dirPath, tool.DirPerms)
	}
	return nil
}

func CreateFile(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return os.WriteFile(filePath, []byte(""), tool.FilePerms)
	}
	return nil
}

func CopyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	return os.WriteFile(dst, data, tool.FilePerms)
}
