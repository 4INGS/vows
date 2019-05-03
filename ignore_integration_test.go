package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Move to integration tests
func TestBuildWhiltelist(t *testing.T) {
	if !*integrationTests {
		return
	}

	var w Ignorelist
	err := w.LoadFromFile("testdata/test-Ignorelist.txt")
	assert.Nil(t, err)
}
