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
	vows := exec.Command(path.Join(dir, binaryName), "--debug=true", "--preview=true")
	vows.Env = append(os.Environ(), "VOWS_GITHUB_ORG=bluewasher", "VOWS_GITHUB_TOKEN=12345678901234567890")

	// Run and verify the output
	output, err := vows.CombinedOutput()
	assert.Nil(t, err)
	assert.Contains(t, string(output), "bluewasher")
}

func TestProgramWithParameter(t *testing.T) {
	if !*integrationTests {
		return
	}
	// Setup the program
	binaryName := "vows"
	dir, err := os.Getwd()
	vows := exec.Command(path.Join(dir, binaryName), "--github_org=redslide", "--debug=true", "--preview=true", "--github_token=12345")

	// Run and verify the output
	output, err := vows.CombinedOutput()
	assert.Nil(t, err)
	assert.Contains(t, string(output), "redslide")
}
