package wallpaper

type Provider interface {
	GetWallpaperURL() (string, error)
}
