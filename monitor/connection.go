package monitor

import (
	"crypto/rand"
	"log"
	"math/big"
	"net"
	"strconv"
	"sync"
	"time"
)

func ConnectedToInternet() bool {
	valid := 0
	sites := randomSites()
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
	return valid > 0
}

func contains(slice []int, numberToCheck int) bool {
	for _, num := range slice {
		if num == numberToCheck {
			return true
		}
	}
	return false
}

func randomSites() []string {
	siteArray := []string{
		"google.com",
		"stackoverflow.com",
		"apple.com",
		"youtube.com",
		"facebook.com",
		"baidu.com",
		"yahoo.com",
		"amazon.com",
		"wikipedia.org",
		"google.co.in",
		"twitter.com",
		"qq.com",
		"live.com",
		"taobao.com",
		"bing.com",
	}

	var sites []string
	var indexes []int

	// Select 3 random sites
	for i := 0; i < 3; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(siteArray))))

		if err != nil {
			log.Fatal(err)
		}

		randomIndex := int(num.Int64())

		for contains(indexes, randomIndex) {
			num, _ = rand.Int(rand.Reader, big.NewInt(int64(len(siteArray))))
			randomIndex = int(num.Int64())
		}

		site := siteArray[randomIndex]
		indexes = append(indexes, randomIndex)
		sites = append(sites, site)
	}

	return sites
}
