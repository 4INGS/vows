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
func TestFetchTeams(t *testing.T) {
	if !*externalTests {
		return
	}
	_, err := GetReposForTeam("All Teams")
	assert.Nil(t, err)
}

func TestTeamPermission(t *testing.T) {
	if !*externalTests {
		return
	}
	var g GithubRepoHost
	access, err := g.TeamAccessToRepo(fetchTestTeamName(), fetchTestRepository())
	assert.Nil(t, err)
	assert.Equal(t, "ADMIN", access)
}

// Only keeping happy path external tests.
// Leaving these commented in case we want to re-enable in the futre.

// func TestFetchTeamsNonexistent(t *testing.T) {
// 	if !*externalTests {
// 		return
// 	}
// 	_, err := GetReposForTeam("VJEIMVSKLDJFIOJEFF")
// 	assert.NotNil(t, err)
// }

// func TestTeamPermissionIncorrectTeam(t *testing.T) {
// 	if !*externalTests {
// 		return
// 	}
// 	var g GithubRepoHost
// 	access, err := g.TeamAccessToRepo("ogres", fetchTestRepository())
// 	assert.NotNil(t, err)
// 	assert.Equal(t, "", access)
// }

// func TestTeamPermissionIncorrectRepo(t *testing.T) {
// 	if !*externalTests {
// 		return
// 	}
// 	var g GithubRepoHost
// 	access, err := g.TeamAccessToRepo(fetchTestTeamName(), "bad-code-repo")
// 	assert.Nil(t, err)
// 	assert.Equal(t, "", access)
// }
