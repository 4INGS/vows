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
func (p GithubRepoHost) AddTeamToRepo(t teamConfig, repo string) error {
	client := getV3Client()
	org := fetchOrganization()
	if len(org) == 0 {
		return fmt.Errorf("Unable to add team %s to %s: No organziation in config", t.Name, repo)
	}
	if isDebug() {
		fmt.Printf("Adding team id %s to repo %s", t.Name, repo)
	}

	ops := &github.TeamAddTeamRepoOptions{Permission: string(t.Permission)}
	resp, err := client.Teams.AddTeamRepo(context.Background(), t.ID, org, repo, ops)
	if err != nil {
		fmt.Printf("Error response is %+v", resp)
	}
	return err
}

// GetTeamID will convert a team name or slug into the team ID
func (p GithubRepoHost) GetTeamID(teamname string) (int64, error) {
	client := getV3Client()
	opt := &github.ListOptions{}
	org := fetchOrganization()
	if len(org) == 0 {
		return 0, fmt.Errorf("No org found, unable to get team ID for %s", teamname)
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
	at := fetchAccessToken()
	if len(at) == 0 {
		panic("No access token found, unable to build v3 client")
	}
	token := oauth2.Token{AccessToken: at}
	ts := oauth2.StaticTokenSource(&token)
	tc = oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)
	return client
}
