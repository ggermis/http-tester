package cli

import (
	"errors"
	"strings"
)

const (
	OUTPUT_DETAIL = "detail"
	OUTPUT_DOT    = "dot"
	OUTPUT_CSV    = "csv"
	OUTPUT_LINE   = "line"
	OUTPUT_SPLIT  = "split"
	OUTPUT_NULL   = "null"
)

func isValidOutputFormat(outputFormat string) bool {
	switch outputFormat {
	case OUTPUT_DETAIL, OUTPUT_DOT, OUTPUT_CSV, OUTPUT_LINE, OUTPUT_SPLIT, OUTPUT_NULL:
		return true
	}
	return false
}

type Options struct {
	NumberOfThreads  int
	NumberOfRequests int
	Url              string
	Wait             int
	OutputFormat     string
	Randomize        int
	Method           string
	Headers          []string
	Data             string
	StatusCode       int
	InputFile        string
	Timeout          int
	BucketSize       int
	Quiet            bool
	ShowVersion      bool
	SlowRequests     int
}

func (o *Options) HeadersAsMap() map[string]string {
	x := map[string]string{}
	for _, header := range o.Headers {
		h := strings.Split(header, ":")
		if len(h) == 2 {
			x[strings.TrimSpace(h[0])] = strings.TrimSpace(h[1])
		}
	}
	return x
}

func (o *Options) Validate() error {
	if Option.ShowVersion {
		return nil
	}
	if Option.Url == "" && Option.InputFile == "" {
		return errors.New("URL is a required parameter")
	}
	if !isValidOutputFormat(Option.OutputFormat) {
		return errors.New("invalid output format specified")
	}
	if Option.Quiet {
		Option.OutputFormat = OUTPUT_NULL
	}
	return nil
}

var Option Options
