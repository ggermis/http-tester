package trace

import (
	"net/http"
	"net/http/httptrace"
	"time"

	"github.com/ggermis/http-tester/pkg/http_tester/cli"
)

type CaptureQueue struct {
	Data chan *Capture
	Done chan bool
}

func NewCaptureQueue() *CaptureQueue {
	var queue CaptureQueue
	queue.Data = make(chan *Capture, cli.Option.NumberOfRequests)
	queue.Done = make(chan bool, 1)
	return &queue
}

type Capture struct {
	ThreadId  int
	RequestId int
	Proto     string
	Method    string
	Url       string
	IpAddress string
	Start     time.Time
	Status    int
	Duration  float64
	Headers   http.Header
	Data      string
	Actions   []Action
}

type Action struct {
	Name     string
	Params   []interface{}
	Duration float64
	Total    float64
}

func (c *Capture) StartTrace(req *http.Request) *http.Request {
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), createTraceConfig(c)))
	c.Start = time.Now()
	c.RecordAction(TraceStart)
	return req
}

func (c *Capture) StopTrace(status int) {
	c.Status = status
	c.Duration = c.RecordAction(TraceFinish)
	defer s.registerCall(c)
}

func (c *Capture) RecordAction(name string, params ...interface{}) float64 {
	duration := float64(time.Since(c.Start)) / float64(time.Millisecond)
	action := Action{Name: name, Total: duration, Params: params}
	if len(c.Actions) > 0 {
		action.Duration = action.Total - c.Actions[len(c.Actions)-1].Total
	}
	c.Actions = append(c.Actions, action)
	return duration
}
