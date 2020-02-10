package http_tester

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

func signalHandler() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		stats.duration = time.Since(stats.start)
		stats.show()
		done <- true
	}()
	<-done
	os.Exit(1)
}
