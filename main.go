package main

import (
	"fmt"
	"time"

	"github.com/rekram1-node/NetworkMonitor/monitor"
)

func main() {
	online := monitor.ConnectedToInternet()
	if !online {
		currentTime := time.Now()
		fmt.Println("failed at " + currentTime.Format("2006-01-02 15:04:05"))
		return
	}
	monitor.GetSpeed()
}
