package exporter

// ServiceList represents the ServiceList response
type ServiceList []Service

// Service represents an item of the ServiceList response
type Service struct {
	Name             string `json:"Name"`
	DisplayName      string `json:"DisplayName"`
	Status           int    `json:"Status"`
	MemoryUsed       int    `json:"MemoryUsed"`
	CPUUsage         int    `json:"CpuUsage"`
	StartStopEnabled bool   `json:"startStopEnabled"`
	RestartEnabled   bool   `json:"restartEnabled"`
}

// ServiceList fetches the service list
func (api *API) ServiceList() (ServiceList, error) {
	response := ServiceList{}
	return response, api.getResponse("ServiceList", &response)
}
