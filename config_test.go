package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	assert := assert.New(t)

	config, err := parseConfig("fixtures/config.json")

	assert.NoError(err)
	assert.Equal("3cx.example.com", config.Hostname)
	assert.Equal("admin", config.Username)
	assert.Equal("secret", config.Password)
}
