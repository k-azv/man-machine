package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config is the configuration for man-machine.
type Config struct {
	APIKey   string `mapstructure:"apiKey"`
	BaseURL  string `mapstructure:"baseURL"`
	Model    string `mapstructure:"model"`
	Language string `mapstructure:"language"`
}

// Load loads the config from the given json file
func Load() (Config, error) {
	cfgFile, err := GetConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("get config file path: %w", err)
	}

	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		fmt.Printf("Config file not detected,\n" +
			"run \"mam setup\" to initialize mam.\n")
		return Config{}, fmt.Errorf("find config file: %w", err)
	}

	v := viper.New()
	v.SetConfigFile(cfgFile)
	if err := v.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("reading config file: %w", err)
	}

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return Config{}, fmt.Errorf("decode into struct: %w", err)
	}
	return c, nil
}

func GetConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("get home dir: %w", err)
	}

	cfgRoot := filepath.Join(home, ".config")

	cfgDir := filepath.Join(cfgRoot, "mam")
	if err := os.Mkdir(cfgDir, 0o755); err != nil && !os.IsExist(err) {
		return "", fmt.Errorf("create config directory: %w", err)
	}

	cfgFile := filepath.Join(cfgDir, "config.yaml")

	return cfgFile, nil
}
