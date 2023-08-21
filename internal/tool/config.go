package tool

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/devkcud/dotfile-tool/internal/config"
)

const (
	DirPerms  fs.FileMode = 0755
	FilePerms fs.FileMode = 0644
)

var (
	WorkDir   = filepath.Join(os.TempDir(), "dotfile-tool")
	ConfigDir = filepath.Join(WorkDir, ".config")
	SharedDir = filepath.Join(WorkDir, "shared")
	EnvFile   = filepath.Join(WorkDir, "env")
)

func createDirIfNotExists(dirPath string, perms fs.FileMode) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return os.MkdirAll(dirPath, perms)
	}
	return nil
}

func createFileIfNotExists(filePath string, perms fs.FileMode) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return os.WriteFile(filePath, []byte(""), perms)
	}
	return nil
}

func copyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	return os.WriteFile(dst, data, FilePerms)
}

func Setup() {
	if err := createDirIfNotExists(WorkDir, DirPerms); err != nil {
		fmt.Println("error: couldn't create WorkDir", err)
		return
	}

	if err := createDirIfNotExists(ConfigDir, DirPerms); err != nil {
		fmt.Println("error: couldn't create ConfigDir", err)
		return
	}

	if err := createDirIfNotExists(SharedDir, DirPerms); err != nil {
		fmt.Println("error: couldn't create SharedDir", err)
		return
	}

	if err := createFileIfNotExists(EnvFile, FilePerms); err != nil {
		fmt.Println("error: couldn't create EnvFile", err)
		return
	}

	if config.Current.AddSelf {
		if _, err := os.Stat(config.Path); err == nil {
			if err := createDirIfNotExists(filepath.Join(ConfigDir, "dotfiles-tool"), DirPerms); err != nil {
				fmt.Println("error: couldn't create dotfiles-tool dir", err)
				return
			}

			if err := copyFile(config.Path, filepath.Join(ConfigDir, "dotfiles-tool", "cfg")); err != nil {
				fmt.Println("error: couldn't copy dotfiles-tool file", err)
			}
		}
	}

	if strings.TrimSpace(config.Current.GitRemote) == "" {
		return
	}

	os.Chdir(WorkDir)

	if err := exec.Command("git", "init").Run(); err != nil {
		fmt.Println("error: couldn't init git", err)
		return
	}
	fmt.Println("running: git init")

	remote := exec.Command("git", "remote", "add", "origin", config.Current.GitRemote)
	if err := remote.Run(); err != nil {
		fmt.Println("error: couldn't add remote", err)
		return
	}
	fmt.Println("running: git remote add origin", config.Current.GitRemote)
}
