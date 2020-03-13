package cmd

import (
	"github.com/ggermis/http-tester/pkg/http_tester"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version of this tool",

	Run: func(cmd *cobra.Command, args []string) {
		http_tester.ShowVersion()
	},
}
