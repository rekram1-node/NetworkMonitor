package monitor

import (
	"net"
	"strconv"
	"sync"
	"time"

	utils "github.com/rekram1-node/NetworkMonitor/utilityFunctions"
)

func ConnectedToInternet() bool {
	valid := 0
	sites := utils.RandomSites()

	var wg = &sync.WaitGroup{}
	for _, siteName := range sites {
		wg.Add(1)

		go func(site string) {
			address := site + ":" + strconv.Itoa(443)
			conn, err := net.DialTimeout("tcp", address, 2*time.Second)

			if err == nil {
				valid++
				defer conn.Close()
			}
			wg.Done()
		}(siteName)
	}
	wg.Wait()

	if valid > 0 {
		return true
	}

	return false
}
