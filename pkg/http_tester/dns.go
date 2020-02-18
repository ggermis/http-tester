package http_tester

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strings"
	"time"

	"github.com/ggermis/http-tester/pkg/http_tester/trace"
)

type Resolver struct {
	hosts map[string][]net.IP
}

func NewResolver() *Resolver {
	resolver := &Resolver{hosts: map[string][]net.IP{}}
	go func() {
		for range time.Tick(1 * time.Minute) {
			resolver.refresh()
		}
	}()
	return resolver
}

func (m *Resolver) DialContext(host string, cap *trace.Capture) func(ctx context.Context, network, addr string) (net.Conn, error) {
	return func(ctx context.Context, network, addr string) (conn net.Conn, err error) {
		cap.IpAddress = m.mustResolve(host)
		separator := strings.LastIndex(addr, ":")
		return net.Dial(network, cap.IpAddress+addr[separator:])
	}
}

func (m *Resolver) refresh() {
	for h := range m.hosts {
		ips, _ := net.LookupIP(h)
		m.hosts[h] = ips
	}
}

func (m *Resolver) mustResolve(host string) string {
	ips := m.resolve(host)
	if len(ips) == 0 {
		log.Panic(fmt.Sprintf("unable to resolve host '%s'", host))
	}
	return ips[rand.Intn(len(ips))].String()
}

func (m *Resolver) resolve(host string) []net.IP {
	ips := m.hosts[host]
	if len(ips) == 0 {
		ips, _ = net.LookupIP(host)
		m.hosts[host] = ips
	}
	return ips
}
