package scheduled

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rekram1-node/NetworkMonitor/logger"
	"github.com/rekram1-node/NetworkMonitor/monitor"
	"gopkg.in/yaml.v2"
)

const (
	ConfigFileName = "config.yaml"
)

type autoUpdate struct {
	Status   bool
	Interval int
}

type Config struct {
	UpdateConfig  autoUpdate
	ScanFrequency string
	PublishScript string
}

func prompt(message string) string {
	//nolint
	fmt.Println(message)
	response, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	response = strings.ToLower(response)
	response = strings.TrimSpace(response)
	return response
}

func handleExistence() bool {
	defaultMsg := "You have already initialized network-monitor do you want to overwrite your current settings? (y/n)"
	response := prompt(defaultMsg)

	for response != "y" && response != "n" {
		response = prompt(defaultMsg + "\nPlease enter a valid response")
	}

	return response == "y"
}

func Initialize(dir string) {
	exists, _ := monitor.Exists(dir)

	if exists && !handleExistence() {
		return
	}

	err := os.MkdirAll(dir, os.ModePerm)

	if err != nil {
		log.Fatal("failed to create directory: network-monitoring")
	}

	filePath := dir + "/" + ConfigFileName

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		log.Fatal("failed to create configuration file at " + filePath)
	}

	autoUpdate := autoUpdate{
		Status:   true,
		Interval: 5,
	}

	cfg := map[string]Config{
		"network-monitor": {
			autoUpdate,
			"5s",
			`echo "no script to run, please add a script file name in your config file"`,
		},
	}

	data, err := yaml.Marshal(&cfg)

	if err != nil {
		log.Fatal("failed to marshal update configuration: " + err.Error())
	}

	err = os.WriteFile(filePath, data, 0)

	if err != nil {
		log.Fatal("failed to write to configuration file: " + err.Error())
	}

	defer f.Close()
	logger.Info.Msg("Successfully Initialized Network Monitor!!!")
}
