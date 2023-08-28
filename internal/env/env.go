package env

import (
	"fmt"
	"os"
	"strings"

	"github.com/devkcud/dotfile-tool/internal/config"
	"github.com/devkcud/dotfile-tool/internal/tool"
	"golang.org/x/exp/slices"
)

func UpdateEnv() {
	// Generally cause it's not a good idea to include them this (--force flag will add them anyway)
	skip := []string{
		"HOME",
		"PWD",
		"OLDPWD",
		"SHELL",
		"SHLVL",
		"TERM",
		"USER",
		"LOGNAME",
		"HOST",
		"HOSTNAME",
		"VISUAL",
		"PAGER",
		"LESS",
		"MAIL",
		"MAILCHECK",
		"SSH_AGENT_PID",
		"SSH_AUTH_SOCK",
		"GPG_AGENT_INFO",
		"GPG_AGENT_PID",
		"COLORTERM",
	}

	var out string

	for _, env := range config.Current.LookupEnv {
		if slices.Contains(skip, env) {
			continue
		}

		envValue := os.Getenv(env)

		envValue = strings.ReplaceAll(envValue, os.Getenv("HOME"), "$HOME")

		if strings.TrimSpace(envValue) != "" {
			out += fmt.Sprintf("$%s=%s\n", env, envValue)
		}
	}

	if err := os.WriteFile(tool.EnvFile, []byte(out), 0644); err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("Updated env file with %d vars\n", len(strings.Split(out, "\n")))
}
