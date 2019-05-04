package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildVariables(t *testing.T) {
	holderOrg, _ := fetchOrganization()
	setConfigValue(GithubOrganization, "blah")

	vars := buildOrgVariables()
	assert.NotNil(t, vars, "Variables not created")
	assert.Contains(t, vars, "login")

	setConfigValue(GithubOrganization, holderOrg)
}
