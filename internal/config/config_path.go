package config

import (
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error retrieving the user's home directory path: %s", err)
	}

	configPath := fmt.Sprintf("%s/%s", homeDir, configFileName)

	return configPath, nil
}