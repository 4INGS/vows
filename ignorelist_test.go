package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhiltelist(t *testing.T) {
	// Setup
	var w Ignorelist

	// Execute
	w.SetLines([]string{"pelican"})

	// Verify
	assert.True(t, w.OnIgnorelist("pelican"), "Existing item not found on list")
	assert.False(t, w.OnIgnorelist("orange"), "Item found not on list")
}
