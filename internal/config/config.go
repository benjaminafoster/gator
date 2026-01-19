package config

import (
	"fmt"
	"os"
	"encoding/json"
	"path/filepath"
)

type Config struct {
	DBURL string `json:"db_url"`
	CurrentUser string `json:"current_user"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read config file: %w", err)
	}
	
	gatorConfig := Config{}
	err = json.Unmarshal(data, &gatorConfig)
	if err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal config file: %w", err)
	}
	
	return gatorConfig, nil
	
}

func (c *Config) SetUser(user string) error {
	// set current user name in struct
	c.CurrentUser = user
	
	// update config file with new data
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	data, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("failed to marshal config file: %w", err)
	}
	
	err = os.WriteFile(configFilePath, data, 0660)
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}
	
	return nil
}


func getConfigFilePath() (string, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}
	
	configFilePath := filepath.Join(userHomeDir, configFileName)
	
	return configFilePath, nil
}