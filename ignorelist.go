package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type ignorelist struct {
	list map[string]bool
}

// Buildignorelist will create the ignorelist based on current configuration
func Buildignorelist() ignorelist {
	var w ignorelist
	w.LoadFromFile("ignorelist.txt")
	return w
}

func (w *ignorelist) LoadFromFile(filename string) error {
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

func (w *ignorelist) SetLines(lines []string) {
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

func (w *ignorelist) Onignorelist(reponame string) bool {
	_, present := w.list[reponame]
	return present
}
