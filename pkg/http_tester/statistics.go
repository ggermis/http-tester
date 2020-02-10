package http_tester

import (
	"fmt"
	"time"

	"github.com/ggermis/http-tester/pkg/http_tester/cli"
	"github.com/ggermis/http-tester/pkg/http_tester/trace"
)

type statistics struct {
	start    time.Time
	duration time.Duration
	ctr      int64
}

func (s *statistics) calculateRps() float64 {
	return float64(stats.ctr) / float64(stats.duration/time.Millisecond) * 1000
}

func (s *statistics) show() {
	if !cli.Option.Quiet {
		fmt.Printf("\n\nFinished %d requests in %s [%0.2f rps]\n\n", stats.ctr, stats.duration, stats.calculateRps())
		trace.ShowSummary()
	}
}

var stats statistics
