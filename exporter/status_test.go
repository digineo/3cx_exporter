package exporter

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStatus(t *testing.T) {
	require := require.New(t)

	data, err := ioutil.ReadFile("fixtures/SystemStatus.json")
	require.NoError(err)

	status := SystemStatus{}
	require.NoError(json.Unmarshal(data, &status))
}
