package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/shurcooL/githubv4"
)

// GithubRepoHost applies and updates branch protections
type GithubRepoHost struct {
	remoteteams []*github.Team
}

// AddBranchProtection will call the Github mutation to add branch protections
func (p GithubRepoHost) AddBranchProtection(repoID string) (BranchProtectionRule, error) {
	if isDebug() {
		fmt.Printf("Adding branch protection on %s\n", repoID)
	}

	rules := fetchBranchProtectionRules()
	input := githubv4.CreateBranchProtectionRuleInput{
		RepositoryID:                 repoID,
		Pattern:                      *githubv4.NewString(githubv4.String(rules.Pattern)),
		DismissesStaleReviews:        githubv4.NewBoolean(githubv4.Boolean(rules.DismissesStaleReviews)),
		IsAdminEnforced:              githubv4.NewBoolean(githubv4.Boolean(rules.IsAdminEnforced)),
		RequiresApprovingReviews:     githubv4.NewBoolean(githubv4.Boolean(rules.RequiresApprovingReviews)),
		RequiredApprovingReviewCount: githubv4.NewInt(githubv4.Int(rules.RequiredApprovingReviewCount)),
		RequiresStatusChecks:         githubv4.NewBoolean(githubv4.Boolean(rules.RequiresStatusChecks)),
	}

	checks := make([]githubv4.String, len(rules.RequiredStatusCheckContexts))
	for i, name := range rules.RequiredStatusCheckContexts {
		checks[i] = *githubv4.NewString(githubv4.String(name))
	}
	input.RequiredStatusCheckContexts = &checks

	var m CreateRuleMutation
	client := buildClient()
	err := client.Mutate(context.Background(), &m, input, nil)
	return m.CreateBranchProtectionRule.BranchProtectionRule, err
}

// UpdateBranchProtection will correct the settings on a branch protection
func (p GithubRepoHost) UpdateBranchProtection(repoID string, rule BranchProtectionRule) error {
	if isDebug() {
		fmt.Printf("Updating branch protection on %s\n", repoID)
	}

	rules := fetchBranchProtectionRules()
	input := githubv4.UpdateBranchProtectionRuleInput{
		BranchProtectionRuleID:       rule.ID,
		Pattern:                      githubv4.NewString(githubv4.String(rules.Pattern)),
		DismissesStaleReviews:        githubv4.NewBoolean(githubv4.Boolean(rules.DismissesStaleReviews)),
		IsAdminEnforced:              githubv4.NewBoolean(githubv4.Boolean(rules.IsAdminEnforced)),
		RequiresApprovingReviews:     githubv4.NewBoolean(githubv4.Boolean(rules.RequiresApprovingReviews)),
		RequiredApprovingReviewCount: githubv4.NewInt(githubv4.Int(rules.RequiredApprovingReviewCount)),
		RequiresStatusChecks:         githubv4.NewBoolean(githubv4.Boolean(rules.RequiresStatusChecks)),
		RequiredStatusCheckContexts: &[]githubv4.String{
			*githubv4.NewString("build"),
		},
	}

	var m UpdateBranchProtectionRuleMutation
	client := buildClient()
	err := client.Mutate(context.Background(), &m, input, nil)
	return err
}
