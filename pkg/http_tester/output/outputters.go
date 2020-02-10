package output

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
	"codenut.org/http-tester/pkg/http_tester/cli"
	"codenut.org/http-tester/pkg/http_tester/trace"
)

func nullOutputter() Outputter {
	return func(queue *trace.CaptureQueue) {
		queue.Done <- true
	}
}

func detailOutputter() Outputter {
	return func(queue *trace.CaptureQueue) {
		for i := range queue.Data {
			data, _ := yaml.Marshal(&i)
			if i.Status == cli.Option.StatusCode {
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
		for i := range queue.Data {
			if i.Status == cli.Option.StatusCode {
				fmt.Print(".")
			} else {
				fmt.Print(fmt.Sprintf("[%d]", i.Status))
			}
		}
		queue.Done <- true
	}
}

func csvOutputter() Outputter {
	return func(queue *trace.CaptureQueue) {
		for i := range queue.Data {
			fmt.Printf("%03d-%06d,%s,%s,%d,%t,%0.2f\n",
				i.ThreadId, i.RequestId, i.Method, i.Url, i.Status, i.TlsHandshake, i.Duration)
		}
		queue.Done <- true
	}
}
