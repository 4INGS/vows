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
func (p GithubProtector) AddTeamToRepo(teamID int64, repo string) error {
	client := getV3Client()
	org, err := fetchOrganization()
	if err != nil {
		return fmt.Errorf("Unable to add teamd id %d to %s: %s", teamID, repo, err.Error())
	}
	// Will not accept Push permissions, ticket open with Github
	//opt := &github.OrganizationAddTeamRepoOptions{Permission: "Push"}
	if err != nil {
		return fmt.Errorf("Unable to add team id %d to %s: %s", teamID, repo, err.Error())
	}
	//fmt.Printf("Adding team id %d to repo %s", teamID, repo)

	resp, err := client.Teams.AddTeamRepo(context.Background(), teamID, org, repo, nil)
	if err != nil {
		fmt.Printf("Error response is %+v", resp)
	}
	return err
}

// GetTeamID will convert a team name or slug into the team ID
func (p GithubProtector) GetTeamID(teamname string) (int64, error) {
	client := getV3Client()
	opt := &github.ListOptions{}
	org, err := fetchOrganization()
	if err != nil {
		return 0, fmt.Errorf("No org found, unable to get team ID for %s: %s", teamname, err.Error())
	}
	teams, _, err := client.Teams.ListTeams(context.Background(), org, opt)
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
