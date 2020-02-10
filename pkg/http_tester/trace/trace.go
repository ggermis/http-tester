package trace

import (
	"crypto/tls"
	"fmt"
	"net/http/httptrace"
	"net/textproto"
)

func createTraceConfig(cap *Capture) *httptrace.ClientTrace {
	return &httptrace.ClientTrace{
		GetConn:              func(hostPort string) { cap.recordAction("GetConn", hostPort) },
		GotConn:              func(i httptrace.GotConnInfo) { cap.recordAction("GotConn", i) },
		PutIdleConn:          func(err error) { cap.recordAction("PutIdleConn", err) },
		GotFirstResponseByte: func() { cap.recordAction("GotFirstResponseByte") },
		Got100Continue:       func() { cap.recordAction("Got100Continue") },
		Got1xxResponse: func(code int, header textproto.MIMEHeader) error {
			cap.recordAction("Got1xxResponse", code, header)
			return nil
		},
		DNSStart:     func(i httptrace.DNSStartInfo) { cap.recordAction("DNSStart", i.Host) },
		DNSDone:      func(i httptrace.DNSDoneInfo) { cap.recordAction("DNSDone", i.Addrs) },
		ConnectStart: func(network, addr string) { cap.recordAction("ConnectStart", network, addr) },
		ConnectDone: func(network, addr string, err error) {
			cap.recordAction("ConnectDone", network, addr, err)
		},
		TLSHandshakeStart: func() { cap.recordAction("TLSHandshakeStart") },
		TLSHandshakeDone: func(state tls.ConnectionState, err error) {
			cap.TlsHandshake = true
			cap.recordAction("TLSHandshakeDone", fmt.Sprintf("%+v", state), err)
		},
		WroteHeaderField: func(key string, value []string) {
			cap.recordAction("WroteHeaderField", fmt.Sprintf("%s", key), fmt.Sprintf("%v", value))
		},
		WroteHeaders:    func() { cap.recordAction("WroteHeaders") },
		Wait100Continue: func() { cap.recordAction("Wait100Continue") },
		WroteRequest:    func(i httptrace.WroteRequestInfo) { cap.recordAction("WroteRequest", i) },
	}
}
