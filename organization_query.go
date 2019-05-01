package main

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"
)

func runOrganizationQuery() []Repository {
	var query OrganizationQuery
	client := buildClient()
	vars := buildOrgVariables()
	return executeQuery(client, &query, vars)
}

func buildOrgVariables() map[string]interface{} {
	org, err := fetchOrganization()
	if err != nil {
		panic(fmt.Sprintf("Unable to build organization variables: %s", err.Error()))
	}
	return map[string]interface{}{
		"login":      githubv4.String(org),
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
