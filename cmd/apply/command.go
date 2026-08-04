package apply

import (
	"fmt"
	"regexp"

	"github.com/spf13/cobra"

	"github.com/jeandeaual/go-locale"

	"github.com/matthiasharzer/daily-wallpaper/curator"
	"github.com/matthiasharzer/daily-wallpaper/local"
	"github.com/matthiasharzer/daily-wallpaper/wallpaper"
	"github.com/matthiasharzer/daily-wallpaper/wallpaper/bing"
)

func getProvider(market string) wallpaper.Provider {
	return bing.NewWallpaperProvider(market)
}

func validateMarket(market string) error {
	match, _ := regexp.MatchString("^[a-z]{2}-[A-Z]{2}$", market)
	if !match {
		return fmt.Errorf("invalid market format: %s", market)
	}

	return nil
}

func resolveMarket(market string) (string, error) {
	if market == "" {
		systemLocale, err := locale.GetLocale()
		if err != nil {
			return "", fmt.Errorf("failed to get system locale: %w", err)
		}

		return systemLocale, nil
	}

	err := validateMarket(market)
	if err != nil {
		return "", fmt.Errorf("invalid market: %w", err)
	}

	return market, nil
}

var marketArg = ""

func init() {
	Command.Flags().StringVarP(&marketArg, "market", "m", "en-US", "Market code for Bing wallpapers (e.g., en-US, de-DE)")
}

var Command = &cobra.Command{
	Use:   "apply",
	Short: "Applies the current wallpaper to the desktop",
	RunE: func(cmd *cobra.Command, args []string) error {
		market, err := resolveMarket(marketArg)
		if err != nil {
			return err
		}
		cmd.SilenceUsage = true

		provider := getProvider(market)
		storage := local.Storage{}

		wallpaperCurator := curator.NewCurator(provider, storage)
		err = wallpaperCurator.Run()
		if err != nil {
			return err
		}

		return nil
	},
}
