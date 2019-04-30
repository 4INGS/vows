// +build integration

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchRepositories(t *testing.T) {
	repos := GetReposHelper()
	assert.True(t, len(repos) > 0, "No repositories found in the organization")
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
