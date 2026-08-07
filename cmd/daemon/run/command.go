package run

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/matthiasharzer/daily-wallpaper/curator"
	"github.com/matthiasharzer/daily-wallpaper/local"
	"github.com/matthiasharzer/daily-wallpaper/logging"
	"github.com/matthiasharzer/daily-wallpaper/utils/cmdutils"
)

var interval = 1 * time.Hour
var marketArg = ""

func init() {
	Command.Flags().DurationVarP(&interval, "interval", "i", interval, "Set the interval to check for new wallpapers")
	Command.Flags().StringVarP(&marketArg, "market", "m", "", "Market code for Bing wallpapers (e.g., en-US, de-DE)")
}

var Command = &cobra.Command{
	Use:   "run",
	Short: "Run the background daemon",
	Run: func(cmd *cobra.Command, args []string) {
		market, err := cmdutils.ResolveMarket(marketArg)
		if err != nil {
			logging.Fatal("error resolving market", "err", err)
			return
		}
		if interval <= 0 {
			logging.Fatal("interval must be positive")
			return
		}
		cmd.SilenceUsage = true

		wallpaperProvider := cmdutils.GetProvider(market)
		storage := local.Storage{}

		wallpaperCurator := curator.NewCurator(wallpaperProvider, storage)

		err = wallpaperCurator.Run()
		if err != nil {
			logging.Error("failed to update wallpaper initially", "err", err)
		}

		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		var daemonErr error
		for daemonErr == nil {
			<-ticker.C
			daemonErr = wallpaperCurator.Run()
		}

		logging.Fatal("daemon error", "err", daemonErr)
	}}
