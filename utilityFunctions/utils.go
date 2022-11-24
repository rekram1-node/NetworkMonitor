package utils

import (
	"math/rand"
	"time"
)

func contains(slice []int, numberToCheck int) bool {
	for _, num := range slice {
		if num == numberToCheck {
			return true
		}
	}
	return false
}

func RandomSites() []string {
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
		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(len(siteArray))

		for contains(indexes, randomIndex) {
			randomIndex = rand.Intn(len(siteArray))
		}

		site := siteArray[randomIndex]
		indexes = append(indexes, randomIndex)
		sites = append(sites, site)
	}

	return sites
}
