package main

import (
	"fmt"
	"os"
)

func main() {
	configInit()
	var gp GithubProtector
	repos := runOrganizationQuery()
	err := ApplyBranchProtection(repos, nil, gp)
	if err != nil {
		fmt.Printf("Unable to apply all branch protections" + err.Error())
	}
	os.Exit(0)
}
