package obsidian

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/joelvaneenwyk/obsidian-cli/pkg/config"
)

var CliConfigPath = config.CliPath
var JsonMarshal = json.Marshal

func (v *Vault) DefaultName() (string, error) {
	if v.Name != "" {
		return v.Name, nil
	}

	// get cliConfig path
	_, cliConfigFile, err := CliConfigPath()
	if err != nil {
		return "", err
	}

	// read file
	content, err := os.ReadFile(cliConfigFile)
	if err != nil {
		return "", errors.New(CliConfigReadError)
	}

	// unmarshal json
	cliConfig := CliConfig{}
	err = json.Unmarshal(content, &cliConfig)

	if err != nil {
		return "", errors.New(CliConfigParseError)
	}

	if cliConfig.DefaultVaultName == "" {
		return "", errors.New(CliConfigParseError)
	}

	v.Name = cliConfig.DefaultVaultName
	return cliConfig.DefaultVaultName, nil
}

func (v *Vault) SetDefaultName(name string) error {
	// marshal obsidian name to json
	cliConfig := CliConfig{DefaultVaultName: name}
	jsonContent, err := JsonMarshal(cliConfig)
	if err != nil {
		return errors.New(CliConfigGenerateJsonError)
	}

	// get cliConfig path
	obsConfigDir, obsConfigFile, err := CliConfigPath()
	if err != nil {
		return err
	}

	// create directory
	err = os.MkdirAll(obsConfigDir, os.ModePerm)
	if err != nil {
		return errors.New(CliConfigDirWriteError)
	}

	// create and write file
	err = os.WriteFile(obsConfigFile, jsonContent, 0644)
	if err != nil {
		return errors.New(CliConfigWriteError)
	}

	v.Name = name

	return nil
}
