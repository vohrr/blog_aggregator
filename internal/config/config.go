package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const (
	configFile = ".gatorconfig.json"
)

func getConfigPath() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	configPath := fmt.Sprintf("%s/%s", path, configFile)
	return configPath, nil
}

func Read() (*Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	configData, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(configData, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (cfg *Config) SetUser(userName string) error {

	cfg.CurrentUserName = userName

	configData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	configPath, err := getConfigPath()

	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, configData, 0644)
	if err != nil {
		return err
	}
	return nil
}
