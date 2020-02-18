package http_tester

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/ggermis/http-tester/pkg/http_tester/trace"
)

var resolver *Resolver

func init() {
	resolver = &Resolver{r: net.Resolver{PreferGo: true}, hosts: map[string][]string{}}
	go func() {
		for range time.Tick(5 * time.Second) {
			resolver.refresh()
		}
	}()
}

type Resolver struct {
	r     net.Resolver
	mu    sync.Mutex
	hosts map[string][]string
}

func (m *Resolver) DialContext(host string, cap *trace.Capture) func(ctx context.Context, network, addr string) (net.Conn, error) {
	return func(ctx context.Context, network, addr string) (conn net.Conn, err error) {
		cap.IpAddress = m.mustResolve(host)
		separator := strings.LastIndex(addr, ":")
		return net.Dial(network, cap.IpAddress+addr[separator:])
	}
}

func (m *Resolver) mustResolve(host string) string {
	ips := m.resolve(host, false)
	if len(ips) == 0 {
		log.Panic(fmt.Sprintf("unable to resolve host '%s'", host))
	}
	return ips[rand.Intn(len(ips))]
}

func (m *Resolver) refresh() {
	for h := range m.hosts {
		m.resolve(h, true)
	}
}

func (m *Resolver) resolve(host string, force bool) []string {
	m.mu.Lock()
	defer m.mu.Unlock()
	ips := m.hosts[host]
	if force || len(ips) == 0 {
		ips, _ = m.r.LookupHost(context.Background(), host)
		m.hosts[host] = ips
	}
	return ips
}
