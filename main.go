package main

import (
	"fmt"
	"os"
)

func main() {
	configInit()

	var list Ignorelist
	list.SetLines(fetchIgnoreRepositories())

	repos := GetReposForOrganization()
	if isDebug() {
		fmt.Printf("Fetched %d repos for organization\n", len(repos))
	}

	var gp GithubRepoHost
	err := ProcessRepositories(repos, list, gp)
	if err != nil {
		fmt.Printf("Error while processing repositories: %s\n", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
