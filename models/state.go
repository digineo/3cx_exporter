package models

import "time"

type AppState struct {
	Connected    bool      `json:"connected"`
	LastRequest  time.Time `json:"last_request"`
	RequestCount int       `json:"requests_count"`
	LicenseKey   string    `json:"key"`
	Version      string    `json:"version"`
}
