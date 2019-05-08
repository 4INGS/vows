package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

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

	var err error
	if isPreview() {
		fmt.Printf("Team %s would be added to repo %s\n", t.Name, repo)
	} else {
		operation := func() error {
			ops := &github.TeamAddTeamRepoOptions{Permission: string(t.Permission)}
			_, err := client.Teams.AddTeamRepo(context.Background(), t.ID, org, repo, ops)
			return err
		}
		err = makeGitHubCall(operation)
	}
	return err
}

// GetTeamID will convert a team name or slug into the team ID
func (p GithubRepoHost) GetTeamID(teamname string) (int64, error) {
	if len(p.remoteteams) == 0 {
		err := populateRemoteTeams(&p)
		if err != nil {
			return 0, err
		}
	}
	for _, team := range p.remoteteams {
		if team.GetName() == teamname || team.GetSlug() == teamname {
			return team.GetID(), nil
		}
	}

	return 0, fmt.Errorf("No team found with name %s", teamname)
}

func populateRemoteTeams(p *GithubRepoHost) error {
	if isDebug() {
		fmt.Println("Populating Remote Teams")
	}
	client := getV3Client()
	opt := &github.ListOptions{}
	opt.PerPage = 100
	org := fetchOrganization()
	if len(org) == 0 {
		return fmt.Errorf("No org found")
	}

	var teams []*github.Team
	var resp *github.Response
	var err error
	if isDebug() {
		fmt.Println("Building team list operation")
	}
	operation := func() error {
		teams, resp, err = client.Teams.ListTeams(context.Background(), org, opt)
		return err
	}
	err = makeGitHubCall(operation)
	if resp.NextPage != 0 {
		fmt.Printf("Non-zero next page found: %d", resp.NextPage)
	}
	if err != nil {
		return fmt.Errorf("Unable to get a list of teams from Github: %s", err.Error())
	}
	if isDebug() {
		fmt.Printf("Found teams %+v\n", teams)
	}
	p.remoteteams = teams
	return nil
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

var oneSecond, _ = time.ParseDuration("1s")
var callsPerHourLimit = float64(4500)
var startTime = time.Now()
var totalCalls = 0

func makeGitHubCall(operation func() error) error {
	totalCalls++
	for {
		avgTime := time.Since(startTime).Seconds() / float64(totalCalls)
		if 3600/avgTime > callsPerHourLimit {
			if isDebug() {
				log.Printf("Lowering speed of calls (rate of %d calls per hour)", int(3600/avgTime))
			}
			time.Sleep(oneSecond)
		} else {
			break
		}
	}

	// Retry a few times if error occurs
	var err error
	retries := 0
	for {
		err = operation()
		if err == nil {
			break
		}
		if isDebug() {
			log.Printf("Retrying github call: %s", err)
		}
		retries++
		if retries > 5 {
			log.Print("Hit retry limit")
			break
		}
	}

	return err
}
