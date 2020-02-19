package output

import (
	"github.com/ggermis/http-tester/pkg/http_tester/cli"
	"github.com/ggermis/http-tester/pkg/http_tester/trace"
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
	case cli.OUTPUT_LINE:
		return lineOutputter()
	case cli.OUTPUT_SPLIT:
		return splitOutputter()
	}
	return dotOutputter()
}
