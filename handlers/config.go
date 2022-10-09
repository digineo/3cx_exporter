package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/digineo/3cx_exporter/models"
	"go.uber.org/zap"
)

func GetStateHandler(s stateProvisor, logger *zap.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		state := s.CheckConnection()
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(state); err != nil {
			logger.Error("Failed encode data",
				zap.Error(err),
			)
		}
	})
}

func GetConfigGetterHandler(logger *zap.Logger, configPath string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conf := models.Config{ConfigPath: configPath}
		config, err := conf.Get()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			logger.Error("Request error", zap.Error(err))
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(config)

	})
}

func GetSetConfigHandler(logger *zap.Logger, configurer configurer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		conf := models.Config{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&conf)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			logger.Debug("Bad request", zap.Error(err))
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		err = conf.Set()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			logger.Debug("Saving config error", zap.Error(err))
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		err = configurer.SetCreds(conf.Host, conf.Login, conf.Password, true)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			logger.Debug("Login error", zap.Error(err))
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return

		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(conf)

	})
}
