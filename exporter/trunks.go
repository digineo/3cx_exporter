package exporter

// Trunk represents a SIP trunk
type Trunk struct {
	ID             string `json:"Id"`
	Number         string `json:"Number"`
	Name           string `json:"Name"`
	Host           string `json:"Host"`
	Type           string `json:"Type"`
	SimCalls       string `json:"SimCalls"`
	ExternalNumber string `json:"ExternalNumber"`
	IsRegistered   bool   `json:"IsRegistered"`
	CanBeDeleted   bool   `json:"CanBeDeleted"`

	// empty strings for time values result in parse errors
	// RegisterOkTime     *time.Time
	// RegisterSentTime   *time.Time
	// RegisterFailedTime *time.Time
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
