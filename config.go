package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/digineo/3cx_exporter/exporter"
)

func parseConfig(path string) (*exporter.API, error) {

	data, err := ioutil.ReadFile("fixtures/config.json")
	if err != nil {
		return nil, err
	}

	api := exporter.API{}
	return &api, json.Unmarshal(data, &api)
}
