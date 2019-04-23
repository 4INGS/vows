// +build integration

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchRepositories(t *testing.T) {
	repos := GetReposHelper()
	//assert.Equal(t, "RepoFetch", oq.Organization.Name)
	assert.Len(t, repos, 1, "Incorrect number of repos found")
	//assert.True(t, len(repos) > 200, "Repo count not large enough, only at ", len(repos))
}

func TestFetchBranchProtection(t *testing.T) {
	repos := GetReposHelper()
	assert.NotNil(t, repos[0].Name, "No name found")
}

var cachedRepos []Repository

// Helper function to fetch the repos once, but allow multiple tests against results
func GetReposHelper() []Repository {
	if cachedRepos == nil {
		cachedRepos = runOrganizationQuery()
	}
	return cachedRepos
}
