package scheduled

import (
	"errors"
	"fmt"
	"os/exec"
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
	data, err := GetConfig(directory)

	if err != nil {
		return errors.New("failed to read configuration file " + err.Error())
	}

	uploadScript := data.PublishScript

	if uploadScript == "" {
		return errors.New("no script to run, please add a script file name in your config file")
	}

	cfg.CustomFile = directory + "/" + uploadScript

	return nil
}
