package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildClient(t *testing.T) {
	client := buildClient()
	assert.NotNil(t, client, "Client could not be created")
}

func TestFetchGithubToken(t *testing.T) {
	token := fetchGithubToken()
	assert.NotNil(t, token, "Unable to find the github token")
	fmt.Println("Token: ", token)
}

func TestBuildVariables(t *testing.T) {
	vars := buildVariables()
	assert.NotNil(t, vars, "Variables not created")
	assert.Len(t, vars, 1, "Length of variables not correct")
}

func TestFetchRepositories(t *testing.T) {
	client := buildClient()
	variables := buildVariables()
	oq := runQuery(client, variables)
	//fmt.Println("Name: ", oq.Organization.URL)

	assert.Equal(t, "4INGS", oq.Organization.Name)
}
