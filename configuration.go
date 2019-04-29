package main

import (
	"errors"
	"os"
)

func fetchGithubToken() string {
	return os.Getenv("GITHUB_TOKEN")
}

func fetchOrganization() string {
	//return os.Getenv("GITHUB_ORG")
	return "RepoFetch"
}

func fetchTestRepositoryID() (string, error) {
	repoID := os.Getenv("GITHUB_TEST_REPOSITORY_ID")
	if len(repoID) < 5 {
		return "", errors.New("Missing enviroment variable GITHUB_TEST_REPOSITORY_ID.  This needs to be set to a Github repository id that can be used for testing")
	}
	return repoID, nil
}
