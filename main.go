package main

import "fmt"

func main() {
	client := buildClient()
	vars := buildVariables()
	var query OrganizationQuery
	runOrganizationQuery(client, &query, vars)
	//printRepos(oq)
}

func printRepos(oq OrganizationQuery) {

	fmt.Println("Name: ", oq.Organization.URL)
}
