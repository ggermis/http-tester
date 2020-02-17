package output

import (
	"fmt"
	"os"

	"github.com/ggermis/http-tester/pkg/http_tester/cli"
	"github.com/ggermis/http-tester/pkg/http_tester/trace"
	"gopkg.in/yaml.v2"
)

func nullOutputter() Outputter {
	return func(queue *trace.CaptureQueue) {
		queue.Done <- true
	}
}

func detailOutputter() Outputter {
	return func(queue *trace.CaptureQueue) {
		for c := range queue.Data {
			data, _ := yaml.Marshal(&c)
			if c.Status == cli.Option.StatusCode {
				fmt.Printf("---\n%s\n", data)
			} else {
				_, _ = fmt.Fprintf(os.Stderr, "---\n%s\n", data)
			}
		}
		queue.Done <- true
	}
}

func dotOutputter() Outputter {
	return func(queue *trace.CaptureQueue) {
		for c := range queue.Data {
			if c.Status == cli.Option.StatusCode {
				if int(c.Duration) >= cli.Option.SlowRequests {
					fmt.Print("S")
				} else {
					fmt.Print(".")
				}
			} else {
				fmt.Print(fmt.Sprintf("[%d]", c.Status))
			}
		}
		queue.Done <- true
	}
}

func csvOutputter() Outputter {
	return func(queue *trace.CaptureQueue) {
		for c := range queue.Data {
			fmt.Printf("%03d-%06d,%s,%s,%d,%t,%0.2f\n",
				c.ThreadId, c.RequestId, c.Method, c.Url, c.Status, c.TlsHandshake, c.Duration)
		}
		queue.Done <- true
	}
}
