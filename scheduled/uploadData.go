package scheduled

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"

	"gopkg.in/yaml.v2"
)

type uploadConfig struct {
	CustomFile string
}

func UploadFile(dir string) error {
	uploadCfg := &uploadConfig{}
	err := uploadCfg.checkScript(dir)

	if err != nil {
		return err
	}

	cmd, err := exec.Command("/bin/sh", uploadCfg.CustomFile).Output()

	if err != nil {
		return err
	}

	fmt.Println("your custom script output:", string(cmd))
	return nil
}

func (cfg *uploadConfig) checkScript(directory string) error {
	configFilePath := directory + "/" + ConfigFileName
	yfile, err := ioutil.ReadFile(configFilePath)

	if err != nil {
		return err
	}

	data := make(map[string]Config)

	err = yaml.Unmarshal(yfile, &data)

	if err != nil {
		return err
	}

	uploadScript := data["network-monitor"].PublishScript

	if uploadScript == "" {
		return errors.New("no script to run, please add a script file name in your config file")
	}

	cfg.CustomFile = directory + "/" + uploadScript

	return nil
}
