package main

import (
	"context"

	"github.com/shurcooL/githubv4"
)

// ApplyBranchProtectionMutation will call the Github mutation to add branch protections
func ApplyBranchProtectionMutation(id string) (string, error) {

	input := githubv4.CreateBranchProtectionRuleInput{
		RepositoryID:                 id,
		Pattern:                      "master",
		DismissesStaleReviews:        githubv4.NewBoolean(true),
		IsAdminEnforced:              githubv4.NewBoolean(true),
		RequiresApprovingReviews:     githubv4.NewBoolean(true),
		RequiredApprovingReviewCount: githubv4.NewInt(4),
		RequiresStatusChecks:         githubv4.NewBoolean(true),
		RequiredStatusCheckContexts: &[]githubv4.String{
			*githubv4.NewString("build"),
		},
	}

	var m CreateRuleMutation
	client := buildClient()
	err := client.Mutate(context.Background(), &m, input, nil)
	return m.CreateBranchProtectionRule.ClientMutationID, err
}
