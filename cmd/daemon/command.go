package daemon

import (
	"github.com/spf13/cobra"

	"github.com/matthiasharzer/daily-wallpaper/cmd/daemon/run"
)

func init() {
	Command.AddCommand(run.Command)
}

var Command = &cobra.Command{
	Use:   "daemon",
	Short: "The daemon runs in the background to update the wallpaper periodically",
}
