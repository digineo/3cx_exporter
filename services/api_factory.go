package services

import (
	"github.com/digineo/3cx_exporter/exporter"
	"github.com/digineo/3cx_exporter/models"
)

type StatusGetter interface {
	SystemStatus() (models.InstanceState, error)
}

type StatusGetterFactory func(instance models.Instance) (StatusGetter, error)

var NewStatusGetter StatusGetterFactory = func(instance models.Instance) (StatusGetter, error) {
	getter, err := exporter.New3CXApi(instance.Host, instance.Login, instance.Password, instance.Port, true, instance.InstanceId)
	return getter, err
}
