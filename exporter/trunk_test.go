package exporter

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTrunks(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	data, err := os.ReadFile("fixtures/TrunkList.json")
	require.NoError(err)

	response := struct {
		List []Trunk `json:"list"`
	}{}

	require.NoError(json.Unmarshal(data, &response))
	require.Len(response.List, 1)

	trunk := response.List[0]
	assert.Equal("SIP-Provider", trunk.Name)
	assert.Equal(true, trunk.IsRegistered)
}
