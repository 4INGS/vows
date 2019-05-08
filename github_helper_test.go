package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildClient(t *testing.T) {
	holder := fetchAccessToken()
	setConfigValue("AccessToken", "blah")

	client := buildClient()
	assert.NotNil(t, client, "Client could not be created")

	setConfigValue("AccessToken", holder)
}
