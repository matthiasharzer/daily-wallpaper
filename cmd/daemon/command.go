package daemon

import "github.com/spf13/cobra"

var Command = &cobra.Command{
	Use:   "daemon",
	Short: "The daemon runs in the background to update the wallpaper periodically",
}
