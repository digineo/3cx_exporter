package exporter

import (
	"time"
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
}

// SystemStatus fetches the system status
func (api *API) SystemStatus() (SystemStatus, error) {
	response := SystemStatus{}
	return response, api.getResponse("SystemStatus", &response)
}
