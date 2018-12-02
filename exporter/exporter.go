package exporter

import (
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const prefix = "pbx_"

var (
	blacklistSizeDesc        = prometheus.NewDesc(prefix+"blacklist_size", "Number of blacklisted IP addresses", nil, nil)
	callsActiveDesc          = prometheus.NewDesc(prefix+"calls_active", "Number of current active calls", nil, nil)
	extensionsTotalDesc      = prometheus.NewDesc(prefix+"extensions_total", "Number of total extensions", nil, nil)
	extensionsRegisteredDesc = prometheus.NewDesc(prefix+"extensions_registered", "Number of registered extensions", nil, nil)
	backupAgeDesc            = prometheus.NewDesc(prefix+"backup_age", "Age of last backup in seconds", nil, nil)
	maintenanceRemainingDesc = prometheus.NewDesc(prefix+"maintenance_remaining", "Remaining time of maintenance in seconds", nil, nil)

	trunkRegisteredDesc = prometheus.NewDesc(prefix+"trunk_registered", "Status of trunk", []string{"name"}, nil)
)

type Exporter struct {
	API
}

func (ex *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- blacklistSizeDesc
	ch <- callsActiveDesc
	ch <- extensionsTotalDesc
	ch <- extensionsRegisteredDesc
	ch <- backupAgeDesc
	ch <- maintenanceRemainingDesc
	ch <- trunkRegisteredDesc
}

func (ex *Exporter) Collect(ch chan<- prometheus.Metric) {
	now := time.Now()

	status, err := ex.API.SystemStatus()
	if err == nil {
		ch <- prometheus.MustNewConstMetric(blacklistSizeDesc, prometheus.GaugeValue, float64(status.BlacklistedIPCount))
		ch <- prometheus.MustNewConstMetric(callsActiveDesc, prometheus.GaugeValue, float64(status.CallsActive))
		ch <- prometheus.MustNewConstMetric(extensionsTotalDesc, prometheus.GaugeValue, float64(status.ExtensionsTotal))
		ch <- prometheus.MustNewConstMetric(extensionsRegisteredDesc, prometheus.GaugeValue, float64(status.ExtensionsRegistered))

		// seconds since last backup
		backupAgo := float64(-1)
		if t := status.LastBackupDateTime; t != nil {
			backupAgo = float64(now.Sub(*t)) / float64(time.Second)
		}
		ch <- prometheus.MustNewConstMetric(backupAgeDesc, prometheus.CounterValue, backupAgo)

		// remaining time of maintenance
		maintenanceRemaining := float64(-1)
		if t := status.MaintenanceExpiresAt; t != nil {
			maintenanceRemaining = float64(t.Sub(now)) / float64(time.Second)
		}
		ch <- prometheus.MustNewConstMetric(maintenanceRemainingDesc, prometheus.CounterValue, maintenanceRemaining)
	} else {
		log.Println(err)
	}

	trunks, err := ex.API.TrunkList()
	if err == nil {
		for i := range trunks {
			trunk := trunks[i]
			labels := []string{trunk.Name}

			registered := 0
			if trunk.IsRegistered {
				registered = 1
			}
			ch <- prometheus.MustNewConstMetric(trunkRegisteredDesc, prometheus.GaugeValue, float64(registered), labels...)
		}
	} else {
		log.Println(err)
	}
}
