package scheduled

import (
	"log"
	"time"

	"github.com/rekram1-node/NetworkMonitor/logger"
	"github.com/rekram1-node/NetworkMonitor/monitor"
)

func ConnectionCheck(dir string, timeLayout string) {
	data, err := GetConfig(dir)

	if err != nil {
		log.Fatal(err)
	}

	// parses time out from string formats such as "5s" or "5h30m40s"
	sleepDuration, err := time.ParseDuration(data.ScanFrequency)

	if err != nil {
		log.Fatal(err)
	}

	for {
		time.Sleep(sleepDuration)
		online := monitor.ConnectedToInternet()

		if !online {
			startedAt := time.Now()
			formattedStart := startedAt.Format(timeLayout)
			logger.Error.Msg("lost connection")

			for !online {
				online = monitor.ConnectedToInternet()
			}

			endedAt := time.Now()
			logger.Recover.Msg("recovered connection")

			cfg := monitor.OutageConfig{
				Start:          startedAt,
				End:            endedAt,
				FormattedStart: formattedStart,
				Directory:      dir,
			}

			err := cfg.AppendLog()

			if err != nil {
				log.Fatal("misconfigured logger")
			}
		}
		logger.Info.Msg("Successfully Connected")
	}
}
