package utils

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func contains(slice []int, numberToCheck int) bool {
	for _, num := range slice {
		if num == numberToCheck {
			return true
		}
	}
	return false
}

func RandomSites() []string {
	siteArray, err := readLines("sites.txt")

	if err != nil {
		log.Fatalf("readLines: %s", err)
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
