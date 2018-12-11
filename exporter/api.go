package exporter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strings"

	"golang.org/x/net/publicsuffix"
)

// API is the interface to 3CX
type API struct {
	Hostname string
	Username string
	Password string
	client   *http.Client
}

// Login creates a user session
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
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
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

func (api *API) buildURI(path string) string {
	return fmt.Sprintf("https://%s/api/%s", api.Hostname, path)
}

// getResponse does a GET request and parses the JSON response
func (api *API) getResponse(path string, response interface{}) error {
	resp, err := api.client.Get(api.buildURI(path))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	if contentType := resp.Header.Get("Content-Type"); !strings.HasPrefix(contentType, "application/json") {
		return fmt.Errorf("unexpected content-type: %s", contentType)
	}

	return json.NewDecoder(resp.Body).Decode(response)
}
