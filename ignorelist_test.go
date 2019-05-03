package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhiltelist(t *testing.T) {
	// Setup
	var w ignorelist

	// Execute
	w.SetLines([]string{"pelican"})

	// Verify
	assert.True(t, w.Onignorelist("pelican"), "Existing item not found on list")
	assert.False(t, w.Onignorelist("orange"), "Item found not on list")
}
