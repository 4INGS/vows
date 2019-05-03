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

	w := Buildignorelist()

	runOrganizationQuery()
	repos := runOrganizationQuery()
	var gp GithubProtector
	err := ApplyBranchProtection(repos, w, gp)
	if err != nil {
		fmt.Printf("Unable to apply all branch protections" + err.Error())
	}
	os.Exit(0)
}
