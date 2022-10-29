package exporter

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStatus(t *testing.T) {
	require := require.New(t)

	data, err := os.ReadFile("fixtures/SystemStatus.json")
	require.NoError(err)

	status := SystemStatus{}
	require.NoError(json.Unmarshal(data, &status))
}
