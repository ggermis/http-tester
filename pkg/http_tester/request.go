package http_tester

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/ggermis/http-tester/pkg/http_tester/cli"
	"github.com/ggermis/http-tester/pkg/http_tester/input"
	"github.com/ggermis/http-tester/pkg/http_tester/interpolator"
	"github.com/ggermis/http-tester/pkg/http_tester/trace"
)

type HttpRequest struct {
	threadId  int
	requestId int
	scenario  *input.Scenario
	task      *input.Task
}

func doHttpRequest(threadId, requestId int, queue *trace.CaptureQueue) {
	scenario := input.LoadScenario()
	for _, task := range scenario.Tasks {
		queue.Data <- doTracedHttpRequest(&HttpRequest{threadId: threadId, requestId: requestId, scenario: &scenario, task: task})
		atomic.AddInt64(&stats.ctr, 1)
	}
}

func doTracedHttpRequest(r *HttpRequest) *trace.Capture {
	request := r.task.AsRequest(r.scenario)

	capture := &trace.Capture{ThreadId: r.threadId, RequestId: r.requestId, Method: request.Method, Url: request.URL.String()}
	client := &http.Client{
		Timeout: time.Duration(cli.Option.Timeout) * time.Second,
		Transport: &http.Transport{
			ForceAttemptHTTP2: true,
			DialTLSContext:    dialTLSContext(request.Host, capture),
		},
	}
	defer client.CloseIdleConnections()

	res, err := client.Do(capture.StartTrace(request))
	if err != nil {
		log.Panic(err)
	}
	capture.Headers = res.Header
	if data, err := ioutil.ReadAll(res.Body); err == nil {
		capture.Data = string(data)
		for key, value := range interpolator.NewParser(res.Header.Get("Content-Type")).Parse(capture.Data, r.task.Variables) {
			r.scenario.Interpolator.Register(key, value)
		}
	}
	_ = res.Body.Close()
	capture.StopTrace(res.StatusCode)

	return capture
}

func dialTLSContext(host string, cap *trace.Capture) func(ctx context.Context, network, addr string) (net.Conn, error) {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		separator := strings.LastIndex(addr, ":")
		port, _ := strconv.Atoi(addr[separator+1:])

		cap.RecordAction(trace.TraceMark, network, addr)

		ip := resolver.mustResolve(host)
		cap.RecordAction(trace.TraceDns, ip.String())

		cap.IpAddress = ip.String()

		raw, err := net.DialTCP(network, nil, &net.TCPAddr{IP: ip.IP, Port: port, Zone: ip.Zone})
		if err != nil {
			log.Panic(err)
		}
		cap.RecordAction(trace.TraceDialTCP)

		conn := tls.Client(raw, &tls.Config{ServerName: addr[:separator], MinVersion: tls.VersionTLS12})

		err = conn.Handshake()
		if err != nil {
			log.Panic(err)
		}
		cap.RecordAction(trace.TraceTLSHandshake, fmt.Sprintf("%+v", conn.ConnectionState()))

		return conn, err
	}
}
