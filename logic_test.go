package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
)

type mockprotector struct {
	mock.Mock
}

func (m *mockprotector) applyBranchProtection(id string) {
	m.Called(id)
}

func TestNoProtector(t *testing.T) {
	// Setup
	repos := []Repository{
		Repository{
			ID: "123",
		},
	}
	// Execute
	err := ApplyBranchProtection(repos, nil, nil)
	// Verify
	assert.NotNil(t, err)
}

func TestSingleRepo(t *testing.T) {
	// Setup
	testObj := new(mockprotector)
	testObj.On("applyBranchProtection", mock.AnythingOfType("string")).Return()
	repos := []Repository{
		Repository{
			ID: "123",
		},
	}
	// Execute
	ApplyBranchProtection(repos, nil, testObj)
	// Verify
	testObj.AssertNumberOfCalls(t, "applyBranchProtection", 1)
}

func TestMultiRepo(t *testing.T) {
	// Setup
	testObj := new(mockprotector)
	testObj.On("applyBranchProtection", mock.AnythingOfType("string")).Return()
	repos := []Repository{
		Repository{
			ID: "123",
		},
		Repository{
			ID: "456",
		},
	}
	// Execute
	ApplyBranchProtection(repos, nil, testObj)
	// Verify
	testObj.AssertNumberOfCalls(t, "applyBranchProtection", 2)
}

func TestSkipWhiteList(t *testing.T) {
	// Setup
	testObj := new(mockprotector)
	testObj.On("applyBranchProtection", mock.AnythingOfType("string")).Return()
	repos := []Repository{
		Repository{
			ID: "123",
		},
	}
	whitelist := []string{"123"}
	// Execute
	ApplyBranchProtection(repos, whitelist, testObj)
	// Verify
	testObj.AssertNumberOfCalls(t, "applyBranchProtection", 0)
}

func TestCorrectBranchProtections(t *testing.T) {
	// Setup
	testObj := new(mockprotector)
	testObj.On("applyBranchProtection", mock.AnythingOfType("string")).Return()
	repos := mockCorrectRepos()
	// Execute
	ApplyBranchProtection(repos, nil, testObj)
	// Verify
	testObj.AssertNumberOfCalls(t, "applyBranchProtection", 0)
}

func TestIncorrectRequiresStatusChecks(t *testing.T) {
	// Setup
	testObj := new(mockprotector)
	testObj.On("applyBranchProtection", mock.AnythingOfType("string")).Return()
	repos := mockCorrectRepos()
	//Break a rule
	repos[0].BranchProtectionRules.Nodes[0].RequiresStatusChecks = false

	// Execute
	ApplyBranchProtection(repos, nil, testObj)
	// Verify
	testObj.AssertNumberOfCalls(t, "applyBranchProtection", 1)
}
func TestIncorrectIsAdminEnforced(t *testing.T) {
	// Setup
	testObj := new(mockprotector)
	testObj.On("applyBranchProtection", mock.AnythingOfType("string")).Return()
	repos := mockCorrectRepos()
	//Break a rule
	repos[0].BranchProtectionRules.Nodes[0].IsAdminEnforced = false

	// Execute
	ApplyBranchProtection(repos, nil, testObj)
	// Verify
	testObj.AssertNumberOfCalls(t, "applyBranchProtection", 1)
}
func TestIncorrectReviewCount(t *testing.T) {
	// Setup
	testObj := new(mockprotector)
	testObj.On("applyBranchProtection", mock.AnythingOfType("string")).Return()
	repos := mockCorrectRepos()
	//Break a rule
	repos[0].BranchProtectionRules.Nodes[0].RequiredApprovingReviewCount = 0

	// Execute
	ApplyBranchProtection(repos, nil, testObj)
	// Verify
	testObj.AssertNumberOfCalls(t, "applyBranchProtection", 1)
}

func mockCorrectRepos() []Repository {
	repos := []Repository{
		Repository{
			ID: "123",
			BranchProtectionRules: struct {
				Nodes []BranchProtectionRule
			}{
				Nodes: []BranchProtectionRule{
					BranchProtectionRule{
						Pattern:                      "master",
						RequiresStatusChecks:         true,
						RequiresApprovingReviews:     true,
						RequiredApprovingReviewCount: 1,
						DismissesStaleReviews:        true,
						IsAdminEnforced:              true,
						RequiresStrictStatusChecks:   true,
					},
				},
			},
		},
	}
	return repos
}
