package app

import (
	"fmt"
	"os"
	"path/filepath"
)

const appName = "daily-wallpaper"

var StorageDirectory string

func init() {
	var err error

	configDir, err := os.UserConfigDir()
	if err != nil {
		panic(fmt.Errorf("failed to get user config dir: %w", err))
	}

	StorageDirectory = filepath.Join(configDir, appName)
	err = os.MkdirAll(StorageDirectory, 0755)
	if err != nil {
		panic(fmt.Errorf("failed to create storage directory: %w", err))
	}
}
