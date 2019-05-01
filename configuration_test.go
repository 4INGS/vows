package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfigValueWithNoConfiguration(t *testing.T) {
	value, err := getConfigValue("blah")
	assert.Empty(t, value)
	assert.NotNil(t, err)
}

func TestFetchGithubToken(t *testing.T) {
	configInit()
	token, err := fetchGithubToken()
	assert.NotEmpty(t, token, "Unable to find the github token, have you set the proper environment variable?")
	assert.Nil(t, err)
}

func TestFetchOrganization(t *testing.T) {
	configInit()
	organization, err := fetchOrganization()
	assert.NotEmpty(t, organization, "Unable to find the github organization, have you set the proper environment variable?")
	assert.Nil(t, err)
}
