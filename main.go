package main

import (
	"fmt"
	"os"
)

func main() {
	configInit()
	value, _ := getConfigValue("GITHUB_ORG")
	fmt.Println(value)
	//runOrganizationQuery()
	// repos := runOrganizationQuery()
	// var gp GithubProtector
	// err := ApplyBranchProtection(repos, nil, gp)
	// if err != nil {
	// 	fmt.Printf("Unable to apply all branch protections" + err.Error())
	// }
	os.Exit(0)
}
