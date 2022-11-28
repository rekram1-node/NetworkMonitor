package scheduled

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/mail"
	"os"

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
	UpdateConfig autoUpdate
	// Email         string
	// Password      string
	PublishScript string
}

func validMailAddress(address string) bool {
	_, err := mail.ParseAddress(address)
	if err != nil {
		return false
	}
	return true
}

func Initialize(dir string) {
	if exists, _ := monitor.Exists(dir); exists {
		log.Fatal("app has already been initialized")
		return
	}

	err := os.MkdirAll(dir, os.ModePerm)

	if err != nil {
		log.Fatal("failed to create directory: network-monitoring")
	}

	filePath := dir + "/" + ConfigFileName

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	defer f.Close()

	if err != nil {
		log.Fatal("failed to create configuration file at " + filePath)
	}

	// var userEmail string = ""
	// var userPassword string = ""

	// fmt.Println("Please enter an email for alert purposes: ")
	// fmt.Scanln(&userEmail)

	// for !validMailAddress(userEmail) {
	// 	fmt.Println("That email was invalid, please enter a valid email address: ")
	// 	fmt.Scanln(&userEmail)
	// }

	// fmt.Println("Please enter the password for your email address: ")
	// fmt.Scanln(&userPassword)

	autoUpdate := autoUpdate{
		Status:   true,
		Interval: 5,
	}

	cfg := map[string]Config{
		"network-monitor": {
			autoUpdate,
			// userEmail,
			// userPassword,
			"",
		},
	}

	data, err := yaml.Marshal(&cfg)

	if err != nil {
		log.Fatal("failed to marshal update configuration: " + err.Error())
	}

	err = ioutil.WriteFile(filePath, data, 0)

	if err != nil {
		log.Fatal("failed to write to configuration file: " + err.Error())
	}

	fmt.Println("successfully initialized network monitor!!!")
}
