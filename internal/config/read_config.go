package config

import (
	"fmt"
	"encoding/json"
	"os"
)


func Read() (Config, error) {
	config_path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(config_path)
	if err != nil {
		return Config{}, fmt.Errorf("error reading file at %s", config_path)
	}

	config := Config{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, fmt.Errorf("error unmarshaling json data from config file")
	}

	return config, nil
}