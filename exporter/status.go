package exporter

import (
	"time"
)

// SystemStatus represents the SystemStatus response
type SystemStatus struct {
	FQDN                            string      `json:"FQDN"`
	Version                         string      `json:"Version"`
	Activated                       bool        `json:"Activated"`
	MaxSimCalls                     int         `json:"MaxSimCalls"`
	MaxSimMeetingParticipants       int         `json:"MaxSimMeetingParticipants"`
	CallHistoryCount                int         `json:"CallHistoryCount"`
	ChatMessagesCount               int         `json:"ChatMessagesCount"`
	ExtensionsRegistered            int         `json:"ExtensionsRegistered"`
	OwnPush                         bool        `json:"OwnPush"`
	AvailableLocalIps               string      `json:"AvailableLocalIps"`
	ExtensionsTotal                 int         `json:"ExtensionsTotal"`
	HasUnregisteredSystemExtensions bool        `json:"HasUnregisteredSystemExtensions"`
	HasNotRunningServices           bool        `json:"HasNotRunningServices"`
	TrunksRegistered                int         `json:"TrunksRegistered"`
	TrunksTotal                     int         `json:"TrunksTotal"`
	CallsActive                     int         `json:"CallsActive"`
	BlacklistedIPCount              int         `json:"BlacklistedIpCount"`
	MemoryUsage                     int         `json:"MemoryUsage"`
	PhysicalMemoryUsage             int         `json:"PhysicalMemoryUsage"`
	FreeFirtualMemory               int64       `json:"FreeFirtualMemory"`
	TotalVirtualMemory              int64       `json:"TotalVirtualMemory"`
	FreePhysicalMemory              int64       `json:"FreePhysicalMemory"`
	TotalPhysicalMemory             int64       `json:"TotalPhysicalMemory"`
	DiskUsage                       int         `json:"DiskUsage"`
	FreeDiskSpace                   int64       `json:"FreeDiskSpace"`
	TotalDiskSpace                  int64       `json:"TotalDiskSpace"`
	CPUUsage                        int         `json:"CpuUsage"`
	MaintenanceExpiresAt            *time.Time  `json:"MaintenanceExpiresAt"`
	Support                         bool        `json:"Support"`
	ExpirationDate                  interface{} `json:"ExpirationDate"`
	OutboundRules                   int         `json:"OutboundRules"`
	BackupScheduled                 bool        `json:"BackupScheduled"`
	LastBackupDateTime              *time.Time  `json:"LastBackupDateTime"`
	ResellerName                    string      `json:"ResellerName"`
	LicenseKey                      string      `json:"LicenseKey"`
	ProductCode                     string      `json:"ProductCode"`
}

// SystemStatus fetches the system status
func (api *API) SystemStatus() (SystemStatus, error) {
	response := SystemStatus{}
	return response, api.getResponse("SystemStatus", &response)
}
