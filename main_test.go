package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	externalTests    = flag.Bool("external", false, "Run tests against external dependencies")
	integrationTests = flag.Bool("integration", false, "Run tests at system boundry (building and running program on command line)")
)

func init() {
	flag.Parse()
	configInit()
}

func TestMain(m *testing.M) {

	if *integrationTests {
		buildProgram()
	}

	code := m.Run()
	os.Exit(code)
}

func buildProgram() {
	build := exec.Command("go", "build")
	err := build.Run()
	if err != nil {
		panic(fmt.Sprintf("Unable to build project: %s", err.Error()))
	}
}

// TestLookup will test the query to github
func TestLookup(t *testing.T) {
	assert.Equal(t, "1", "1", "Expected then actual")
	//t.Fatal("This test fails")
}
