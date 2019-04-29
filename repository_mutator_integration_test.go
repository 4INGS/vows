// +build integration

package main

import (
	"context"
	"testing"

	"github.com/shurcooL/githubv4"

	"github.com/stretchr/testify/assert"
)

func TestAddBranchProtectionMutationInvalidID(t *testing.T) {
	_, err := AddBranchProtectionMutation("123")
	assert.NotNil(t, err)
}

func TestAddBranchProtectionMutationValidID(t *testing.T) {
	// Setup: This test requires a Repository ID to run.
	repoID, err := fetchTestRepositoryID()
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	// Add protection
	rule, err2 := AddBranchProtectionMutation(repoID)
	assert.Nil(t, err2)
	assert.NotEmpty(t, rule.ID, "No branch rule ID returned from the mutation")

	// Clean up
	err3 := helperRemoveBranchProtectionRule(rule.ID)
	assert.Nil(t, err3, "Unable to remove the added branch protection.  Please remove manually")
}

func TestUpdateBranchProtectionMutationValidID(t *testing.T) {
	// Setup: This test requires a Repository ID to run.
	repoID, err := fetchTestRepositoryID()
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	// Add protection
	rule, err2 := AddBranchProtectionMutation(repoID)
	assert.Nil(t, err2)
	assert.NotEmpty(t, rule.ID, "No branch rule ID returned from the mutation")

	// Update protection
	err3 := UpdateBranchProtectionMutation(repoID, rule)
	assert.Nil(t, err3)

	// Clean up
	err4 := helperRemoveBranchProtectionRule(rule.ID)
	assert.Nil(t, err4, "Unable to remove the added branch protection.  Please remove manually")
}

// This will attempt to clean up after a test
// Helper function only for now.
type DeleteBranchProtectionRuleMutation struct {
	DeleteBranchProtectionRule struct {
		ClientMutationID string
	} `graphql:"deleteBranchProtectionRule(input: $input)"`
}

func helperRemoveBranchProtectionRule(ruleID string) error {
	input := githubv4.DeleteBranchProtectionRuleInput{
		BranchProtectionRuleID: ruleID,
	}
	var m DeleteBranchProtectionRuleMutation
	client := buildClient()
	err := client.Mutate(context.Background(), &m, input, nil)
	return err
}
