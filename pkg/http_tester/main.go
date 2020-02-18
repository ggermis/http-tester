package http_tester

import (
	"fmt"
	"math/rand"
	"syscall"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	setMaxOpenFiles()
	go signalHandler()
}

func StartWithStatistics() {
	defer stats.show()
	stats.start = time.Now()
	start()
	stats.duration = time.Since(stats.start)
}

func ShowVersion() {
	fmt.Println(version)
}

func setMaxOpenFiles() {
	var rLimit syscall.Rlimit
	_ = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	rLimit.Cur = 16 * 1024
	_ = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
}
