package http_tester

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ggermis/http-tester/pkg/http_tester/trace"
)

var resolver *Resolver

func init() {
	resolver = &Resolver{r: net.Resolver{PreferGo: true}, hosts: map[string][]net.IPAddr{}}
	go func() {
		for range time.Tick(60 * time.Second) {
			resolver.refresh()
		}
	}()
}

type Resolver struct {
	r     net.Resolver
	mu    sync.Mutex
	hosts map[string][]net.IPAddr
}

func (r *Resolver) DialTLSContext(host string, cap *trace.Capture) func(ctx context.Context, network, addr string) (net.Conn, error) {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		separator := strings.LastIndex(addr, ":")
		port, _ := strconv.Atoi(addr[separator+1:])

		cap.RecordAction(trace.TraceMark, network, addr)

		ip := r.mustResolve(host)
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

func (r *Resolver) mustResolve(host string) net.IPAddr {
	ips := r.resolve(host, false)
	if len(ips) == 0 {
		log.Panic(fmt.Sprintf("unable to resolve host '%s'", host))
	}
	return ips[rand.Intn(len(ips))]
}

func (r *Resolver) refresh() {
	for h := range r.hosts {
		r.resolve(h, true)
	}
}

func (r *Resolver) resolve(host string, force bool) []net.IPAddr {
	r.mu.Lock()
	defer r.mu.Unlock()
	ips := r.hosts[host]
	if force || len(ips) == 0 {
		ips, _ = r.r.LookupIPAddr(context.Background(), host)
		r.hosts[host] = ips
	}
	return ips
}
