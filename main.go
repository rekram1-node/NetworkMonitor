package main

import (
	"flag"
	"fmt"
	"log"
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
	uploadData := flag.Bool("upload", false, "upload file")

	flag.Parse()

	if *init {
		fmt.Println("initializing application...")
		scheduled.Initialize(dir)
	}

	if *uploadData {
		fmt.Println("uploading...")
		err := scheduled.UploadFile(dir)
		checkErr(err)
	} else {
		fmt.Println("checking connection...")
		scheduled.ConnectionCheck(dir, timeLayout)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
