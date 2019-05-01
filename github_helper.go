package main

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

func buildClient() *githubv4.Client {
	token, err := fetchGithubToken()
	if err != nil {
		panic(fmt.Sprintf("Unable to build Github Client: %s", err.Error()))
	}
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	client := githubv4.NewClient(httpClient)
	return client
}
