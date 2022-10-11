package models

import "time"

type InstanceState struct {
	Id                  int        `json:"instance_state_id"`
	InstanceId          int        `json:"instance_id"`
	BlacklistSize       int        `json:"blacklist_size"`
	CallsActive         int        `json:"calls_active"`
	CallsLimit          int        `json:"calls_limit"`
	ExtensionsTotal     int        `json:"extensions_total"`
	ExtensionsRegistred int        `json:"extensions_registred"`
	LastBackUp          *time.Time `json:"last_backup"`
	MaintenceUntil      *time.Time `json:"maintence_until"`
	ServiceStatus       string     `json:"service_status"`
	ServiceCPU          string     `json:"service_cpu"`
	ServiceMemory       string     `json:"service_memory"`
	TruncRegistred      string     `json:"trunc_registred"`
	CreatedAt           time.Time  `json:"created_at"`
}
