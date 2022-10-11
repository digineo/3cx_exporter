package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/digineo/3cx_exporter/models"
	"github.com/gorilla/mux"
)

type statusGetter interface {
	GetLastStatus() (status models.InstanceState, err error)
}

func NewRouter(getter statusGetter) *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		status, _ := getter.GetLastStatus()
		json.NewEncoder(w).Encode(status)

	})

	return r

}
