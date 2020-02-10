package trace

import (
	"fmt"
	"sort"
	"sync"

	"codenut.org/http-tester/pkg/http_tester/cli"
)

var s summary

type summary struct {
	mu           sync.Mutex
	distribution map[int]int
	response     map[int]int
}

func (s *summary) registerCall(i *Capture) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.distribution[ceil(int(i.Duration), cli.Option.BucketSize)] += 1
	s.response[i.Status] += 1
}

func ceil(value, ceiling int) int {
	return ((value + ceiling - 1) / ceiling) * ceiling
}

func ShowSummary() {
	fmt.Printf("Response Status Distribution:\n\n")
	for k, v := range s.response {
		fmt.Printf("%d: %d\n", k, v)
	}

	fmt.Printf("\nResponse Time Distribution:\n\n")
	keys := make([]int, 0, len(s.distribution))
	for k := range s.distribution {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Printf("%4dms - %4dms: %d\n", k-cli.Option.BucketSize, k, s.distribution[k])
	}
}
