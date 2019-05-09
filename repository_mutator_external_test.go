package main

import (
	"context"
	"testing"

	"github.com/shurcooL/githubv4"

	"github.com/stretchr/testify/assert"
)

func TestAddBranchProtectionInvalidID(t *testing.T) {
	if !*externalTests {
		return
	}
	var gp GithubRepoHost
	_, err := gp.AddBranchProtection("123")
	assert.NotNil(t, err)
}

func TestAddBranchProtectionValidID(t *testing.T) {
	// Setup: This test requires a Repository ID to run.
	if !*externalTests {
		return
	}
	repoID := fetchTestRepositoryID()
	if len(repoID) == 0 {
		assert.Fail(t, "No test repository provided in the configuration")
		return
	}

	// Add protection
	var gp GithubRepoHost
	rule, err2 := gp.AddBranchProtection(repoID)
	assert.Nil(t, err2)
	assert.NotEmpty(t, rule.ID, "No branch rule ID returned from the mutation")

	// Clean up
	err3 := helperRemoveBranchProtectionRule(rule.ID)
	assert.Nil(t, err3, "Unable to remove the added branch protection.  Please remove manually")
}

func TestUpdateBranchProtectionValidID(t *testing.T) {
	// Setup: This test requires a Repository ID to run.
	if !*externalTests {
		return
	}
	repoID := fetchTestRepositoryID()
	if len(repoID) == 0 {
		assert.Fail(t, "No test repository provided in the configuration")
		return
	}

	// Add protection
	var gp GithubRepoHost
	rule, err2 := gp.AddBranchProtection(repoID)
	assert.Nil(t, err2)
	assert.NotEmpty(t, rule.ID, "No branch rule ID returned from the mutation")

	// Update protection
	err3 := gp.UpdateBranchProtection(repoID, rule)
	assert.Nil(t, err3)

	// Clean up
	err4 := helperRemoveBranchProtectionRule(rule.ID)
	assert.Nil(t, err4, "Unable to remove the added branch protection.  Please remove manually")
}

func TestAddTeamToRepo(t *testing.T) {
	if !*externalTests {
		//return
	}
	var gp GithubRepoHost
	var tc teamConfig
	tc.Name = fetchTestTeamName()
	tc.Permission = admin
	teamID, err := gp.GetTeamID(tc.Name)
	assert.Nil(t, err)
	tc.ID = teamID

	err = gp.AddTeamToRepo(&tc, fetchTestRepository())
	assert.Nil(t, err)
}

func TestGetTeamID(t *testing.T) {
	if !*externalTests {
		return
	}

	var gp GithubRepoHost
	teamID, err := gp.GetTeamID("All Teams")

	assert.Nil(t, err)
	assert.True(t, teamID > 0)
}

// This will attempt to clean up after a test
// Helper function only for now.
type DeleteBranchProtectionRule struct {
	DeleteBranchProtectionRule struct {
		ClientMutationID string
	} `graphql:"deleteBranchProtectionRule(input: $input)"`
}

func helperRemoveBranchProtectionRule(ruleID string) error {
	input := githubv4.DeleteBranchProtectionRuleInput{
		BranchProtectionRuleID: ruleID,
	}
	var m DeleteBranchProtectionRule
	client := buildClient()
	err := client.Mutate(context.Background(), &m, input, nil)
	return err
}
