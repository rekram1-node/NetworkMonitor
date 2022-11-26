package scheduled

import (
	"fmt"
	"log"
	"time"

	"github.com/rekram1-node/NetworkMonitor/monitor"
)

func ConnectionCheck(dir string, timeLayout string) {
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
}

func UploadFile(dir string) error {
	// f := dir + "/" + monitor.OutagesFileName
	// outagesFile, err := os.Open(f)

	// if err != nil {
	// 	return err
	// }
	// configurationFile := dir + "/config.yaml"

	// err := clearFile(outagesFile)

	// if err != nil {
	// 	return err
	// }

	return nil
}

// func clearFile(filePath string) error {
// 	err := os.Remove(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
