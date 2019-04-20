package main

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"
)

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
	DismissesStaleReviews        bool
	isAdminEnforced              bool
	requiredApprovingReviewCount int
	requiresStrictStatusChecks   bool
}

func runOrganizationQuery() []Repository {
	var query OrganizationQuery
	client := buildClient()
	vars := buildOrgVariables()
	return executeQuery(client, &query, vars)
}

func buildOrgVariables() map[string]interface{} {
	return map[string]interface{}{
		"login":      githubv4.String("RepoFetch"),
		"repoCursor": (*githubv4.String)(nil),
	}
}

func executeQuery(client *githubv4.Client, query *OrganizationQuery, variables map[string]interface{}) []Repository {
	var allRepos []Repository
	for {
		err := client.Query(context.Background(), &query, variables)
		if err != nil {
			fmt.Println("Unable to run the query", err)
		}
		allRepos = append(allRepos, query.Organization.Repositories.Nodes...)
		if !query.Organization.Repositories.PageInfo.HasNextPage {
			break
		}
		variables["repoCursor"] = githubv4.NewString(query.Organization.Repositories.PageInfo.EndCursor)
	}
	return allRepos
}
