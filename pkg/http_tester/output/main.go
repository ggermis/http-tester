package output

import (
	"codenut.org/http-tester/pkg/http_tester/cli"
	"codenut.org/http-tester/pkg/http_tester/trace"
)

type Outputter func(*trace.CaptureQueue)

func NewOutputter() Outputter {
	switch cli.Option.OutputFormat {
	case cli.OUTPUT_DETAIL:
		return detailOutputter()
	case cli.OUTPUT_NULL:
		return nullOutputter()
	case cli.OUTPUT_CSV:
		return csvOutputter()
	}
	return dotOutputter()
}
