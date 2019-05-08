package main

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestGetConfigValueWithNoConfiguration(t *testing.T) {
	value := getConfigValue("blah")
	assert.Empty(t, value)
}

func TestfetchAccessToken(t *testing.T) {
	holder := viper.Get("AccessToken")

	viper.Set("AccessToken", "water")
	token := fetchAccessToken()
	assert.NotEmpty(t, token, "Unable to find the github token, have you set the proper environment variable?")
	assert.Equal(t, "water", token)

	viper.Set("AccessToken", holder)
}

func TestFetchOrganization(t *testing.T) {
	holder := viper.Get("Organization")

	viper.Set("Organization", "rock")
	organization := fetchOrganization()
	assert.NotEmpty(t, organization, "Unable to find the github organization, have you set the proper environment variable?")
	assert.Equal(t, "rock", organization)

	viper.Set("Organization", holder)
}

func TestFetchTestRepoID(t *testing.T) {
	holder := viper.Get("GITHUB_TEST_REPOSITORY_ID")

	viper.Set("GITHUB_TEST_REPOSITORY_ID", "fire")
	repoID := fetchTestRepositoryID()
	assert.NotEmpty(t, repoID, "Unable to find the github test repo id, have you set the proper environment variable?")
	assert.Equal(t, "fire", repoID)

	viper.Set("GITHUB_TEST_REPOSITORY_ID", holder)
}

func TestTeams(t *testing.T) {
	// TODO: Have proper setup for this instead of depending on config file
	teams := fetchTeams()
	assert.Equal(t, 2, len(teams))
	assert.Equal(t, "all-teams", teams[0].Name)
	assert.Equal(t, push, teams[0].Permission)
}
