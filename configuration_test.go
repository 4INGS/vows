package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchGithubToken(t *testing.T) {
	token := fetchGithubToken()
	assert.NotEmpty(t, token, "Unable to find the github token, have you set the proper environment variable?")
}

func TestFetchOrganization(t *testing.T) {
	organization := fetchOrganization()
	assert.NotEmpty(t, organization, "Unable to find the github organization, have you set the proper environment variable?")
}
