package services

import (
	"fmt"
	"sync"

	"github.com/digineo/3cx_exporter/models"
	"go.uber.org/zap"
)

type statusSaver interface {
	NewStatus(status models.InstanceState) (err error)
}

type Processor struct {
	saver        statusSaver
	statusGetter StatusGetterFactory
	logger       *zap.Logger
}

func (p *Processor) ProcessInstances(instances []models.Instance) {
	var wg sync.WaitGroup
	for _, instance := range instances {
		p.logger.Debug(fmt.Sprintf("Instance ID %v processing started", instance.InstanceId))
		wg.Add(1)
		go func(inst models.Instance) {
			client, err := p.statusGetter(inst)
			if err != nil {
				p.logger.Error(fmt.Sprintf("Instance ID %v processing error. Could not initialize client", inst.InstanceId), zap.Error(err))
				return
			}
			status, err := client.SystemStatus()
			if err != nil {
				p.logger.Error(fmt.Sprintf("Instance ID %v processing error. Could not get status", inst.InstanceId), zap.Error(err))
				return
			}
			err = p.saver.NewStatus(status)
			if err != nil {
				p.logger.Error(fmt.Sprintf("Instance ID %v processing error. Could not save state", inst.InstanceId), zap.Error(err))
			}
		}(instance)
	}
	wg.Wait()

}

func NewMonitor(saver statusSaver, getter StatusGetterFactory, logger *zap.Logger) *Processor {
	return &Processor{
		saver:        saver,
		statusGetter: getter,
		logger:       logger,
	}
}
