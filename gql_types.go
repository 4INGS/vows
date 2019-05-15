package main

import "github.com/shurcooL/githubv4"

// TeamQuery gets the team from GitHub
type TeamQuery struct {
	Organization struct {
		Name     string
		Location string
		URL      string
		Teams    struct {
			Nodes      []Team
			TotalCount int
		} `graphql:"teams(first:100,query:$teamname)"`
	} `graphql:"organization(login: $login)"`
}

// Team represents a Github Team
type Team struct {
	Name         string
	Repositories struct {
		Nodes    []Repository
		PageInfo struct {
			EndCursor   githubv4.String
			HasNextPage bool
		}
	} `graphql:"repositories(first:100, after:$repoCursor)"`
}

// TeamAccessQuery checks if a team has permission to a repo
type TeamAccessQuery struct {
	Organization struct {
		Name  string
		Teams struct {
			TotalCount int
			Nodes      []struct {
				Name         string
				Repositories struct {
					TotalCount int
					Edges      []struct {
						Permission string
						Node       struct {
							Name string
						}
					}
				} `graphql:"repositories(first:10, query: $repository)"`
			}
			PageInfo struct {
				HasNextPage bool
			}
		} `graphql:"teams(first:10, query: $team)"`
	} `graphql:"organization(login: $organization)"`
}

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
		} `graphql:"repositories(first:100, after:$repoCursor)"`
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
	ID                           string
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
		ClientMutationID     string
		BranchProtectionRule BranchProtectionRule
	} `graphql:"createBranchProtectionRule(input: $input)"`
}

// UpdateBranchProtectionRuleMutation will update the rule in Github
type UpdateBranchProtectionRuleMutation struct {
	UpdateBranchProtectionRule struct {
		ClientMutationID string
	} `graphql:"updateBranchProtectionRule(input: $input)"`
}
