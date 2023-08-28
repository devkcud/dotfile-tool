package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const (
	DirPerm  os.FileMode = 0755
	FilePerm os.FileMode = 0644
)

func CreateDir(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, DirPerm); err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
	}
	return nil
}

func CreateFile(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if err := os.WriteFile(filePath, []byte(""), FilePerm); err != nil {
			return fmt.Errorf("failed to create file: %v", err)
		}
	}
	return nil
}

func WriteFile(filePath string, out string) error {
    if err := os.WriteFile(filePath, []byte(out), FilePerm); err != nil {
        return fmt.Errorf("failed to create file: %v", err)
    }
	return nil
}

func CopyFolder(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing %s: %v", path, err)
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return fmt.Errorf("error getting relative path: %v", err)
		}

		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			if err := CreateDir(dstPath); err != nil {
				return fmt.Errorf("error creating directory %s: %v", dstPath, err)
			}
		} else if info.Mode().IsRegular() {
			if err := CopyFile(path, dstPath); err != nil {
				return fmt.Errorf("error copying file %s: %v", path, err)
			}
		}

		return nil
	})
}

func CopyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %v", err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %v", err)
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("error copying data: %v", err)
	}

	if err := dstFile.Chmod(FilePerm); err != nil {
		return fmt.Errorf("error setting file permissions: %v", err)
	}

	return nil
}
