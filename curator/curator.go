package curator

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	systemWallpaper "github.com/reujab/wallpaper"

	"github.com/matthiasharzer/daily-wallpaper/app"
	"github.com/matthiasharzer/daily-wallpaper/local"
	"github.com/matthiasharzer/daily-wallpaper/logging"
	"github.com/matthiasharzer/daily-wallpaper/wallpaper"
)

type Curator struct {
	wallpaperProvider wallpaper.Provider
	storage           local.Storage
}

func NewCurator(wallpaperProvider wallpaper.Provider, storage local.Storage) Curator {
	return Curator{
		wallpaperProvider: wallpaperProvider,
		storage:           storage,
	}
}

func (c *Curator) downloadWallpaper(url string, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create wallpaper file: %w", err)
	}
	defer func() {
		_ = file.Close()
	}()

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download wallpaper: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download wallpaper: received status code %d", resp.StatusCode)
	}
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write wallpaper to file: %w", err)
	}
	return nil
}

func (c *Curator) Run() error {
	logging.Info("Checking for new wallpaper")

	latestWallpaperURL, err := c.wallpaperProvider.GetWallpaperURL()
	if err != nil {
		return fmt.Errorf("failed to get latest wallpaper URL: %w", err)
	}
	if latestWallpaperURL == "" {
		return fmt.Errorf("unexpectedly empty wallpaper URL")
	}

	config, err := c.storage.ReadConfig()
	if err != nil {
		return fmt.Errorf("failed to get current wallpaper URL: %w", err)
	}

	currentSystemWallpaper, err := systemWallpaper.Get()
	if err != nil {
		return fmt.Errorf("failed to get current system wallpaper: %w", err)
	}

	didWallpaperURLChange := config.CurrentWallpaperURL != latestWallpaperURL       // there is a new wallpaper
	didWallpaperFileChange := config.CurrentWallpaperFile != currentSystemWallpaper // another app set the wallpaper
	didWallpaperChange := didWallpaperURLChange || didWallpaperFileChange

	if !didWallpaperChange {
		logging.Info("No new wallpaper found.")
		return nil
	}

	wallpaperFile := filepath.Join(app.StorageDirectory, "current_wallpaper")
	err = c.downloadWallpaper(latestWallpaperURL, wallpaperFile)
	if err != nil {
		return fmt.Errorf("failed to download wallpaper: %w", err)
	}

	err = systemWallpaper.SetFromFile(wallpaperFile)
	if err != nil {
		return fmt.Errorf("failed to set wallpaper: %w", err)
	}

	err = c.storage.WriteConfig(local.Config{
		CurrentWallpaperURL:  latestWallpaperURL,
		CurrentWallpaperFile: wallpaperFile,
	})
	if err != nil {
		return fmt.Errorf("failed to set current wallpaper URL: %w", err)
	}

	logging.Info("New wallpaper set successfully.", "url", latestWallpaperURL, "file", wallpaperFile)
	return nil
}
