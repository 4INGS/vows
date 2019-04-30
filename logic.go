package main

import "errors"

type protector interface {
	AddBranchProtection(repoID string) (BranchProtectionRule, error)
	UpdateBranchProtection(repoID string, rule BranchProtectionRule) error
}

// ApplyBranchProtection does things
func ApplyBranchProtection(repos []Repository, whitelist []string, protector protector) error {
	if protector == nil {
		return errors.New("No protector passed in")
	}
	// Loop over repos
	for _, v := range repos {
		// Skip if in white list
		if Contains(whitelist, v.ID) {
			continue
		}

		var ruleSet = false
		// Check if branch protection already in place and correct
		for _, r := range v.BranchProtectionRules.Nodes {
			if r.Pattern == "master" {
				if !ValidBranchProtectionRule(r) {
					protector.UpdateBranchProtection(v.ID, r)
				}
				ruleSet = true
			}
		}
		if !ruleSet {
			protector.AddBranchProtection(v.ID)
		}
	}
	return nil
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
