package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/devkcud/dotfile-tool/internal/config"
)

func main() {
	config.Setup()

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("error: Need at least one argument")
		return
	}

	command := strings.ToLower(args[0])

	switch command {
	default:
		fmt.Printf("error: Unknown command: %s\n", command)
	}
}
