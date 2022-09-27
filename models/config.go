package models

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Host       string `json:"Hostname"`
	Login      string `json:"Username"`
	Password   string `json:"Password"`
	ConfigPath string `json:"-"`
}

func (c *Config) Set() error {

	return nil
}

func (c *Config) Get() (*Config, error) {
	data, err := ioutil.ReadFile(c.ConfigPath)
	if err != nil {
		return c, err
	}
	return c, json.Unmarshal(data, &c)
}
