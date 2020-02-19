package cmd

import (
	"log"

	"github.com/ggermis/http-tester/pkg/http_tester"
	"github.com/ggermis/http-tester/pkg/http_tester/cli"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().IntVarP(&cli.Option.NumberOfRequests, "number", "n", 1, "Number of calls to perform per thread")
	rootCmd.PersistentFlags().IntVarP(&cli.Option.NumberOfThreads, "concurrent", "c", 1, "Number of threads to spin up")
	rootCmd.PersistentFlags().StringVarP(&cli.Option.Url, "url", "u", "", "URL to call")
	rootCmd.PersistentFlags().IntVarP(&cli.Option.Wait, "wait", "w", 0, "Time to wait between sequential calls (in ms)")
	rootCmd.PersistentFlags().StringVarP(&cli.Option.OutputFormat, "output", "o", cli.OUTPUT_DOT, "Output format to use. (dot, csv, detail, line, split, null)")
	rootCmd.PersistentFlags().IntVarP(&cli.Option.Randomize, "randomize", "r", 0, "Perform HTTP calls within random intervals between 0 and the specified number of milliseconds")
	rootCmd.PersistentFlags().StringVarP(&cli.Option.Method, "method", "X", "GET", "HTTP method to use")
	rootCmd.PersistentFlags().StringSliceVarP(&cli.Option.Headers, "headers", "H", []string{}, "Headers to pass (key=value). Can be specified multiple times")
	rootCmd.PersistentFlags().StringVarP(&cli.Option.Data, "data", "d", "", "data to send over HTTP")
	rootCmd.PersistentFlags().IntVarP(&cli.Option.StatusCode, "statuscode", "s", 200, "HTTP response code indicating success")
	rootCmd.PersistentFlags().StringVarP(&cli.Option.InputFile, "file", "f", "", "Input file containing HTTP calls to run in sequence on each thread")
	rootCmd.PersistentFlags().IntVar(&cli.Option.Timeout, "timeout", 30, "HTTP client timeout in seconds")
	rootCmd.PersistentFlags().IntVarP(&cli.Option.BucketSize, "bucket-size", "b", 50, "Request time statistics bucket size in ms")
	rootCmd.PersistentFlags().BoolVarP(&cli.Option.ShowVersion, "version", "v", false, "Show version")
	rootCmd.PersistentFlags().BoolVarP(&cli.Option.Quiet, "quiet", "q", false, "Be quiet")
	rootCmd.PersistentFlags().IntVar(&cli.Option.SlowRequests, "slow-requests", -1, "Indicate slow requests in the dot outputter with an 'S'. Specified in ms")
}

var rootCmd = &cobra.Command{
	Use:   "http-tester",
	Short: "High throughput HTTP call tester",

	PreRunE: func(cmd *cobra.Command, args []string) error {
		return cli.Option.Validate()
	},

	Run: func(cmd *cobra.Command, args []string) {
		if cli.Option.ShowVersion {
			http_tester.ShowVersion()
		} else {
			http_tester.StartWithStatistics()
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
