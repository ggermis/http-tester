package http_tester

import (
	"math/rand"
	"sync"
	"time"

	"github.com/ggermis/http-tester/pkg/http_tester/cli"
	"github.com/ggermis/http-tester/pkg/http_tester/output"
	"github.com/ggermis/http-tester/pkg/http_tester/trace"
)

func start() {
	var wg sync.WaitGroup
	for i := 1; i <= cli.Option.NumberOfThreads; i++ {
		wg.Add(1)
		go newWorkerThread(i, &wg)
	}
	wg.Wait()
}

func newWorkerThread(threadId int, wg *sync.WaitGroup) {
	defer wg.Done()

	queue := trace.NewCaptureQueue()

	go output.NewOutputter()(queue)

	for i := 1; i <= cli.Option.NumberOfRequests; i++ {
		if cli.Option.Randomize > 0 {
			time.Sleep(time.Duration(rand.Intn(cli.Option.Randomize)) * time.Millisecond)
		}
		doHttpRequest(threadId, i, queue)
		if cli.Option.Wait > 0 {
			time.Sleep(time.Duration(cli.Option.Wait) * time.Millisecond)
		}
	}
	close(queue.Data)
	<-queue.Done
}
