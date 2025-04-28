package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type config struct {
	APIKey   string `mapstructure:"apiKey"`
	BaseURL  string `mapstructure:"baseURL"`
	Model    string `mapstructure:"model"`
	Language string `mapstructure:"language"`
}

var Config config

// loadConfig loads the config from the given json file
func LoadConfig() error {
	cfgFile, err := GetConfigFilePath()
	if err != nil {
		return fmt.Errorf("get config file path: %w", err)
	}

	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		fmt.Printf("Config file not detected,\n" +
			"run \"mam setup\" to initialize mam.\n")
		return fmt.Errorf("find config file: %w", err)
	}

	viper.SetConfigFile(cfgFile)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("reading config file: %w", err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		return fmt.Errorf("decode into struct: %w", err)
	}
	return nil
}

func GetConfigFilePath() (string, error) {
	cfgRoot, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("get config root: %w", err)
	}

	cfgDir := filepath.Join(cfgRoot, "mam")
	if err := os.Mkdir(cfgDir, 0o755); err != nil && !os.IsExist(err) {
		return "", fmt.Errorf("create config directory: %w", err)
	}

	cfgFile := filepath.Join(cfgDir, "config.yaml")

	return cfgFile, nil
}
