package exporter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"

	"golang.org/x/net/publicsuffix"
)

// API is the interface to 3CX
type API struct {
	Hostname string
	Username string
	Password string
	client   *http.Client
}

func (api *API) Login() error {
	jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	client := &http.Client{Jar: jar}

	credentials := struct {
		Username string
		Password string
	}{api.Username, api.Password}

	body, _ := json.Marshal(&credentials)

	resp, err := client.Post(api.buildURI("login"), "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if string(respBody) != "AuthSuccess" {
		return fmt.Errorf("failed to login: %s", respBody)
	}

	api.client = client
	return nil
}

func (api *API) SystemStatus() (*SystemStatus, error) {
	resp, err := api.client.Get(api.buildURI("SystemStatus"))
	if err != nil {
		return nil, err
	}
	status := SystemStatus{}
	return &status, json.NewDecoder(resp.Body).Decode(&status)
}

func (api *API) TrunkList() ([]Trunk, error) {
	resp, err := api.client.Get(api.buildURI("TrunkList"))
	if err != nil {
		return nil, err
	}

	trunkList := struct {
		List []Trunk `json:"list"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&trunkList)
	if err != nil {
		return nil, err
	}

	return trunkList.List, nil
}

func (api *API) buildURI(path string) string {
	return fmt.Sprintf("https://%s/api/%s", api.Hostname, path)
}
