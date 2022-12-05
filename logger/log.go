package logger

import (
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

const (
	timeLayout = "2006-01-02 3:04:05 PM"
)

var (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	// blue          = "\033[34m"
	// purple        = "\033[35m"
	// cyan          = "\033[36m"
	// gray          = "\033[37m"
	white = "\033[97m"

	errPrefix     = red + "ERROR:\t"
	warnPrefix    = yellow + "WARN:\t"
	infoPrefix    = white + "INFO:\t"
	recoverPrefix = green + "RECOVERED:\t"
	// format        = ""
)

type NetworkLog struct {
	log *log.Logger
}

var Info *NetworkLog
var Recover *NetworkLog
var Warn *NetworkLog
var Error *NetworkLog

func init() {
	if runtime.GOOS == "windows" {
		reset = ""
		red = ""
		green = ""
		yellow = ""
		// blue = ""
		// purple = ""
		// cyan = ""
		// gray = ""
		white = ""
	}

	Info = &NetworkLog{
		log: log.New(os.Stdout, infoPrefix, 0),
	}
	Recover = &NetworkLog{
		log: log.New(os.Stdout, recoverPrefix, 0),
	}
	Warn = &NetworkLog{
		log: log.New(os.Stdout, warnPrefix, 0),
	}
	Error = &NetworkLog{
		log: log.New(os.Stdout, errPrefix, 0),
	}
}

func (l *NetworkLog) Msg(message string) {
	format := time.Now().Format(timeLayout) + " : "
	prefix := l.log.Prefix()
	switch {
	case strings.Contains(prefix, errPrefix):
		l.log.Println(red + format + message + reset)
	case strings.Contains(prefix, warnPrefix):
		l.log.Println(yellow + format + message + reset)
	case strings.Contains(prefix, recoverPrefix):
		l.log.Println(green + format + message + reset)
	default:
		l.log.Println(white + format + message + reset)
	}
}
