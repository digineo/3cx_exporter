package exporter

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatus(t *testing.T) {
	assert := assert.New(t)
	data, err := ioutil.ReadFile("fixtures/SystemStatus.json")

	assert.NoError(err)

	status := SystemStatus{}
	assert.NoError(json.Unmarshal(data, &status))
}
