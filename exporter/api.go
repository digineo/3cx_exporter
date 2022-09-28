package exporter

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strings"

	"golang.org/x/net/publicsuffix"
)

// API is the interface to 3CX
type API struct {
	hostname string
	username string
	password string
	client   *http.Client
}

// ErrAuthentication is returned on HTTP status 401
var ErrAuthentication = errors.New("authentication failed")

func (api *API) SetCreds(hostname, username, password string, skipVerify bool) error {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: skipVerify},
	}
	api.client = &http.Client{Transport: tr}
	api.hostname = hostname
	api.username = username
	api.password = password
	return api.login()

}

// Login creates a user session
func (api *API) login() error {
	jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	api.client.Jar = jar

	credentials := struct {
		Username string
		Password string
	}{api.username, api.password}

	body, _ := json.Marshal(&credentials)

	resp, err := api.client.Post(api.buildURI("login"), "application/json", bytes.NewReader(body))
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

	return nil
}

func (api *API) buildURI(path string) string {
	return fmt.Sprintf("https://%s/api/%s", api.hostname, path)
}

// getResponse does a GET request and parses the JSON response
func (api *API) getResponse(path string, response interface{}) error {
	retried := false

request:
	resp, err := api.client.Get(api.buildURI(path))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		if contentType := resp.Header.Get("Content-Type"); !strings.HasPrefix(contentType, "application/json") {
			return fmt.Errorf("unexpected content-type: %s", contentType)
		}

		return json.NewDecoder(resp.Body).Decode(response)
	case 401:
		// session expired
		if retried {
			return ErrAuthentication
		}

		// try to login again
		err = api.login()
		if err != nil {
			return err
		}

		// login successful, retry request
		retried = true
		goto request
	default:
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
}
