package scheduled

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type autoUpdate struct {
	Status   bool
	Interval int
}

type secrets struct {
	ApiKey        string
	PublishScript string
}

type Config struct {
	UpdateConfig autoUpdate
	Secrets      secrets
}

func Initialize(dir string) {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Fatal("failed to create directory: network-monitoring")
	}

	filePath := dir + "/config.yaml"

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	defer f.Close()

	if err != nil {
		log.Fatal("failed to create configuration file at " + filePath)
	}

	autoUpdate := autoUpdate{
		Status:   true,
		Interval: 5,
	}

	s := secrets{
		ApiKey: "fakeApiKey",
	}

	cfg := map[string]Config{
		"network-monitor": {
			autoUpdate,
			s,
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
