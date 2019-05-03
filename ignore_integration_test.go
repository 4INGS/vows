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

	var w ignorelist
	err := w.LoadFromFile("testdata/test-ignorelist.txt")
	assert.Nil(t, err)
}
