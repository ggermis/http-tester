package cli

import (
	"errors"
	"strings"
)

type Options struct {
	NumberOfThreads   int
	NumberOfRequests  int
	Url               string
	Wait              int
	OutputFormat      string
	ForceTLSHandshake bool
	Randomize         int
	Method            string
	Headers           []string
	Data              string
	StatusCode        int
	InputFile         string
	Timeout           int
	BucketSize        int
	Quiet             bool
	ShowVersion       bool
}

func (o *Options) HeadersAsMap() map[string]string {
	x := map[string]string{}
	for _, header := range o.Headers {
		h := strings.Split(header, ":")
		x[strings.TrimSpace(h[0])] = strings.TrimSpace(h[1])
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
