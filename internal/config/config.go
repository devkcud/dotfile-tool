package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type config struct {
	// List of environment variables to look for.
	LookupEnv []string

	// Initialize git repository with it's remote url. It won't be used if empty.
	GitRemote string

	// Add self .config files to the dotfiles directory
	AddSelf bool

	// Print more stuff
	Verbose bool
}

var (
	Current *config
	Path    string
)

var WorkDir, ConfigDir, SharedDir, EnvFile string

func Setup() {
	cfg, _ := os.UserCacheDir()

	WorkDir = filepath.Join(cfg, "workdir.dotfile-tool")
	ConfigDir = filepath.Join(WorkDir, "config")
	SharedDir = filepath.Join(WorkDir, "shared")
	EnvFile = filepath.Join(WorkDir, "user.env")

	Current = &config{
		LookupEnv: []string{},
		GitRemote: "",
		AddSelf:   true,
		Verbose:   false,
	}

	dir, _ := os.UserConfigDir()

	Path = filepath.Join(dir, "dotfile-tool", "cfg")

	if _, err := os.Stat(Path); err != nil {
		fmt.Printf("error: Config file not found at path: %s\n", Path)
		return
	}

	data, err := os.ReadFile(Path)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	for id, line := range strings.Split(string(data), "\n") {
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, value, found := strings.Cut(line, "=")
		if !found {
			fmt.Printf("error: Invalid line '%s' at line '%d'\n", line, id+1)
			continue
		}

		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)

		if key == "" || value == "" {
			fmt.Printf("error: Empty key or value at line '%d'\n", id+1)
			continue
		}

		switch strings.ToLower(key) {
		case "gitremote":
			Current.GitRemote = value
		case "verbose":
			val, err := strconv.ParseBool(value)
			if err != nil {
				fmt.Printf("error: Invalid value '%s' for key '%s' at line '%d'\n", value, key, id+1)
				continue
			}
			Current.Verbose = val
		case "addself":
			val, err := strconv.ParseBool(value)
			if err != nil {
				fmt.Printf("error: Invalid value '%s' for key '%s' at line '%d'\n", value, key, id+1)
				continue
			}
			Current.AddSelf = val
		case "lookupenv":
			values := strings.Split(strings.Trim(strings.ReplaceAll(value, " ", ""), ","), ",")
			for _, value := range values {
				value = strings.TrimSpace(value)
				if value != "" {
					Current.LookupEnv = append(Current.LookupEnv, value)
				}
			}
		default:
			fmt.Printf("error: Unknown key '%s' at line '%d'\n", key, id+1)
		}
	}
}
