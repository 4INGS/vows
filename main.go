package main

import (
	"fmt"
	"os"
)

func main() {
	configInit()

	repos := GetReposForOrganization()
	var list Ignorelist
	list.SetLines(fetchIgnoreRepositories())

	var gp GithubRepoHost
	err := ProcessRepositories(repos, list, gp)
	if err != nil {
		fmt.Printf("Error while processing repositories: %s\n", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
