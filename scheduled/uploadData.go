package scheduled

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/rekram1-node/NetworkMonitor/logger"
)

type uploadConfig struct {
	Directory  string
	CustomFile string
	Script     string
}

func UploadFile(dir string) error {
	uploadCfg := &uploadConfig{
		Directory: dir,
	}

	if err := uploadCfg.checkScript(); err != nil {
		return err
	}

	uploadCfg.createFile()
	//nolint
	cmd, err := exec.Command("/bin/sh", uploadCfg.CustomFile).Output()

	if err != nil {
		logger.Error.Msg("failed to run script: " + err.Error())
		return err
	}

	scriptOutput := fmt.Sprintf("\nYour Script Output: \n%s", string(cmd))
	logger.Info.Msg(scriptOutput)
	return nil
}

func (cfg *uploadConfig) checkScript() error {
	data, err := GetConfig(cfg.Directory)

	if err != nil {
		return errors.New("failed to read configuration file " + err.Error())
	}

	uploadScript := data.PublishScript

	if uploadScript == "" {
		return errors.New("no script to run, please add a script file name in your config file")
	}

	cfg.CustomFile = cfg.Directory + "/" + "scriptFile"
	cfg.Script = uploadScript

	return nil
}

func (cfg *uploadConfig) createFile() error {
	f, err := os.OpenFile(cfg.CustomFile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	datawriter := bufio.NewWriter(f)
	lines := strings.Split(cfg.Script, "\n")

	for _, line := range lines {
		_, err = datawriter.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	datawriter.Flush()
	defer f.Close()

	return nil
}
