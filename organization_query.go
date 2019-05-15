package main

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"
)

// RepoNameSorter will order repositories by name
type RepoNameSorter []Repository

func (a RepoNameSorter) Len() int           { return len(a) }
func (a RepoNameSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a RepoNameSorter) Less(i, j int) bool { return a[i].Name < a[j].Name }

// GetReposForOrganization will fetch all repositories and their branch protections for an organization
func GetReposForOrganization() []Repository {
	var query OrganizationQuery
	client := buildClient()
	vars := buildOrgVariables()
	vars["repoCursor"] = (*githubv4.String)(nil)
	return executeOrganizationQuery(client, &query, vars)
}

// TeamAccessToRepo checks if the given teamname has the given permission to the repo specified
func (p GithubRepoHost) TeamAccessToRepo(teamname string, repo string) (string, error) {
	if isDebug() {
		fmt.Printf("Checking if %s has access to %s", teamname, repo)
	}

	var query TeamAccessQuery
	client := buildClient()
	vars := buildTeamVariables(teamname, repo)
	err := client.Query(context.Background(), &query, vars)
	if err != nil {
		fmt.Println("Unable to run the query", err)
	}

	// See if the permission node exists and matches our requested permission
	if query.Organization.Teams.TotalCount != 1 {
		return "", fmt.Errorf("No teamname found named %s in the organization %s", teamname, fetchOrganization())
	}
	teamObj := query.Organization.Teams.Nodes[0]
	if teamObj.Repositories.TotalCount == 0 {
		fmt.Println("NO teamname access")
		// This is fine, the teamname just doesn't have access
		return "", nil
	}
	for _, edge := range query.Organization.Teams.Nodes[0].Repositories.Edges {
		if edge.Node.Name == repo {
			return edge.Permission, nil
		}
	}
	return "", nil
}

func executeOrganizationQuery(client *githubv4.Client, query *OrganizationQuery, variables map[string]interface{}) []Repository {
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
	return RepoNameSorter(allRepos)
}

// GetReposForTeam for fetch all repositories a teamname has access to
func GetReposForTeam(teamname string) ([]Repository, error) {
	var query TeamQuery
	client := buildClient()
	vars := buildOrgVariables()
	vars["teamname"] = githubv4.String(teamname)
	vars["repoCursor"] = (*githubv4.String)(nil)
	return executeTeamQuery(client, &query, vars)
}

func executeTeamQuery(client *githubv4.Client, query *TeamQuery, variables map[string]interface{}) ([]Repository, error) {
	var allRepos []Repository
	for {
		err := client.Query(context.Background(), &query, variables)
		if err != nil {
			fmt.Println("Unable to run the query", err)
		}
		// If debug, log this out
		//fmt.Printf("query %+v\n", query)
		if query.Organization.Teams.TotalCount != 1 {
			return nil, fmt.Errorf("No teamname found with name %s", variables["teamname"])
		}
		allRepos = append(allRepos, query.Organization.Teams.Nodes[0].Repositories.Nodes...)
		if !query.Organization.Teams.Nodes[0].Repositories.PageInfo.HasNextPage {
			break
		}
		variables["repoCursor"] = githubv4.NewString(query.Organization.Teams.Nodes[0].Repositories.PageInfo.EndCursor)
	}
	return allRepos, nil
}

func buildOrgVariables() map[string]interface{} {
	org := fetchOrganization()
	if len(org) == 0 {
		panic(fmt.Sprintf("Unable to build organization variables, no organization found in the config"))
	}
	return map[string]interface{}{
		"login": githubv4.String(org),
	}

}

func buildTeamVariables(teamname string, repo string) map[string]interface{} {
	org := fetchOrganization()
	if len(org) == 0 {
		panic(fmt.Sprintf("Unable to build organization variables, no organization found in the config"))
	}
	return map[string]interface{}{
		"organization": githubv4.String(org),
		"team":         githubv4.String(teamname),
		"repository":   githubv4.String(repo),
	}
}
