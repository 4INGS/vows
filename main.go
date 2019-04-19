package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

// OrganizationQuery queries github for things
type OrganizationQuery struct {
	Organization struct {
		Name     string
		Location string
		URL      string
	} `graphql:"organization(login: $login)"`
}

func main() {
	client := buildClient()
	variables := buildVariables()

	oq := runQuery(client, variables)
	printRepos(oq)
}

func buildClient() *githubv4.Client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: fetchGithubToken()},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)
	return client
}

func fetchGithubToken() string {
	return os.Getenv("GITHUB_TOKEN")
}

func buildVariables() map[string]interface{} {
	return map[string]interface{}{
		"login": githubv4.String("4ings"),
	}
}

func runQuery(client *githubv4.Client, variables map[string]interface{}) OrganizationQuery {
	var query OrganizationQuery

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		// Handle error
		fmt.Println("Unable to run the query", err)
	}
	return query
}

func printRepos(oq OrganizationQuery) {

	fmt.Println("Name: ", oq.Organization.URL)
}
