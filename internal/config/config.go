package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type config struct {
	APIKey  string `json:"apiKey"`
	BaseURL string `json:"baseURL"`
	Model   string `json:"model"`
	Prompt  string `json:"prompt"`
}

var Config config

// LoadConfig loads the config from the given json file
func LoadConfig() error {
	configPath := filepath.Join(getExecutableDir(), "configs/config.json")
	configData, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	err = json.Unmarshal(configData, &Config)
	if err != nil {
		return fmt.Errorf("failed to decode config: %w", err)
	}

	return nil
}

// getExecutableDir returns the directory of the executable
func getExecutableDir() string {
	exePath, err := os.Executable()

	if err != nil {
		fmt.Println("Error getting executable path:", err)
		os.Exit(1)
	}

	println(exePath)
	return filepath.Dir(exePath)
}
