package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
)

type mockprotector struct {
	mock.Mock
}

func (m *mockprotector) AddBranchProtection(id string) (BranchProtectionRule, error) {
	m.Called(id)
	var br BranchProtectionRule
	return br, nil
}
func (m *mockprotector) UpdateBranchProtection(id string, r BranchProtectionRule) error {
	m.Called(id)
	return nil
}

func TestNoProtector(t *testing.T) {
	// Setup
	repos := []Repository{
		Repository{
			ID: "834",
		},
	}
	var w ignorelist

	// Execute
	err := ApplyBranchProtection(repos, w, nil)
	// Verify
	assert.NotNil(t, err)
}

func TestSingleRepo(t *testing.T) {
	// Setup
	testObj := new(mockprotector)
	//testObj.On("AddBranchProtection", mock.AnythingOfType("string")).Return()
	testObj.On("AddBranchProtection", "456").Return()
	repos := []Repository{
		Repository{
			ID: "456",
		},
	}
	var w ignorelist

	// Execute
	ApplyBranchProtection(repos, w, testObj)
	// Verify
	testObj.AssertNumberOfCalls(t, "AddBranchProtection", 1)
}

func TestMultiRepo(t *testing.T) {
	// Setup
	testObj := new(mockprotector)
	testObj.On("AddBranchProtection", mock.AnythingOfType("string")).Return()
	repos := []Repository{
		Repository{
			ID: "345",
		},
		Repository{
			ID: "567",
		},
	}
	var w ignorelist
	// Execute
	ApplyBranchProtection(repos, w, testObj)
	// Verify
	testObj.AssertNumberOfCalls(t, "AddBranchProtection", 2)
}

func TestSkipignorelist(t *testing.T) {
	// Setup
	testObj := new(mockprotector)
	testObj.On("AddBranchProtection", mock.AnythingOfType("string")).Return()
	repos := []Repository{
		Repository{
			ID:   "123",
			Name: "abc",
		},
	}
	var w ignorelist
	w.SetLines([]string{"abc"})
	// Execute
	ApplyBranchProtection(repos, w, testObj)
	// Verify
	testObj.AssertNumberOfCalls(t, "AddBranchProtection", 0)
}

func TestCorrectBranchProtections(t *testing.T) {
	// Setup
	testObj := new(mockprotector)
	testObj.On("AddBranchProtection", mock.AnythingOfType("string")).Return()
	repos := mockCorrectRepos()
	var w ignorelist
	// Execute
	ApplyBranchProtection(repos, w, testObj)
	// Verify
	testObj.AssertNumberOfCalls(t, "AddBranchProtection", 0)
}

func TestIncorrectRequiresStatusChecks(t *testing.T) {
	// Setup
	testObj := new(mockprotector)
	testObj.On("UpdateBranchProtection", "2468").Return()
	repos := mockCorrectRepos()
	var w ignorelist
	//Break a rule
	repos[0].BranchProtectionRules.Nodes[0].RequiresStatusChecks = false

	// Execute
	ApplyBranchProtection(repos, w, testObj)
	// Verify
	testObj.AssertNumberOfCalls(t, "UpdateBranchProtection", 1)
}
func TestIncorrectIsAdminEnforced(t *testing.T) {
	// Setup
	testObj := new(mockprotector)
	testObj.On("UpdateBranchProtection", "2468").Return()
	repos := mockCorrectRepos()
	var w ignorelist
	//Break a rule
	repos[0].BranchProtectionRules.Nodes[0].IsAdminEnforced = false

	// Execute
	ApplyBranchProtection(repos, w, testObj)
	// Verify
	testObj.AssertNumberOfCalls(t, "UpdateBranchProtection", 1)
}
func TestIncorrectReviewCount(t *testing.T) {
	// Setup
	testObj := new(mockprotector)
	testObj.On("UpdateBranchProtection", "2468").Return()
	repos := mockCorrectRepos()
	var w ignorelist
	//Break a rule
	repos[0].BranchProtectionRules.Nodes[0].RequiredApprovingReviewCount = 0

	// Execute
	ApplyBranchProtection(repos, w, testObj)
	// Verify
	testObj.AssertNumberOfCalls(t, "UpdateBranchProtection", 1)
}

func TestPreviewModeAdd(t *testing.T) {
	// Setup
	setConfigValue("preview", "true")
	testObj := new(mockprotector)
	testObj.On("AddBranchProtection", "456").Return()
	repos := []Repository{
		Repository{
			ID:   "456",
			Name: "test-preview-repo",
		},
	}
	var w ignorelist
	// Execute
	ApplyBranchProtection(repos, w, testObj)
	// Verify
	testObj.AssertNumberOfCalls(t, "AddBranchProtection", 0)
}

func TestPreviewModeUpdate(t *testing.T) {
	// Setup
	setConfigValue("preview", "true")
	testObj := new(mockprotector)
	testObj.On("UpdateBranchProtection", "2468").Return()
	repos := mockCorrectRepos()
	//Break a rule
	repos[0].BranchProtectionRules.Nodes[0].RequiresStatusChecks = false
	var w ignorelist

	// Execute
	ApplyBranchProtection(repos, w, testObj)
	// Verify
	testObj.AssertNumberOfCalls(t, "UpdateBranchProtection", 0)
}

func mockCorrectRepos() []Repository {
	repos := []Repository{
		Repository{
			ID:   "2468",
			Name: "test-repo",
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
