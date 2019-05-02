package main

import (
	"bytes"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestProgramWithConfigfile(t *testing.T) {
	var jsonExample = []byte(`{"car": "hatchback"}`)
	viper.SetConfigType("json")
	viper.ReadConfig(bytes.NewBuffer(jsonExample))

	// Run
	organization, err := getConfigValue("car")

	// Verify
	assert.NotEmpty(t, organization, "Unable to find the github organization, have you set the proper environment variable?")
	assert.Nil(t, err)
	assert.Equal(t, "hatchback", organization)
}

func TestGetConfigValueWithNoConfiguration(t *testing.T) {
	value, err := getConfigValue("blah")
	assert.Empty(t, value)
	assert.NotNil(t, err)
}

func TestFetchGithubToken(t *testing.T) {
	viper.Set(GithubToken, "water")
	token, err := fetchGithubToken()
	assert.NotEmpty(t, token, "Unable to find the github token, have you set the proper environment variable?")
	assert.Nil(t, err)
	assert.Equal(t, "water", token)
}

func TestFetchOrganization(t *testing.T) {
	viper.Set(GithubOrganization, "rock")
	organization, err := fetchOrganization()
	assert.NotEmpty(t, organization, "Unable to find the github organization, have you set the proper environment variable?")
	assert.Nil(t, err)
	assert.Equal(t, "rock", organization)
}

func TestFetchTestRepoID(t *testing.T) {
	viper.Set(GithubTestRepositoryID, "fire")
	repoID, err := fetchTestRepositoryID()
	assert.NotEmpty(t, repoID, "Unable to find the github test repo id, have you set the proper environment variable?")
	assert.Nil(t, err)
	assert.Equal(t, "fire", repoID)
}
