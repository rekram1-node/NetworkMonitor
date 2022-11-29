package logger

import (
	"log"
	"os"
	"runtime"
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
	white         = "\033[97m"
	errPrefix     = red + "ERROR:\t"
	warnPrefix    = yellow + "WARN:\t"
	infoPrefix    = white + "INFO:\t" + reset
	recoverPrefix = green + "RECOVERED:\t" + reset
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
		log: log.New(os.Stdout, infoPrefix, log.Ldate|log.Ltime|log.Lshortfile),
	}
	Recover = &NetworkLog{
		log: log.New(os.Stdout, recoverPrefix, log.Ldate|log.Ltime|log.Lshortfile),
	}
	Warn = &NetworkLog{
		log: log.New(os.Stdout, warnPrefix, log.Ldate|log.Ltime|log.Lshortfile),
	}
	Error = &NetworkLog{
		log: log.New(os.Stdout, errPrefix, log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *NetworkLog) Msg(message string) {
	switch l.log.Prefix() {
	case errPrefix:
		l.log.Println(red + message + reset)
	case warnPrefix:
		l.log.Println(yellow + message + reset)
	case recoverPrefix:
		l.log.Println(green + message + reset)
	default:
		l.log.Println(white + message + reset)
	}
}
