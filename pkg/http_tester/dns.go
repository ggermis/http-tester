package http_tester

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"
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
