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
func (m *mockprotector) AddTeamToRepo(teamID int64, repoName string) error {
	m.Called(teamID)
	return nil
}
func (m *mockprotector) GetTeamID(teamname string) (int64, error) {
	m.Called(teamname)
	return 1, nil
}

func TestNoProtector(t *testing.T) {
	// Setup
	repos := []Repository{
		Repository{
			ID: "834",
		},
	}
	var w Ignorelist

	// Execute
	err := ProcessRepositories(repos, w, nil, "tn")
	// Verify
	assert.NotNil(t, err)
}

func TestSingleRepo(t *testing.T) {
	// Setup
	testObj := new(mockprotector)
	//testObj.On("AddBranchProtection", mock.AnythingOfType("string")).Return()
	testObj.On("AddBranchProtection", "456").Return()
	testObj.On("GetTeamID", mock.AnythingOfType("string")).Return()
	testObj.On("AddTeamToRepo", mock.Anything, mock.Anything).Return()
	repos := []Repository{
		Repository{
			ID: "456",
		},
	}
	var w Ignorelist

	// Execute
	err := ProcessRepositories(repos, w, testObj, "tn")
	// Verify
	assert.Nil(t, err)
	testObj.AssertNumberOfCalls(t, "AddBranchProtection", 1)
}

func TestMultiRepo(t *testing.T) {
	// Setup
	testObj := new(mockprotector)
	testObj.On("AddBranchProtection", mock.AnythingOfType("string")).Return()
	testObj.On("GetTeamID", mock.AnythingOfType("string")).Return()
	testObj.On("AddTeamToRepo", mock.Anything, mock.Anything).Return()
	repos := []Repository{
		Repository{
			ID: "345",
		},
		Repository{
			ID: "567",
		},
	}
	var w Ignorelist
	// Execute
	ProcessRepositories(repos, w, testObj, "tn")
	// Verify
	testObj.AssertNumberOfCalls(t, "AddBranchProtection", 2)
}

func TestRepoOnIgnorelist(t *testing.T) {
	// Setup
	testObj := new(mockprotector)
	testObj.On("AddBranchProtection", mock.AnythingOfType("string")).Return()
	testObj.On("GetTeamID", mock.AnythingOfType("string")).Return()
	repos := []Repository{
		Repository{
			ID:   "123",
			Name: "abc",
		},
	}
	var w Ignorelist
	w.SetLines([]string{"abc"})
	// Execute
	ProcessRepositories(repos, w, testObj, "tn")
	// Verify
	testObj.AssertNumberOfCalls(t, "AddBranchProtection", 0)
}

func TestCorrectBranchProtections(t *testing.T) {
	// Setup
	testObj := new(mockprotector)
	testObj.On("AddBranchProtection", mock.AnythingOfType("string")).Return()
	testObj.On("GetTeamID", mock.AnythingOfType("string")).Return()
	testObj.On("AddTeamToRepo", mock.Anything, mock.Anything).Return()
	repos := mockCorrectRepos()
	var w Ignorelist
	// Execute
	ProcessRepositories(repos, w, testObj, "tn")
	// Verify
	testObj.AssertNumberOfCalls(t, "AddBranchProtection", 0)
}

func TestIncorrectRequiresStatusChecks(t *testing.T) {
	// Setup
	testObj := new(mockprotector)
	testObj.On("UpdateBranchProtection", "2468").Return()
	testObj.On("GetTeamID", mock.AnythingOfType("string")).Return()
	testObj.On("AddTeamToRepo", mock.Anything, mock.Anything).Return()
	repos := mockCorrectRepos()
	var w Ignorelist
	//Break a rule
	repos[0].BranchProtectionRules.Nodes[0].RequiresStatusChecks = false

	// Execute
	ProcessRepositories(repos, w, testObj, "tn")
	// Verify
	testObj.AssertNumberOfCalls(t, "UpdateBranchProtection", 1)
}
func TestIncorrectIsAdminEnforced(t *testing.T) {
	// Setup
	testObj := new(mockprotector)
	testObj.On("UpdateBranchProtection", "2468").Return()
	testObj.On("GetTeamID", mock.AnythingOfType("string")).Return()
	testObj.On("AddTeamToRepo", mock.Anything, mock.Anything).Return()
	repos := mockCorrectRepos()
	var w Ignorelist
	//Break a rule
	repos[0].BranchProtectionRules.Nodes[0].IsAdminEnforced = false

	// Execute
	ProcessRepositories(repos, w, testObj, "tn")
	// Verify
	testObj.AssertNumberOfCalls(t, "UpdateBranchProtection", 1)
}
func TestIncorrectReviewCount(t *testing.T) {
	// Setup
	testObj := new(mockprotector)
	testObj.On("UpdateBranchProtection", "2468").Return()
	testObj.On("GetTeamID", mock.AnythingOfType("string")).Return()
	testObj.On("AddTeamToRepo", mock.Anything, mock.Anything).Return()
	repos := mockCorrectRepos()
	var w Ignorelist
	//Break a rule
	repos[0].BranchProtectionRules.Nodes[0].RequiredApprovingReviewCount = 0

	// Execute
	ProcessRepositories(repos, w, testObj, "tn")
	// Verify
	testObj.AssertNumberOfCalls(t, "UpdateBranchProtection", 1)
}

func TestPreviewModeAdd(t *testing.T) {
	// Setup
	setConfigValue("preview", "true")
	testObj := new(mockprotector)
	testObj.On("AddBranchProtection", "456").Return()
	testObj.On("GetTeamID", mock.AnythingOfType("string")).Return()
	testObj.On("AddTeamToRepo", mock.Anything, mock.Anything).Return()
	repos := []Repository{
		Repository{
			ID:   "456",
			Name: "test-preview-repo",
		},
	}
	var w Ignorelist
	// Execute
	ProcessRepositories(repos, w, testObj, "tn")
	// Verify
	testObj.AssertNumberOfCalls(t, "AddBranchProtection", 0)
}

func TestPreviewModeUpdate(t *testing.T) {
	// Setup
	setConfigValue("preview", "true")
	testObj := new(mockprotector)
	testObj.On("UpdateBranchProtection", "2468").Return()
	testObj.On("GetTeamID", mock.AnythingOfType("string")).Return()
	testObj.On("AddTeamToRepo", mock.Anything, mock.Anything).Return()
	repos := mockCorrectRepos()
	//Break a rule
	repos[0].BranchProtectionRules.Nodes[0].RequiresStatusChecks = false
	var w Ignorelist

	// Execute
	ProcessRepositories(repos, w, testObj, "tn")
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
