package main

import (
	"fmt"
	"os"

	"github.com/matthiasharzer/daily-wallpaper/cmd/apply"
	"github.com/matthiasharzer/daily-wallpaper/cmd/version"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "daily-wallpaper",
	Short: "Update your desktop wallpaper with a new picture every day",
	RunE: func(cmd *cobra.Command, _ []string) error {
		return cmd.Help()
	},
}

func init() {
	rootCommand.AddCommand(version.Command)
	rootCommand.AddCommand(apply.Command)
}

func main() {
	err := rootCommand.Execute()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
