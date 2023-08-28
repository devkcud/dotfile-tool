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

func Setup() {
	if err := utils.CreateDir(config.WorkDir); err != nil {
		fmt.Println("error: couldn't create WorkDir", err)
		return
	}

	if err := utils.CreateDir(config.ConfigDir); err != nil {
		fmt.Println("error: couldn't create ConfigDir", err)
		return
	}

	if err := utils.CreateDir(config.SharedDir); err != nil {
		fmt.Println("error: couldn't create SharedDir", err)
		return
	}

	if err := utils.CreateFile(config.EnvFile); err != nil {
		fmt.Println("error: couldn't create EnvFile", err)
		return
	}

	if config.Current.AddSelf {
		if _, err := os.Stat(config.Path); err == nil {
			if err := utils.CreateDir(filepath.Join(config.ConfigDir, "dotfiles-tool")); err != nil {
				fmt.Println("error: couldn't create dotfiles-tool dir", err)
				return
			}

			if err := utils.CopyFile(config.Path, filepath.Join(config.ConfigDir, "dotfiles-tool", "cfg")); err != nil {
				fmt.Println("error: couldn't copy dotfiles-tool file", err)
			}
		}
	}

	// NOTE: idk if it's optimal to call that many git commands
	if _, err := os.Stat(filepath.Join(config.WorkDir, ".git")); err == nil {
		gitUrl := exec.Command("git", "remote", "get-url", "origin")
		gitUrl.Dir = config.WorkDir

		output, _ := gitUrl.Output()

		if strings.TrimSpace(string(output)) != strings.TrimSpace(config.Current.GitRemote) {
			fmt.Println("warning: updating git remote to", config.Current.GitRemote)

			remote := exec.Command("git", "remote", "set-url", "origin", config.Current.GitRemote)
			remote.Dir = config.WorkDir

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
	os.Chdir(config.WorkDir)

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

func Add(args []string, shared bool) {
	if shared {
		fmt.Printf("Using shared directory: %s\n", config.SharedDir)
	}

	for _, path := range args {
		outdir := filepath.Join(config.ConfigDir, path)

		if shared {
			outdir = filepath.Join(config.SharedDir, path)
		}

		if config.Current.Verbose {
			fmt.Printf("Adding: %s\n", path)
		}

		smt, err := os.Stat(path)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}

		// Remove the file/folder if it already exists to evict merge
		if _, err := os.Stat(outdir); err == nil {
			if config.Current.Verbose {
				fmt.Printf("Removing '%s' to evict merge\n", outdir)
			}
			os.RemoveAll(outdir)
		}

		if config.Current.Verbose {
			fmt.Printf("Copying file '%s' to '%s'\n", path, outdir)
		}

		if smt.IsDir() {
			if utils.CopyFolder(path, outdir) != nil {
				fmt.Println("error:", err)
			}
		} else {
			if utils.CopyFile(path, outdir) != nil {
				fmt.Println("error:", err)
			}
		}
	}
}

func Rem(args []string, shared bool) {
	if shared {
		fmt.Printf("Using shared directory: %s\n", config.SharedDir)
	}

	for _, path := range args {
		outdir := filepath.Join(config.ConfigDir, path)

		if shared {
			outdir = filepath.Join(config.SharedDir, path)
		}

		if config.Current.Verbose {
			fmt.Printf("Removing: %s\n", outdir)
		}

		if _, err := os.Stat(outdir); err == nil {
			os.RemoveAll(outdir)
		} else {
			fmt.Println("error:", err)
		}
	}
}

func List(args []string, shared bool) {
	from := ""
	if len(args) > 0 {
		from = args[0]
	}

	readfrom := filepath.Join(config.ConfigDir, from)

	if shared {
		readfrom = filepath.Join(config.SharedDir, from)
	}

	if fileinfo, err := os.Stat(readfrom); err != nil {
		fmt.Println("error:", err)
	} else {
		if !fileinfo.IsDir() {
			fmt.Println(fileinfo.Name(), "[FILE]")
			return
		}
	}

	entries, err := os.ReadDir(filepath.Join(readfrom, from))

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	for _, entry := range entries {
		name := entry.Name()

		if entry.IsDir() {
			name += " [DIR]"
		} else {
			name += " [FILE]"
		}

		fmt.Println(name)
	}
}
