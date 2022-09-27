package handlers

import (
	"github.com/digineo/3cx_exporter/models"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

type stateProvisor interface {
	RegisterRequest()
	CheckConnection() (appState models.AppState)
}

type apiConfigurer interface {
	SetCreds(hostname, username, password string) error
}

func NewRouter(st stateProvisor, configurer apiConfigurer, configPath string, logger *zap.Logger) *mux.Router {

	r := mux.NewRouter()
	reqCountMiddleware := getRequestCountMidleware(st)

	statusHandler := GetStateHandler(st, logger)
	confGetter := GetConfigGetterHandler(logger, configPath)
	confSetter := GetSetConfigHandler(logger, configurer)
	metrix := r.PathPrefix("/metrics").Subrouter()
	metrix.Use(reqCountMiddleware)
	metrix.Handle("", promhttp.Handler())

	r.Handle("/status", statusHandler).Methods("GET")
	r.Handle("/config", confGetter).Methods("GET")
	r.Handle("/config", confSetter).Methods("PUT")

	return r

}
