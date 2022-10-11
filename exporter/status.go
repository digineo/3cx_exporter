package exporter

import (
	"fmt"
	"time"

	"github.com/digineo/3cx_exporter/models"
)

// SystemStatus represents the SystemStatus response
type SystemStatus struct {
	FQDN                      string
	Version                   string
	Activated                 bool
	MaxSimCalls               int
	MaxSimMeetingParticipants int
	CallHistoryCount          int
	ChatMessagesCount         int
	ExtensionsRegistered      int
	OwnPush                   bool
	ExtensionsTotal           int
	TrunksRegistered          int
	TrunksTotal               int
	CallsActive               int
	BlacklistedIPCount        int
	MemoryUsage               int
	PhysicalMemoryUsage       int
	FreeFirtualMemory         int64
	TotalVirtualMemory        int64
	FreePhysicalMemory        int64
	TotalPhysicalMemory       int64
	DiskUsage                 int
	FreeDiskSpace             int64
	TotalDiskSpace            int64
	CPUUsage                  int
	MaintenanceExpiresAt      *time.Time
	Support                   bool
	ExpirationDate            interface{}
	OutboundRules             int
	BackupScheduled           bool
	LastBackupDateTime        *time.Time
	ResellerName              string
	LicenseKey                string
	ProductCode               string
	WebMeetingBestMCU         int `json:"WebMeetingBestMCU"`
}

// SystemStatus fetches the system status
func (api *API) SystemStatus() (models.InstanceState, error) {
	response := SystemStatus{}
	status := models.InstanceState{}
	err := api.getResponse("SystemStatus", &response)
	if err != nil {
		return status, err
	}
	status.BlacklistSize = response.BlacklistedIPCount
	status.CallsActive = response.CallsActive
	status.CallsLimit = response.MaxSimCalls
	status.CreatedAt = time.Now()
	status.ExtensionsRegistred = response.ExtensionsRegistered
	status.ExtensionsTotal = response.ExtensionsTotal
	status.InstanceId = api.instanceId
	status.LastBackUp = *response.LastBackupDateTime
	status.MaintenceUntil = *response.MaintenanceExpiresAt
	status.ServiceCPU = fmt.Sprint(response.CPUUsage)
	status.ServiceMemory = fmt.Sprint(response.MemoryUsage)
	if response.WebMeetingBestMCU == 0 {
		status.ServiceStatus = "Auto"
	} else {
		status.ServiceStatus = fmt.Sprint(response.WebMeetingBestMCU)
	}
	status.TruncRegistred = fmt.Sprint(response.TrunksRegistered)
	return status, nil

}
