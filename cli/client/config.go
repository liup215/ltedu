package client

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const (
	EnvBaseURL = "EDU_BASE_URL"
	EnvToken   = "EDU_TOKEN"
)

// Config holds the CLI configuration.
type Config struct {
	BaseURL string `json:"base_url"`
	Token   string `json:"token"`
}

// configFilePath returns the path to the config file.
func configFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".edu-cli", "config.json"), nil
}

// LoadConfig loads configuration from environment variables first,
// falling back to the config file.
func LoadConfig() *Config {
	cfg := &Config{}

	// Load from file first
	if path, err := configFilePath(); err == nil {
		if data, err := os.ReadFile(path); err == nil {
			_ = json.Unmarshal(data, cfg)
		}
	}

	// Environment variables override file config
	if v := os.Getenv(EnvBaseURL); v != "" {
		cfg.BaseURL = v
	}
	if v := os.Getenv(EnvToken); v != "" {
		cfg.Token = v
	}

	return cfg
}

// SaveConfig persists the configuration to the config file.
func SaveConfig(cfg *Config) error {
	path, err := configFilePath()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return err
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0600)
}
