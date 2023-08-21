package tool

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/devkcud/dotfile-tool/internal/config"
	"github.com/devkcud/dotfile-tool/internal/utils"
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

func Setup() {
	if err := utils.CreateDir(WorkDir); err != nil {
		fmt.Println("error: couldn't create WorkDir", err)
		return
	}

	if err := utils.CreateDir(ConfigDir); err != nil {
		fmt.Println("error: couldn't create ConfigDir", err)
		return
	}

	if err := utils.CreateDir(SharedDir); err != nil {
		fmt.Println("error: couldn't create SharedDir", err)
		return
	}

	if err := utils.CreateFile(EnvFile); err != nil {
		fmt.Println("error: couldn't create EnvFile", err)
		return
	}

	if config.Current.AddSelf {
		if _, err := os.Stat(config.Path); err == nil {
			if err := utils.CreateDir(filepath.Join(ConfigDir, "dotfiles-tool")); err != nil {
				fmt.Println("error: couldn't create dotfiles-tool dir", err)
				return
			}

			if err := utils.CopyFile(config.Path, filepath.Join(ConfigDir, "dotfiles-tool", "cfg")); err != nil {
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
