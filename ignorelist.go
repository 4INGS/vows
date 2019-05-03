package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// Ignorelist is a list of repositories that should be ignore while processing
type Ignorelist struct {
	list map[string]bool
}

// BuildIgnorelist will create the Ignorelist based on current configuration
func BuildIgnorelist() Ignorelist {
	var w Ignorelist
	w.LoadFromFile("Ignorelist.txt")
	return w
}

// LoadFromFile will import the ignore list from a file
func (w *Ignorelist) LoadFromFile(filename string) error {
	dir, err := os.Getwd()
	fullPath := path.Join(dir, filename)
	content, err := ioutil.ReadFile(fullPath)
	if err != nil {
		err = fmt.Errorf("Error reading file %s: %s", fullPath, err.Error())
		return err
	}
	lines := strings.Split(string(content), "\n")
	w.SetLines(lines)
	return nil
}

// SetLines will set the list of repos that should be ignored
func (w *Ignorelist) SetLines(lines []string) {
	set := make(map[string]bool)
	for _, v := range lines {
		if len(v) > 0 {
			if isDebug() {
				fmt.Printf("Storing item of %s\n", v)
			}
			set[v] = true
		}
	}
	w.list = set
}

// OnIgnorelist will check to see if the repository name should be ignored
func (w *Ignorelist) OnIgnorelist(reponame string) bool {
	_, present := w.list[reponame]
	return present
}
