package main

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestGetConfigValueWithNoConfiguration(t *testing.T) {
	value, err := getConfigValue("blah")
	assert.Empty(t, value)
	assert.NotNil(t, err)
}

func TestFetchGithubToken(t *testing.T) {
	holder := viper.Get(GithubToken)

	viper.Set(GithubToken, "water")
	token, err := fetchGithubToken()
	assert.NotEmpty(t, token, "Unable to find the github token, have you set the proper environment variable?")
	assert.Nil(t, err)
	assert.Equal(t, "water", token)

	viper.Set(GithubToken, holder)
}

func TestFetchOrganization(t *testing.T) {
	holder := viper.Get(GithubOrganization)

	viper.Set(GithubOrganization, "rock")
	organization, err := fetchOrganization()
	assert.NotEmpty(t, organization, "Unable to find the github organization, have you set the proper environment variable?")
	assert.Nil(t, err)
	assert.Equal(t, "rock", organization)

	viper.Set(GithubOrganization, holder)
}

func TestFetchTestRepoID(t *testing.T) {
	holder := viper.Get(GithubTestRepositoryID)

	viper.Set(GithubTestRepositoryID, "fire")
	repoID, err := fetchTestRepositoryID()
	assert.NotEmpty(t, repoID, "Unable to find the github test repo id, have you set the proper environment variable?")
	assert.Nil(t, err)
	assert.Equal(t, "fire", repoID)

	viper.Set(GithubTestRepositoryID, holder)
}
