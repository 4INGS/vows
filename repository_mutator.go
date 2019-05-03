package main

import (
	"context"

	"github.com/shurcooL/githubv4"
)

// GithubProtector applies and updates branch protections
type GithubProtector struct{}

// AddBranchProtection will call the Github mutation to add branch protections
func (p GithubProtector) AddBranchProtection(repoID string) (BranchProtectionRule, error) {

	// TODO: Allow this to be set in configuration
	input := githubv4.CreateBranchProtectionRuleInput{
		RepositoryID:                 repoID,
		Pattern:                      "master",
		DismissesStaleReviews:        githubv4.NewBoolean(true),
		IsAdminEnforced:              githubv4.NewBoolean(false),
		RequiresApprovingReviews:     githubv4.NewBoolean(true),
		RequiredApprovingReviewCount: githubv4.NewInt(1),
		RequiresStatusChecks:         githubv4.NewBoolean(true),
		RequiredStatusCheckContexts: &[]githubv4.String{
			*githubv4.NewString("build"),
		},
	}

	var m CreateRuleMutation
	client := buildClient()
	err := client.Mutate(context.Background(), &m, input, nil)
	return m.CreateBranchProtectionRule.BranchProtectionRule, err
}

// UpdateBranchProtection will correct the settings on a branch protection
func (p GithubProtector) UpdateBranchProtection(repoID string, rule BranchProtectionRule) error {

	// TODO: Allow this to be set in configuration
	input := githubv4.UpdateBranchProtectionRuleInput{
		BranchProtectionRuleID:       rule.ID,
		Pattern:                      githubv4.NewString("master"),
		DismissesStaleReviews:        githubv4.NewBoolean(true),
		IsAdminEnforced:              githubv4.NewBoolean(true),
		RequiresApprovingReviews:     githubv4.NewBoolean(true),
		RequiredApprovingReviewCount: githubv4.NewInt(1),
		RequiresStatusChecks:         githubv4.NewBoolean(true),
		RequiredStatusCheckContexts: &[]githubv4.String{
			*githubv4.NewString("build"),
		},
	}

	var m UpdateBranchProtectionRuleMutation
	client := buildClient()
	err := client.Mutate(context.Background(), &m, input, nil)
	return err
}
