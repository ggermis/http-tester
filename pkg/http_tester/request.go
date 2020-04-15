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
	"github.com/ggermis/http-tester/pkg/http_tester/trace"
)

func httpRequestFactory() []*http.Request {
	var requests []*http.Request
	for _, scenario := range input.LoadScenario() {
		req, err := http.NewRequest(scenario.Method, scenario.Url, strings.NewReader(scenario.Data))
		if err != nil {
			log.Panic(err)
		}
		for key, value := range scenario.Headers {
			req.Header.Set(key, value)
		}
		requests = append(requests, req)
	}
	return requests
}

func doTracedHttpRequest(client *http.Client, req *http.Request, cap *trace.Capture) *trace.Capture {
	res, err := client.Do(cap.StartTrace(req))
	if err != nil {
		log.Panic(err)
	}
	cap.Headers = res.Header
	if data, err := ioutil.ReadAll(res.Body); err == nil {
		cap.Data = string(data)
	}
	_ = res.Body.Close()
	cap.StopTrace(res.StatusCode)

	return cap
}

func doHttpRequest(threadId, requestId int, queue *trace.CaptureQueue) {
	for _, request := range httpRequestFactory() {
		capture := &trace.Capture{ThreadId: threadId, RequestId: requestId, Method: request.Method, Url: request.URL.String()}
		client := &http.Client{
			Timeout: time.Duration(cli.Option.Timeout) * time.Second,
			Transport: &http.Transport{
				ForceAttemptHTTP2: true,
				DialTLSContext:    dialTLSContext(request.Host, capture),
			},
		}
		queue.Data <- doTracedHttpRequest(client, request, capture)
		atomic.AddInt64(&stats.ctr, 1)
		client.CloseIdleConnections()
	}
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
