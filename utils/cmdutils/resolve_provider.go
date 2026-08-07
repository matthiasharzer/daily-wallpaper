package cmdutils

import (
	"github.com/matthiasharzer/daily-wallpaper/wallpaper"
	"github.com/matthiasharzer/daily-wallpaper/wallpaper/bing"
)

func GetProvider(market string) wallpaper.Provider {
	return bing.NewWallpaperProvider(market)
}
