package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// Unfortunately Github GQL v4 API does not support adding teams to a repo.
// We fall back to the v3 REST API for this.

// AddTeamToRepo will assign a team to the repo
func AddTeamToRepo(team string, repo string) error {
	client := getV3Client()
	org, err := fetchOrganization()
	if err != nil {
		return fmt.Errorf("Unable to add %s to %s: %s", team, repo, err.Error())
	}
	// Will not accept Push permissions, ticket open with Github
	//opt := &github.OrganizationAddTeamRepoOptions{Permission: "Push"}
	teamID, err := getTeamID(team, client, org)
	if err != nil {
		return fmt.Errorf("Unable to add %s to %s: %s", team, repo, err.Error())
	}
	//fmt.Printf("Adding team id %d to repo %s", teamID, repo)
	resp, err := client.Organizations.AddTeamRepo(context.Background(), teamID, org, repo, nil)
	if err != nil {
		fmt.Printf("Error response is %+v", resp)
	}
	return err
}

func getTeamID(teamname string, client *github.Client, org string) (int, error) {
	opt := &github.ListOptions{}
	teams, _, err := client.Organizations.ListTeams(context.Background(), org, opt)
	if err != nil {
		return 0, fmt.Errorf("Unable to get a list of teams from Github: %s", err.Error())
	}
	for _, team := range teams {
		if team.GetName() == teamname || team.GetSlug() == teamname {
			return team.GetID(), nil
		}
	}

	return 0, fmt.Errorf("No team found with name %s", teamname)
}

func getV3Client() *github.Client {
	var tc *http.Client
	envToken, _ := fetchGithubToken()
	if len(envToken) > 0 {
		token := oauth2.Token{AccessToken: envToken}
		ts := oauth2.StaticTokenSource(&token)
		tc = oauth2.NewClient(oauth2.NoContext, ts)
	}
	client := github.NewClient(tc)
	return client
}
