package output

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/ggermis/http-tester/pkg/http_tester/cli"
	"github.com/ggermis/http-tester/pkg/http_tester/trace"
	"gopkg.in/yaml.v2"
)

const (
	TimeFormat = "2006-01-02T15:04:05.000Z07:00"
)

func nullOutputter() Outputter {
	return func(queue *trace.CaptureQueue) {
		queue.Done <- true
	}
}

func detailOutputter() Outputter {
	return func(queue *trace.CaptureQueue) {
		for c := range queue.Data {
			if c.Status != cli.Option.StatusCode {
				data, _ := yaml.Marshal(&c)
				_, _ = fmt.Fprintf(os.Stderr, "---\n%s\n", data)
			} else {
				if cli.Option.SlowRequests < 0 || (cli.Option.SlowRequests > 0 && int(c.Duration) >= cli.Option.SlowRequests) {
					data, _ := yaml.Marshal(&c)
					fmt.Printf("---\n%s\n", data)
				}
			}
		}
		queue.Done <- true
	}
}

func dotOutputter() Outputter {
	return func(queue *trace.CaptureQueue) {
		for c := range queue.Data {
			if c.Status == cli.Option.StatusCode {
				if cli.Option.SlowRequests > 0 && int(c.Duration) >= cli.Option.SlowRequests {
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
			fmt.Printf("%03d-%06d,%s,%s,%s,%d,%0.2f\n",
				c.ThreadId, c.RequestId, c.IpAddress, c.Method, c.Url, c.Status, c.Duration)
		}
		queue.Done <- true
	}
}

func lineOutputter() Outputter {
	return func(queue *trace.CaptureQueue) {
		for c := range queue.Data {
			if cli.Option.SlowRequests < 0 || (cli.Option.SlowRequests > 0 && int(c.Duration) >= cli.Option.SlowRequests) {
				buf := bytes.NewBufferString("")
				for _, a := range c.Actions {
					buf.Write([]byte(fmt.Sprintf("%s %0.2f ", a.Name, a.Duration)))
				}
				fmt.Println(fmt.Sprintf("%s %s %0.2f %s", c.Headers.Get("X-Joyn-Request-Id"), c.IpAddress, c.Duration, buf.String()))
			}
		}
		queue.Done <- true
	}
}

func splitOutputter() Outputter {
	return func(queue *trace.CaptureQueue) {
		for c := range queue.Data {
			if cli.Option.SlowRequests < 0 || (cli.Option.SlowRequests > 0 && int(c.Duration) >= cli.Option.SlowRequests) {
				var connection, response float64
				for _, a := range c.Actions {
					switch a.Name {
					case trace.TraceDns, trace.TraceDialTCP, trace.TraceTLSHandshake:
						connection += a.Duration
					default:
						response += a.Duration
					}
				}
				fmt.Printf("%s\t%s\t%s\t%0.2f\tconnection\t%0.2f\tresponse\t%0.2f\n", time.Now().Format(TimeFormat), c.Headers.Get("X-Joyn-Request-Id"), c.IpAddress, c.Duration, connection, response)
			}
		}
		queue.Done <- true
	}
}
