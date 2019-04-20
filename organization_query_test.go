package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildVariables(t *testing.T) {
	vars := buildVariables()
	assert.NotNil(t, vars, "Variables not created")
	assert.Contains(t, vars, "login")
	assert.Contains(t, vars, "repoCursor")
}
