package main

import "fmt"

func main() {
	var gp GithubProtector
	repos := runOrganizationQuery()
	err := ApplyBranchProtection(repos, nil, gp)
	if err != nil {
		fmt.Printf("Unable to apply all branch protections" + err.Error())
	}
}
