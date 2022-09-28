package handlers

import (
	"github.com/digineo/3cx_exporter/exporter"
	"github.com/digineo/3cx_exporter/models"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

type stateProvisor interface {
	RegisterRequest()
	CheckConnection() (appState models.AppState)
}

type configurer interface {
	SetCreds(hostname, username, password string, skipVerify bool) error
}

func NewRouter(st stateProvisor, api *exporter.API, configPath string, logger *zap.Logger) *mux.Router {
	prometheus.MustRegister(&exporter.Exporter{API: *api, Logger: logger})

	r := mux.NewRouter()
	reqCountMiddleware := getRequestCountMidleware(st)

	statusHandler := GetStateHandler(st, logger)
	confGetter := GetConfigGetterHandler(logger, configPath)
	confSetter := GetSetConfigHandler(logger, api)
	metrix := r.PathPrefix("/metrics").Subrouter()
	metrix.Use(reqCountMiddleware)
	metrix.Handle("", promhttp.Handler())

	r.Handle("/status", statusHandler).Methods("GET")
	r.Handle("/config", confGetter).Methods("GET")
	r.Handle("/config", confSetter).Methods("PUT")

	return r

}
