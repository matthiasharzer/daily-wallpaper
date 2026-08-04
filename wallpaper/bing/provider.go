package bing

import (
	"fmt"

	"github.com/matthiasharzer/daily-wallpaper/wallpaper"
	"github.com/matthiasharzer/daily-wallpaper/wallpaper/bing/api"
)

type WallpaperProvider struct {
	bingAPI *api.API
}

func NewWallpaperProvider(market string) wallpaper.Provider {
	return &WallpaperProvider{
		bingAPI: api.NewAPI(market),
	}
}

func (w *WallpaperProvider) GetWallpaperURL() (string, error) {
	wallpapers, err := w.bingAPI.GetWallpapers(0, 1)
	if err != nil {
		return "", fmt.Errorf("failed to get wallpapers: %w", err)
	}

	if len(wallpapers) == 0 {
		return "", fmt.Errorf("no wallpapers found")
	}

	return wallpapers[0].FullSizeURL(), nil
}
