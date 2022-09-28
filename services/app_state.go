package services

import (
	"time"

	"github.com/digineo/3cx_exporter/exporter"
	"github.com/digineo/3cx_exporter/models"
)

type StatusGetter interface {
	SystemStatus() (exporter.SystemStatus, error)
}

type AppState struct {
	LastRequest  time.Time
	RequestCount int
	Status       StatusGetter
}

func (a *AppState) CheckConnection() (appState models.AppState) {
	state, err := a.Status.SystemStatus()
	if err == nil {
		appState.Connected = true
	}
	appState.Version = state.Version
	appState.LicenseKey = state.LicenseKey
	appState.LastRequest = a.LastRequest
	appState.RequestCount = a.RequestCount
	return

}
func (a *AppState) RegisterRequest() {
	a.LastRequest = time.Now()
	a.RequestCount += 1

}
