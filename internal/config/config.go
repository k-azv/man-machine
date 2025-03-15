package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type config struct {
	APIKey  string `json:"apiKey"`
	BaseURL string `json:"baseURL"`
	Model   string `json:"model"`
	Prompt  string `json:"prompt"`
}

var Config config

// LoadConfig loads the config from the given json file
func LoadConfig(configFile string) error {
	configData, err := os.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	err = json.Unmarshal(configData, &Config)
	if err != nil {
		return fmt.Errorf("failed to decode config: %w", err)
	}

	return nil
}
