package cmd

import (
	LOG "log"

	"github.com/myaser/zettel-synch/internal/app"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "run the main application",
	Run: func(cmd *cobra.Command, args []string) {
		app, err := app.GetApplication()
		if err != nil {
			LOG.Fatalf("can't configure application: %v", err)
		}
		if err := app.Run(); err != nil {
			LOG.Fatalf("can't configure application: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(launchCmd)
}
