package monitor

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func CleanLogs(dir string) {

	fmt.Println("Start Cleaning.")
	filePath := dir + "/" + OutagesFileName

	f, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Could not open log", err)
	}

	tempPath := dir + "/templog"

	// Create Temp File
	templog, err := os.OpenFile(tempPath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("could not create templog", err)
	}

	lines := strings.Split(string(f), "\n")

	for i, l := range lines {
		if i == 0 {
			templog.WriteString(l)
			continue
		} else {
			currLine := strings.Split(l, " ")
			lineTime, err := time.Parse(time.Layout, currLine[0])
			if err != nil {
				log.Fatal("could not read time from log", err)
			}

			if time.Since(lineTime).Hours() < 720 {
				templog.WriteString(l)
			}
		}
	}

	os.Remove(filePath)
	os.Rename(tempPath, filePath)

}
