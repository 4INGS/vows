package main

import (
	"os"
)

func fetchGithubToken() string {
	return os.Getenv("GITHUB_TOKEN")
}

func fetchOrganization() string {
	//return os.Getenv("GITHUB_ORG")
	return "RepoFetch"
}
