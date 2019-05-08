package main

import (
	"fmt"
)

// Ignorelist is a list of repositories that should be ignore while processing
type Ignorelist struct {
	list map[string]bool
}

// SetLines will set the list of repos that should be ignored
func (w *Ignorelist) SetLines(lines []string) {
	set := make(map[string]bool)
	for _, v := range lines {
		if len(v) > 0 {
			if isDebug() {
				fmt.Printf("Will ignore repo %s\n", v)
			}
			set[v] = true
		}
	}
	w.list = set
}

// OnIgnorelist will check to see if the repository name should be ignored
func (w *Ignorelist) OnIgnorelist(reponame string) bool {
	_, present := w.list[reponame]
	if isDebug() {
		fmt.Printf("Checking ignore list for repo %s.  Found: %t", reponame, present)
	}
	return present
}
