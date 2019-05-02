package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Setup
	build := exec.Command("go", "build")
	err := build.Run()
	if err != nil {
		panic(fmt.Sprintf("Unable to build project: %s", err.Error()))
	}
	configInit()

	code := m.Run()

	// Teardown
	os.Exit(code)
}

func TestProgramWithEnvironmentVariable(t *testing.T) {
	// Setup the program
	binaryName := "vows"
	dir, err := os.Getwd()
	vows := exec.Command(path.Join(dir, binaryName))
	vows.Env = append(os.Environ(), "VOWS_GITHUB_ORG=bluewasher")

	// Run and verify the output
	output, err := vows.CombinedOutput()
	assert.Nil(t, err)
	assert.Contains(t, string(output), "bluewasher")
}

func TestProgramWithParameter(t *testing.T) {
	// Setup the program
	binaryName := "vows"
	dir, err := os.Getwd()
	vows := exec.Command(path.Join(dir, binaryName), "--github_org=redslide")

	// Run and verify the output
	output, err := vows.CombinedOutput()
	assert.Nil(t, err)
	assert.Contains(t, string(output), "redslide")
}

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
