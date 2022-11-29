package scheduled

import (
	"fmt"
	"log"
	"time"

	"github.com/rekram1-node/NetworkMonitor/monitor"
)

func ConnectionCheck(dir string, timeLayout string) {
	data, err := GetConfig(dir)
	var sleep time.Duration = data.ScanFrequency

	if err != nil {
		log.Fatal(err)
	}

	for {
		time.Sleep(sleep * time.Second)
		online := monitor.ConnectedToInternet()

		if !online {
			startedAt := time.Now()
			formattedStart := startedAt.Format(timeLayout)
			fmt.Println("failed at " + formattedStart)

			for !online {
				online = monitor.ConnectedToInternet()
			}

			endedAt := time.Now()
			fmt.Println("recovered at " + endedAt.Format(timeLayout))

			cfg := monitor.LogConfig{
				Start:          startedAt,
				End:            endedAt,
				FormattedStart: formattedStart,
				Directory:      dir,
			}

			err := monitor.AppendLog(&cfg)

			if err != nil {
				log.Fatal("misconfigured logger")
			}
		}
		fmt.Println("successfully connected")
	}
}
