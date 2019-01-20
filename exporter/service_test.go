package exporter

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServiceList(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	data, err := ioutil.ReadFile("fixtures/ServiceList.json")
	require.NoError(err)

	serviceList := ServiceList{}
	require.NoError(json.Unmarshal(data, &serviceList))
	require.Len(serviceList, 12)
	assert.Equal("3CXCfgServ01", serviceList[0].Name)
	assert.Equal(4, serviceList[0].Status)
}
