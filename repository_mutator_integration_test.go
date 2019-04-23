// +build integration

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBranchProtectionMutationInvalidID(t *testing.T) {
	_, err := ApplyBranchProtectionMutation("123")
	assert.NotNil(t, err)
}

// func TestAddBranchProtectionMutationValidID(t *testing.T) {
// 	id, err := ApplyBranchProtectionMutation("")
// 	assert.Nil(t, err)
// 	assert.NotEmpty(t, id)
// }

// var cachedRepos []Repository

// // Helper function to fetch the repos once, but allow multiple tests against results
// func BuildBranchProtectionHelper() {
// 	if cachedRepos == nil {
// 		cachedRepos = runOrganizationQuery()
// 	}
// 	return cachedRepos
// }
