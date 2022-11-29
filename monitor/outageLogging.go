package monitor

import (
	"encoding/csv"
	"errors"
	"os"
	"time"
)

const (
	OutagesFileName = "outages.csv"
)

type OutageConfig struct {
	Start          time.Time
	End            time.Time
	FormattedStart string
	Directory      string
}

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func AppendLog(cfg *OutageConfig) error {
	writeHeader := false
	filePath := cfg.Directory + "/" + OutagesFileName

	if ok, _ := Exists(filePath); !ok {
		writeHeader = true
	}

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		return err
	}

	defer f.Close()
	w := csv.NewWriter(f)
	defer w.Flush()

	duration := cfg.End.Sub(cfg.Start)

	l := []string{
		cfg.FormattedStart,
		duration.String(),
	}

	if writeHeader {
		w.Write([]string{
			"Time",
			"Duration",
		})
	}

	err = w.Write(l)

	if err != nil {
		return err
	}

	return nil
}
