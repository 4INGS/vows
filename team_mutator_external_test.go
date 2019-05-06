package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTeamToRepo(t *testing.T) {
	if !*externalTests {
		return
	}
	err := AddTeamToRepo("All Teams", "fuzzy-octo-parakeet")
	assert.Nil(t, err)
}

func TestGetTeamID(t *testing.T) {
	if !*externalTests {
		return
	}
	client := getV3Client()
	org, _ := fetchOrganization()

	teamID, err := getTeamID("All Teams", client, org)

	assert.Nil(t, err)
	assert.True(t, teamID > 0)
}
