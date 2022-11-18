package main

import (
	"fmt"
	"net"
	"path/filepath"

	ptoml "github.com/pelletier/go-toml"
	"github.com/pkg/errors"
)

func getModeFromConfig(configPath string) (string, error) {
	initConfigHAPath := initConfigHAPathFlags.path
	if len(configPath) > 0 {
		initConfigHAPath = configPath
	}
	fmt.Println("Sanju Changes")
	fmt.Println("Hello")
	fmt.Println("Sahiba Changes 2")
	if checkIfFileExist(initConfigHAPath) {
		config, err := ptoml.LoadFile(initConfigHAPath)
		if err != nil {
			writer.Println(err.Error())
			return "", err
		}
		if config.Get("architecture.existing_infra") != nil {
			return EXISTING_INFRA_MODE, nil
		} else if config.Get("architecture.aws") != nil {
			return AWS_MODE, nil
		} else {
			return AUTOMATE, nil
		}
	} else if checkIfFileExist(filepath.Join(initConfigHabA2HAPathFlag.a2haDirPath, "a2ha.rb")) {
		return HA_MODE, nil
	} else {
		return AUTOMATE, nil
	}
}

func checkIPAddress(ip string) error {
	if net.ParseIP(ip) == nil {
		return errors.New("Ip Address is invalid.")
	} else {
		return nil
	}
}
