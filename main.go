package main

import (
	"os"
)

func main() {
	configInit()
	debug, _ := getConfigValue("debug")
	if debug == "true" {
		printConfiguration()
	}
	//runOrganizationQuery()
	// repos := runOrganizationQuery()
	// var gp GithubProtector
	// err := ApplyBranchProtection(repos, nil, gp)
	// if err != nil {
	// 	fmt.Printf("Unable to apply all branch protections" + err.Error())
	// }
	os.Exit(0)
}
