package main

import (
	"fmt"
	"os"
)

func main() {
	configInit()
	debug, _ := getConfigValue("debug")
	if debug == "true" {
		printConfiguration()
	}

	w := BuildIgnorelist()

	teamname, err := fetchDefaultTeam()
	if err != nil {
		fmt.Printf("Unable to find the default team to assign to all repos: %s", err.Error())
		os.Exit(1)
	}

	repos := GetReposForOrganization()
	var gp GithubProtector
	err = ProcessRepositories(repos, w, gp, teamname)
	if err != nil {
		fmt.Printf("Unable to apply all branch protections" + err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
