package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Config holds CLI runtime configurations.
type Config struct {
	ConfigFile string
	Username   string
	Token      string
	Verbose    bool
	BaseURL    string
}

// defaultConfigFile returns the default config file path ($HOME/.pixela.yaml).
func defaultConfigFile() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("cannot find home directory: %w", err)
	}
	return filepath.Join(home, ".pixela.yaml"), nil
}

// loadConfigFile loads key-value pairs from a simple YAML-like config file.
// If the file does not exist, it returns empty values without error.
func loadConfigFile(path string) (username, token string, err error) {
	if path == "" {
		return "", "", nil
	}

	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", "", nil
		}
		return "", "", fmt.Errorf("failed to open config file %s: %w", path, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])
		val = strings.Trim(val, `"'`)

		switch key {
		case "username":
			username = val
		case "token":
			token = val
		}
	}

	if err := scanner.Err(); err != nil {
		return "", "", fmt.Errorf("failed to read config file %s: %w", path, err)
	}

	return username, token, nil
}

// saveConfigFile writes username and token to the target config file.
// If the file already exists, it reports an error and prints instructions to out.
func saveConfigFile(path, username, token string, out io.Writer) error {
	targetPath := path
	if targetPath == "" {
		def, err := defaultConfigFile()
		if err != nil {
			return err
		}
		targetPath = def
	}

	if _, err := os.Stat(targetPath); err == nil {
		fmt.Fprintf(out, "`user create` succeeded but there is already a config file at %s\n", targetPath)
		fmt.Fprintln(out, "You should take a note of the following settings (save to yaml file):")
		fmt.Fprintf(out, "username: %s\n", username)
		fmt.Fprintf(out, "token: %s\n", token)
		return fmt.Errorf("config file already exists at %s", targetPath)
	}

	content := fmt.Sprintf("username: %s\ntoken: %s\n", username, token)
	if err := os.WriteFile(targetPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to save config file: %w", err)
	}

	return nil
}

// resolveConfig resolves the configuration from config file, environment variables, and flags.
func resolveConfig(explicitConfig string) (*Config, error) {
	cfg := &Config{}

	configPath := explicitConfig
	if configPath == "" {
		def, err := defaultConfigFile()
		if err == nil {
			configPath = def
		}
	}
	cfg.ConfigFile = configPath

	// 1. Load from file
	fileUser, fileToken, err := loadConfigFile(configPath)
	if err != nil {
		return nil, err
	}
	cfg.Username = fileUser
	cfg.Token = fileToken

	// 2. Override with ENV
	if envUser := os.Getenv("PIXELA_USERNAME"); envUser != "" {
		cfg.Username = envUser
	}
	if envToken := os.Getenv("PIXELA_TOKEN"); envToken != "" {
		cfg.Token = envToken
	}
	if envVerbose := os.Getenv("PIXELA_VERBOSE"); envVerbose == "true" || envVerbose == "1" {
		cfg.Verbose = true
	}
	if envEndpoint := os.Getenv("PIXELA_API_ENDPOINT"); envEndpoint != "" {
		cfg.BaseURL = envEndpoint
	}

	return cfg, nil
}
