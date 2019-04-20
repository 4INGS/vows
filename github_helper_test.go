package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildClient(t *testing.T) {
	client := buildClient()
	assert.NotNil(t, client, "Client could not be created")
}
