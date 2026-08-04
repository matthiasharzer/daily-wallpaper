package local

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = "config.json"

type Config struct {
	CurrentWallpaperURL  string `json:"currentWallpaperURL"`
	CurrentWallpaperFile string `json:"currentWallpaperFile"`
}

func getConfigPath() string {
	return filepath.Join(StorageDirectory, configFileName)
}

type Storage struct{}

func (s *Storage) ReadConfig() (Config, error) {
	configPath := getConfigPath()

	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		return Config{}, nil
	} else if err != nil {
		return Config{}, fmt.Errorf("failed to stat config file: %w", err)
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	return config, nil
}

func (s *Storage) WriteConfig(config Config) error {
	configPath := getConfigPath()

	data, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}
	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}
	return nil
}
