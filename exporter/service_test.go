package exporter

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceList(t *testing.T) {
	assert := assert.New(t)
	data, err := ioutil.ReadFile("fixtures/ServiceList.json")

	assert.NoError(err)

	serviceList := ServiceList{}
	assert.NoError(json.Unmarshal(data, &serviceList))
	assert.Len(serviceList, 12)
	assert.Equal("3CXCfgServ01", serviceList[0].Name)
	assert.Equal(4, serviceList[0].Status)
}
