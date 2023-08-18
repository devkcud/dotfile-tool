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
	case "c", "config":
		fmt.Println("Reading from:", config.Path)

		subcommand := args[1:]
		if len(subcommand) == 0 {
			fmt.Printf("LookupEnv: $%s\nGitRemote: %s\nVerbose: %t\n", strings.Join(config.Current.LookupEnv, ", $"), config.Current.GitRemote, config.Current.Verbose)
			return
		}

		switch subcommand[0] {
		case "e", "env", "lookupenv":
			fmt.Println("$" + strings.Join(config.Current.LookupEnv, ", $"))
		case "g", "git", "gitremote":
			fmt.Println(config.Current.GitRemote)
		case "v", "verb", "verbose":
			fmt.Println(config.Current.Verbose)
		default:
			fmt.Printf("error: Unknown subcommand: %s\n", subcommand[0])
		}
	default:
		fmt.Printf("error: Unknown command: %s\n", command)
	}
}
