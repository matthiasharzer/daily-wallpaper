package apply

import (
	"github.com/matthiasharzer/service-kit/logging"
	"github.com/spf13/cobra"

	"github.com/matthiasharzer/daily-wallpaper/curator"
	"github.com/matthiasharzer/daily-wallpaper/local"
	"github.com/matthiasharzer/daily-wallpaper/utils/cmdutils"
)

var marketArg = ""

func init() {
	Command.Flags().StringVarP(&marketArg, "market", "m", "", "Market code for Bing wallpapers (e.g., en-US, de-DE)")
}

var Command = &cobra.Command{
	Use:   "apply",
	Short: "Applies the latest available wallpaper to the desktop",
	RunE: func(cmd *cobra.Command, args []string) error {
		market, err := cmdutils.ResolveMarket(marketArg)
		if err != nil {
			logging.Warn("failed to resolve market, omitting market argument: %v", err)
			market = ""
		}
		cmd.SilenceUsage = true

		provider := cmdutils.GetProvider(market)
		storage := local.Storage{}

		wallpaperCurator := curator.NewCurator(provider, storage)
		err = wallpaperCurator.Run()
		if err != nil {
			return err
		}

		return nil
	},
}
