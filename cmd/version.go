package cmd

import (
	LOG "log"

	"github.com/myaser/zettel-synch/internal/version"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Zettel",
	Run: func(cmd *cobra.Command, args []string) {
		LOG.Println("Build Date:", version.BuildDate)
		LOG.Println("Git Commit:", version.GitCommit)
		LOG.Println("Version:", version.Version)
		LOG.Println("Go Version:", version.GoVersion)
		LOG.Println("OS / Arch:", version.OsArch)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
