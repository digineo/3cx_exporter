package exporter

// ServiceList represents the ServiceList response
type ServiceList []Service

// Service represents an item of the ServiceList response
type Service struct {
	Name       string
	Status     int
	MemoryUsed int
	CPUUsage   int
}

// ServiceList fetches the service list
func (api *API) ServiceList() (ServiceList, error) {
	response := ServiceList{}
	return response, api.getResponse("ServiceList", &response)
}
