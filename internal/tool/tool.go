package tool

import (
	"fmt"
	"os"
	"path/filepath"

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

		fmt.Printf("Adding: %s\n", path)

		smt, err := os.Stat(path)
		if err != nil {
			fmt.Println("error:", err)
			continue
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

		fmt.Println("Done")
	}
}
