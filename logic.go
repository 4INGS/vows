package main

import (
	"errors"
	"fmt"
)

// RepoHost is a host of repositories (example: Gitlab, Github, etc.)
type RepoHost interface {
	AddBranchProtection(repoID string) (BranchProtectionRule, error)
	UpdateBranchProtection(repoID string, rule BranchProtectionRule) error
	AddTeamToRepo(team *teamConfig, repoName string) error
	GetTeamID(teamname string) (int64, error)
	TeamAccessToRepo(team string, repo string) (string, error)
}

// ProcessRepositories applies branch protections and proper teams to all repos
func ProcessRepositories(repos []Repository, list Ignorelist, p RepoHost) error {
	if p == nil {
		return errors.New("No RepoHost passed in")
	}

	// Populate ids for teams
	teams := fetchTeams()
	for _, t := range teams {
		teamID, err := p.GetTeamID(t.Name)
		if err != nil {
			return fmt.Errorf("Unable to find a team with the given name, %s: %s", t.Name, err.Error())
		}
		t.ID = teamID
	}

	// Loop over repos
	for _, r := range repos {
		// Skip if in white list
		if list.OnIgnorelist(r.Name) {
			continue
		}
		checkRepoForBranchProtections(r, p)

		for _, t := range teams {
			checkRepoForTeamAccess(t, r, p)
		}
		fmt.Printf("Processed repository %s\n", r.Name)
	}
	return nil
}

func checkRepoForTeamAccess(t *teamConfig, r Repository, p RepoHost) error {
	access, err := p.TeamAccessToRepo(t.Name, r.Name)
	if err != nil {
		return err
	}
	if !accessMatches(access, t.Permission) {
		if isPreview() {
			fmt.Printf("Would give team %s access to repo %s\n", t.Name, r.Name)
		} else {
			p.AddTeamToRepo(t, r.Name)
		}
	}
	return nil
}

func accessMatches(access string, permission teamPermission) bool {
	return (access == "READ" && permission == pull) ||
		(access == "WRITE" && permission == push) ||
		(access == "ADMIN" && permission == admin)
}

func checkRepoForBranchProtections(v Repository, p RepoHost) {
	var ruleSet = false
	configRules := fetchBranchProtectionRules()

	// Check if branch protection already in place and correct
	for _, r := range v.BranchProtectionRules.Nodes {
		if r.Pattern == configRules.Pattern {
			if !validBranchProtectionRule(r) {
				updateBranchProtection(v, r, p)
			}
			ruleSet = true
		}
	}
	if !ruleSet {
		addBranchProtection(v, p)
	}
}

func updateBranchProtection(v Repository, r BranchProtectionRule, p RepoHost) {
	if isPreview() {
		fmt.Printf("Repo %s: Incorrect branch protection found and would be updated.\n", v.Name)
	} else {
		if isDebug() {
			fmt.Printf("Updating branch protection for %s\n", v.Name)
		}
		p.UpdateBranchProtection(v.ID, r)
	}
}
func addBranchProtection(v Repository, p RepoHost) {
	if isPreview() {
		fmt.Printf("Repo %s: No branch protection found and would be added.\n", v.Name)
	} else {
		if isDebug() {
			fmt.Printf("Adding branch protection to %s\n", v.Name)
		}
		p.AddBranchProtection(v.ID)
	}
}

func validBranchProtectionRule(rule BranchProtectionRule) bool {
	configRules := fetchBranchProtectionRules()

	result := rule.RequiresStatusChecks == configRules.RequiresStatusChecks &&
		rule.RequiresApprovingReviews == configRules.RequiresApprovingReviews &&
		rule.RequiredApprovingReviewCount == configRules.RequiredApprovingReviewCount &&
		rule.DismissesStaleReviews == configRules.DismissesStaleReviews &&
		rule.IsAdminEnforced == configRules.IsAdminEnforced &&
		rule.RequiresStrictStatusChecks == configRules.RequiresStrictStatusChecks

	if !result && isDebug() {
		fmt.Printf("Unmatched rule '%s'.  Differs from branch protection listed in configuration\nRule: %+v\nConfig: %+v\n", rule.Pattern, rule, configRules)
	}

	return result
}

// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
