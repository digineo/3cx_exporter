package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Hostname string `json:"Hostname"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func parseConfig(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	return config, err
}
