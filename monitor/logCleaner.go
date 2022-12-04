package monitor

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
	"time"

	"github.com/rekram1-node/NetworkMonitor/logger"
)

func CleanLogs(dir string) {
	logger.Info.Msg("Start Cleaning...")
	filePath := dir + "/" + OutagesFileName

	f, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal("Could not open log", err)
	}

	tempPath := dir + "/templog"

	// Create Temp File
	templog, err := os.OpenFile(tempPath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("could not create templog ", err)
	}

	if err := templog.Close(); err != nil {
		log.Fatal(err)
	}

	w := csv.NewWriter(templog)

	lines := strings.Split(string(f), "\n")

	for i, l := range lines {
		if l == "" {
			continue
		}

		if i == 0 {
			w.Write([]string{
				"Time",
				"Duration",
			})

			continue
		} else {
			currLine := strings.Split(l, ",")
			lineTime, err := time.Parse(TimeLayout, currLine[0])

			if err != nil {
				log.Fatal("could not read time from log ", err)
			}

			if time.Since(lineTime).Hours() < 720 {
				w.Write(currLine)
			}
		}
	}

	defer w.Flush()
	if err = os.Rename(tempPath, filePath); err != nil {
		log.Fatal("could not rename file")
	}
}
