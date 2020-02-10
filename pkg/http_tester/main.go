package http_tester

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
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
