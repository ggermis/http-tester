package cmd

import (
	"fmt"

	"github.com/ggermis/http-tester/pkg/http_tester"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(http_tester.Version)
	},
}
