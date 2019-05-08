package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/spf13/viper"
)

var (
	externalTests    = flag.Bool("external", false, "Run tests against external dependencies")
	integrationTests = flag.Bool("integration", false, "Run tests at system boundry (building and running program on command line)")
)

func init() {
	flag.Parse()
	configInit()

	if *integrationTests {
		viper.Reset()
		dir, _ := os.Getwd()
		viper.SetConfigFile(path.Join(dir, "testdata/config.test.json"))
		viper.ReadInConfig()
	}

	fmt.Printf("Config loaded: %s\n", viper.ConfigFileUsed())
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
