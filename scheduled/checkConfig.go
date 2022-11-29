package scheduled

import (
	"os"

	"gopkg.in/yaml.v2"
)

func GetConfig(dir string) (*Config, error) {
	configFilePath := dir + "/" + ConfigFileName
	yfile, err := os.ReadFile(configFilePath)

	if err != nil {
		return nil, err
	}

	data := make(map[string]*Config)

	err = yaml.Unmarshal(yfile, &data)

	if err != nil {
		return nil, err
	}

	return data["network-monitor"], nil
}
