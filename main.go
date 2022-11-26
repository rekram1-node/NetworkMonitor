package main

import (
	"flag"
	"os"

	"github.com/rekram1-node/NetworkMonitor/scheduled"
)

const (
	timeLayout = "2006-01-02 3:04:05 PM"
)

func main() {
	home, _ := os.UserHomeDir()
	dir := home + "/network-monitoring"

	init := flag.Bool("init", false, "initialize scripts")
	flag.Parse()

	if *init {
		scheduled.Initialize(dir)
	}

	uploadData := flag.Bool("upload", false, "upload file")
	flag.Parse()

	if *uploadData {
		scheduled.UploadFile(dir)
	} else {
		scheduled.ConnectionCheck(dir, timeLayout)
	}
}
