package main

import (
	"errors"
	"fmt"
)

type protector interface {
	AddBranchProtection(repoID string) (BranchProtectionRule, error)
	UpdateBranchProtection(repoID string, rule BranchProtectionRule) error
	AddTeamToRepo(teamID int64, repoName string) error
	GetTeamID(teamname string) (int64, error)
}

// ProcessRepositories applies branch protections and proper teams to all repos
func ProcessRepositories(repos []Repository, w Ignorelist, p protector, teamname string) error {
	if p == nil {
		return errors.New("No protector passed in")
	}
	teamID, err := p.GetTeamID(teamname)
	if err != nil {
		return fmt.Errorf("Unable to find a team with the given name, %s: %s", teamname, err.Error())
	}

	// Loop over repos
	for _, r := range repos {
		// Skip if in white list
		if w.OnIgnorelist(r.Name) {
			continue
		}
		checkRepoForBranchProtections(r, p)
		p.AddTeamToRepo(teamID, r.Name)
		fmt.Printf("Processed repository %s\n", r.Name)
	}
	return nil
}

func checkRepoForBranchProtections(v Repository, p protector) {
	var ruleSet = false
	// Check if branch protection already in place and correct
	for _, r := range v.BranchProtectionRules.Nodes {
		if r.Pattern == "master" {
			if !ValidBranchProtectionRule(r) {
				if isPreview() {
					fmt.Printf("Repo %s: Incorrect branch protection found and would be updated.\n", v.Name)
				} else {
					p.UpdateBranchProtection(v.ID, r)
				}
			}
			ruleSet = true
		}
	}
	if !ruleSet {
		if isPreview() {
			fmt.Printf("Repo %s: No branch protection found and would be added.\n", v.Name)
		} else {
			p.AddBranchProtection(v.ID)
		}
	}
}

// ValidBranchProtectionRule checks to see if a branch protection matches the standards
func ValidBranchProtectionRule(rule BranchProtectionRule) bool {
	// TODO, allow this to be set in a configuration file or something
	return rule.RequiresStatusChecks == true &&
		rule.RequiresApprovingReviews == true &&
		rule.RequiredApprovingReviewCount > 0 &&
		rule.DismissesStaleReviews == true &&
		rule.IsAdminEnforced == true &&
		rule.RequiresStrictStatusChecks == true

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
