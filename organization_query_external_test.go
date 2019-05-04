package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchRepositories(t *testing.T) {
	if !*externalTests {
		return
	}
	repos := GetReposForOrganization()
	assert.True(t, len(repos) > 0, "No repositories found in the organization")
	if len(repos) > 0 {
		assert.NotNil(t, repos[0].Name, "No name found")
	}
}

func TestFetchTeamsNonexistent(t *testing.T) {
	if !*externalTests {
		return
	}
	_, err := GetReposForTeam("VJEIMVSKLDJFIOJEFF")
	assert.NotNil(t, err)
}

func TestFetchTeams(t *testing.T) {
	if !*externalTests {
		return
	}
	repos, err := GetReposForTeam("All Teams")
	assert.Nil(t, err)
	assert.True(t, len(repos) > 0, "No repositories found for the team")
	if len(repos) > 0 {
		assert.NotNil(t, repos[0].Name, "No name found")
	}
}
