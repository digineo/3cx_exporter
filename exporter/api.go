package exporter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

func NewAPI(hostname, username, password string) API {
	return API{
		hostname: hostname,
		username: username,
		password: password,
	}
}

// ErrAuthentication is returned on HTTP status 401
var ErrAuthentication = errors.New("authentication failed")

// Login creates a user session
func (api *API) Login() error {
	jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	client := &http.Client{Jar: jar}

	credentials := struct {
		Username string
		Password string
	}{api.username, api.password}

	body, _ := json.Marshal(&credentials)

	resp, err := client.Post(api.buildURI("login"), "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
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
		err = api.Login()
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
