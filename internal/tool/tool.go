package tool

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/devkcud/dotfile-tool/internal/utils"
)

func Add(args []string) {
	for _, path := range args {
		fmt.Printf("Adding: %s\n", path)

		smt, err := os.Stat(path)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}

		if smt.IsDir() {
			if utils.CopyFolder(path, filepath.Join(ConfigDir, path)) != nil {
				fmt.Println("error:", err)
			}
		} else {
			if utils.CopyFile(path, filepath.Join(ConfigDir, path)) != nil {
				fmt.Println("error:", err)
			}
		}

		fmt.Println("Done")
	}
}
