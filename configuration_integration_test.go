package main

import (
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProgramWithEnvironmentVariable(t *testing.T) {
	if !*integrationTests {
		return
	}
	// Setup the program
	binaryName := "vows"
	dir, err := os.Getwd()
	vows := exec.Command(path.Join(dir, binaryName), "--debug", "--preview")
	vows.Env = append(os.Environ(), "VOWS_ACCESSTOKEN=08642")

	// Run and verify the output
	output, _ := vows.CombinedOutput()
	assert.Nil(t, err)
	assert.Contains(t, string(output), "08642")
}

func TestProgramWithParameter(t *testing.T) {
	if !*integrationTests {
		return
	}
	// Setup the program
	binaryName := "vows"
	dir, err := os.Getwd()
	vows := exec.Command(path.Join(dir, binaryName), "--debug", "--preview", "--accesstoken=97531")

	// Run and verify the output
	output, _ := vows.CombinedOutput()
	assert.Nil(t, err)
	assert.Contains(t, string(output), "97531")
}

func TestProgramWithConfigFile(t *testing.T) {
	if !*integrationTests {
		return
	}
	// Testing results of configInit(), which is part of every test setup

	// Verify
	token := fetchAccessToken()
	assert.NotEmpty(t, token, "No github token found in configuration")
	repoID := fetchTestRepositoryID()
	assert.NotEmpty(t, repoID, "No github org found in configuration")
}

// TODO: Need to inject configuration for this
func TestIgnoreRepos(t *testing.T) {
	if !*integrationTests {
		return
	}
	repos := fetchIgnoreRepositories()
	assert.True(t, len(repos) > 0)
}
