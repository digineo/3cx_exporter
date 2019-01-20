package exporter

// Trunk represents a SIP trunk
type Trunk struct {
	Name         string
	IsRegistered bool
}

// TrunkList fetches the trunk list
func (api *API) TrunkList() ([]Trunk, error) {
	response := struct {
		List []Trunk `json:"list"`
	}{}

	err := api.getResponse("TrunkList", &response)
	if err != nil {
		return nil, err
	}

	return response.List, nil
}
