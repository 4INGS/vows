package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildVariables(t *testing.T) {
	holderOrg := fetchOrganization()
	setConfigValue("Organization", "blah")

	vars := buildOrgVariables()
	assert.NotNil(t, vars, "Variables not created")
	assert.Contains(t, vars, "login")

	setConfigValue("Organization", holderOrg)
}

func TestRepoNameSort(t *testing.T) {
	repos := []Repository{
		{ID: "123", Name: "CCC"},
		{ID: "123", Name: "AAA"},
		{ID: "123", Name: "BBB"},
	}
	sort.Sort(RepoNameSorter(repos))
	assert.Equal(t, 3, len(repos), "Wrong length")
	assert.Equal(t, "AAA", repos[0].Name)
	assert.Equal(t, "BBB", repos[1].Name)
	assert.Equal(t, "CCC", repos[2].Name)
}
