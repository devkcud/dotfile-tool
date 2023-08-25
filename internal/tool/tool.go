package tool

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/devkcud/dotfile-tool/internal/config"
	"github.com/devkcud/dotfile-tool/internal/utils"
)

func Add(args []string, shared bool) {
	if shared {
		fmt.Printf("Using shared directory: %s\n", SharedDir)
	}

	for _, path := range args {
		outdir := filepath.Join(ConfigDir, path)

		if shared {
			outdir = filepath.Join(SharedDir, path)
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
		fmt.Printf("Using shared directory: %s\n", SharedDir)
	}

	for _, path := range args {
		outdir := filepath.Join(ConfigDir, path)

		if shared {
			outdir = filepath.Join(SharedDir, path)
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

func List(shared bool) {
	entries, err := os.ReadDir(ConfigDir)

	if shared {
		entries, err = os.ReadDir(SharedDir)
	}

	if err != nil {
		fmt.Println("error", err)
		return
	}

	for _, entry := range entries {
		name := entry.Name()

		if entry.IsDir() {
			name += "/"
		}

		fmt.Println(name)
	}
}
