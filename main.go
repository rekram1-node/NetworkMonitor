package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/rekram1-node/NetworkMonitor/logger"
	"github.com/rekram1-node/NetworkMonitor/scheduled"
)

const (
	timeLayout = "2006-01-02 3:04:05 PM"
)

var (
	dir = ""
)

func main() {
	if home, _ := os.UserHomeDir(); home != "/" {
		dir = home + "/network-monitoring"
	} else {
		dir = "network-monitoring"
	}

	init := flag.Bool("init", false, "initialize scripts")
	uploadData := flag.Bool("upload", false, "upload file")
	flag.Parse()

	if *init {
		logger.Info.Msg("Initializing Application...")
		scheduled.Initialize(dir)
		return
	}

	if *uploadData {
		logger.Info.Msg("Uploading...")
		err := scheduled.UploadFile(dir)
		checkErr(err)
		return
	}
	logger.Info.Msg("Checking Connection...")
	cmd, _ := exec.Command("ls").Output()
	fmt.Println(string(cmd))
	fmt.Println()
	cmd, _ = exec.Command("pwd").Output()
	fmt.Println(string(cmd))
	// scheduled.ConnectionCheck(dir, timeLayout)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
