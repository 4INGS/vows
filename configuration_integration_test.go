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
	// Setup uses data from test configuration, loaded as part of test init

	// Verify
	token := fetchAccessToken()
	assert.NotEmpty(t, token, "No github token found in configuration")
	assert.Equal(t, "1111111111111111111111111111111111111111", token)
	repoID := fetchTestRepositoryID()
	assert.NotEmpty(t, repoID, "No github org found in configuration")
}

func TestIgnoreRepos(t *testing.T) {
	if !*integrationTests {
		return
	}
	// Setup uses data from test configuration, loaded as part of test init

	repos := fetchIgnoreRepositories()
	assert.True(t, len(repos) == 2)
	assert.Equal(t, "Repo1", repos[0])
}

func TestTeams(t *testing.T) {
	if !*integrationTests {
		return
	}
	// Setup uses data from test configuration, loaded as part of test init

	teams := fetchTeams()
	assert.Equal(t, 2, len(teams))
	assert.Equal(t, "team1", teams[0].Name)
	assert.Equal(t, pull, teams[0].Permission)
}
