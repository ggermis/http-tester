package trace

import (
	"fmt"
	"net/http/httptrace"
	"net/textproto"
)

const (
	TraceStart        = "Start"
	TraceMark         = "MarkCustom"
	TraceDns          = "DNS"
	TraceDialTCP      = "DialTCP"
	TraceTLSHandshake = "TLSHandshake"
	TraceFinish       = "Finished"
)

func init() {
	s = summary{distribution: map[int]int{}, response: map[int]int{}}
}

func createTraceConfig(cap *Capture) *httptrace.ClientTrace {
	return &httptrace.ClientTrace{
		GetConn:              func(hostPort string) { cap.RecordAction("GetConn", hostPort) },
		GotConn:              func(info httptrace.GotConnInfo) { cap.RecordAction("GotConn", fmt.Sprintf("%+v", info)) },
		PutIdleConn:          func(err error) { cap.RecordAction("PutIdleConn", err) },
		GotFirstResponseByte: func() { cap.RecordAction("GotFirstResponseByte") },
		Got100Continue:       func() { cap.RecordAction("Got100Continue") },
		Got1xxResponse: func(code int, header textproto.MIMEHeader) error {
			cap.RecordAction("Got1xxResponse", code, header)
			return nil
		},
		DNSStart:     func(i httptrace.DNSStartInfo) { cap.RecordAction("DNSStart", i.Host) },
		DNSDone:      func(i httptrace.DNSDoneInfo) { cap.RecordAction("DNSDone", i.Addrs) },
		ConnectStart: func(network, addr string) { cap.RecordAction("ConnectStart", network, addr) },
		ConnectDone: func(network, addr string, err error) {
			cap.RecordAction("ConnectDone", network, addr, err)
		},
		WroteHeaderField: func(key string, value []string) {
			cap.RecordAction("WroteHeaderField", fmt.Sprintf("%s", key), fmt.Sprintf("%v", value))
		},
		WroteHeaders:    func() { cap.RecordAction("WroteHeaders") },
		Wait100Continue: func() { cap.RecordAction("Wait100Continue") },
		WroteRequest:    func(i httptrace.WroteRequestInfo) { cap.RecordAction("WroteRequest", i) },
	}
}
