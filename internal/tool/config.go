package tool

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/devkcud/dotfile-tool/internal/config"
	"github.com/devkcud/dotfile-tool/internal/utils"
)

var (
	WorkDir   = filepath.Join(os.TempDir(), "dotfile-tool")
	ConfigDir = filepath.Join(WorkDir, "config")
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

	// NOTE: idk if it's optimal to call that many git commands
	if _, err := os.Stat(filepath.Join(WorkDir, ".git")); err == nil {
		gitUrl := exec.Command("git", "remote", "get-url", "origin")
		gitUrl.Dir = WorkDir

		output, _ := gitUrl.Output()

		if strings.TrimSpace(string(output)) != strings.TrimSpace(config.Current.GitRemote) {
			fmt.Println("warning: updating git remote to", config.Current.GitRemote)

			remote := exec.Command("git", "remote", "set-url", "origin", config.Current.GitRemote)
			remote.Dir = WorkDir

			if err := remote.Run(); err != nil {
				fmt.Println("error: couldn't set url remote", err)
				return
			}

		}

		return
	}
	if strings.TrimSpace(config.Current.GitRemote) == "" {
		return
	}

	last, _ := os.Getwd()
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

	os.Chdir(last)
}
