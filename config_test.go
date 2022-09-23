package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	assert := assert.New(t)
	os.Setenv("ENV", "DEV")

	config, err := parseConfig()

	assert.NoError(err)
	assert.Equal("localhost", config.Host)
	assert.Equal("admin", config.Username)
	assert.Equal("admin", config.Password)
}
