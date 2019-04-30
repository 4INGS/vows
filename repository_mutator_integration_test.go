// +build integration

package main

import (
	"context"
	"testing"

	"github.com/shurcooL/githubv4"

	"github.com/stretchr/testify/assert"
)

func TestAddBranchProtectionInvalidID(t *testing.T) {
	var gp GithubProtector
	_, err := gp.AddBranchProtection("123")
	assert.NotNil(t, err)
}

func TestAddBranchProtectionValidID(t *testing.T) {
	// Setup: This test requires a Repository ID to run.
	repoID, err := fetchTestRepositoryID()
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	// Add protection
	var gp GithubProtector
	rule, err2 := gp.AddBranchProtection(repoID)
	assert.Nil(t, err2)
	assert.NotEmpty(t, rule.ID, "No branch rule ID returned from the mutation")

	// Clean up
	err3 := helperRemoveBranchProtectionRule(rule.ID)
	assert.Nil(t, err3, "Unable to remove the added branch protection.  Please remove manually")
}

func TestUpdateBranchProtectionValidID(t *testing.T) {
	// Setup: This test requires a Repository ID to run.
	repoID, err := fetchTestRepositoryID()
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	// Add protection
	var gp GithubProtector
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
