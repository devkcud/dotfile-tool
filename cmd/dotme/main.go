package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/devkcud/dotfile-tool/internal/config"
	"github.com/devkcud/dotfile-tool/internal/env"
	"github.com/devkcud/dotfile-tool/internal/tool"
	"golang.org/x/exp/slices"
)

func main() {
	config.Setup()

	args := make([]string, 0)
	flags := make([]string, 0) // I don't really like the 'flag' package :/

	for _, e := range os.Args[1:] {
		if !strings.HasPrefix(e, "-") {
			args = append(args, e)
		} else {
			flags = append(flags, e)
		}
	}

	force := slices.Contains(flags, "-f") || slices.Contains(flags, "--force")
	shared := slices.Contains(flags, "-s") || slices.Contains(flags, "--shared")
	verbose := slices.Contains(flags, "-V") || slices.Contains(flags, "--verbose")

	if force {
		fmt.Println("warning: -force may cause unexpected results")
	}

	if verbose {
		config.Current.Verbose = true
	}

	tool.Setup()

	if len(args) == 0 {
		fmt.Println("error: Need at least one argument")
		return
	}

	command := strings.ToLower(args[0])

	switch command {
	case "+", "a", "add":
		tool.Add(args[1:], shared)
	case "-", "r", "rem", "remove":
		tool.Rem(args[1:], shared)
	case "l", "list":
		tool.List(args[1:], shared)
	case "c", "config":
		fmt.Println("Reading from:", config.Path)
		lookupenv := strings.Join(config.Current.LookupEnv, ", $")

		if len(lookupenv) > 0 {
			lookupenv = fmt.Sprintf("$%s", lookupenv)
		}

		subcommand := args[1:]
		if len(subcommand) == 0 {
			fmt.Printf("LookupEnv: %s\nGitRemote: %s\nVerbose: %t\n", lookupenv, config.Current.GitRemote, config.Current.Verbose)
			return
		}

		switch subcommand[0] {
		case "e", "env", "lookupenv":
			fmt.Println(lookupenv)
		case "g", "git", "gitremote":
			fmt.Println(config.Current.GitRemote)
		case "v", "verb", "verbose":
			fmt.Println(config.Current.Verbose)
		default:
			fmt.Printf("error: Unknown subcommand: %s\n", subcommand[0])
		}
	case "eu", "ue", "env-up", "up-env", "env-update", "update-env":
		env.UpdateEnv(force)
	default:
		fmt.Printf("error: Unknown command: %s\n", command)
	}

	if config.Current.Verbose {
		fmt.Println("Done")
	}
}
