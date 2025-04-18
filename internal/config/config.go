package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DB_URL  string `json:"db_url"`
	CurrentUser     string `json:"current_user_name"`
}

func (c *Config) SetUser(user_name string) error {
	c.CurrentUser = user_name

	data, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("error marshaling configuration data: %s", err)
	}

	configPath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	err = os.WriteFile(configPath, data, 0700)

	return err

}