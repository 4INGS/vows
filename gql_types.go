package main

import "github.com/shurcooL/githubv4"

// OrganizationQuery queries github for things
type OrganizationQuery struct {
	Organization struct {
		Name         string
		Location     string
		URL          string
		Repositories struct {
			Nodes    []Repository
			PageInfo struct {
				EndCursor   githubv4.String
				HasNextPage bool
			}
		} `graphql:"repositories(first:10, after:$repoCursor)"`
	} `graphql:"organization(login: $login)"`
}

// Repository represents a Github repo
type Repository struct {
	ID                    string
	Name                  string
	BranchProtectionRules struct {
		Nodes []BranchProtectionRule
	} `graphql:"branchProtectionRules(first:10)"`
}

// BranchProtectionRule is the branch rules applied to a repository
type BranchProtectionRule struct {
	Pattern                      string
	RequiresStatusChecks         bool
	RequiresApprovingReviews     bool
	RequiredApprovingReviewCount int
	DismissesStaleReviews        bool
	IsAdminEnforced              bool
	RequiresStrictStatusChecks   bool
}

// CreateRuleMutation will create a branch protection rule in Github
type CreateRuleMutation struct {
	CreateBranchProtectionRule struct {
		ClientMutationID string
	} `graphql:"createBranchProtectionRule(input: $input)"`
}
