package main

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// GithubToken is the security token used by the app.  It must be for an account that has branch protection write permissions.  More details at https://help.github.com/en/articles/creating-a-personal-access-token-for-the-command-line
const GithubToken = "GITHUB_TOKEN"

// GithubOrganization is the organization name in Github that rules are enforced against.
const GithubOrganization = "GITHUB_ORG"

// GithubTestRepositoryID is only used for integraion tests.  Branch rules will be created and removed on this repo.
const GithubTestRepositoryID = "GITHUB_TEST_REPOSITORY_ID"

type configuration struct {
	Name         string
	Abbreviation string
	Default      string
	HelpText     string
}

var configs = []configuration{
	configuration{
		Name:         GithubOrganization,
		Abbreviation: "o",
		HelpText:     "The organization name in Github that rules are enforced against",
	},
	configuration{
		Name:         GithubToken,
		Abbreviation: "t",
		HelpText:     "The security token used by the app.  It must be for an account that has branch protection write permissions.  More details at https://help.github.com/en/articles/creating-a-personal-access-token-for-the-command-line",
	},
	configuration{
		Name:         GithubTestRepositoryID,
		Abbreviation: "r",
		HelpText:     "Only used for integraion tests.  Branch rules will be created and removed on this repo",
	},
	configuration{
		Name:         "debug",
		Abbreviation: "d",
		HelpText:     "Enable debug logging",
	},
}

func configInit() {
	// Setup environment
	viper.SetEnvPrefix("vows")
	for _, config := range configs {
		viper.BindEnv(strings.ToUpper(config.Name))
	}

	// Setup config file
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err == nil {
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("Config file changed:", e.Name)
		})
	}

	// Setup command line flags
	for _, config := range configs {
		pflag.StringP(strings.ToLower(config.Name), config.Abbreviation, config.Default, config.HelpText)
		pflag.Parse()
		viper.BindPFlags(pflag.CommandLine)
	}
}

func getConfigValue(key string) (string, error) {
	var err error
	value := viper.GetString(key)
	if len(value) == 0 {
		err = fmt.Errorf("Unable to find config setting for %s.  This needs to be set on the command line, in the config file, or in an environment variable", key)
	}
	return value, err
}

func fetchGithubToken() (string, error) {
	return getConfigValue(GithubToken)
}

func fetchOrganization() (string, error) {
	return getConfigValue(GithubOrganization)
}

func fetchTestRepositoryID() (string, error) {
	return getConfigValue(GithubTestRepositoryID)
}

func setConfigValue(key string, value string) {
	viper.Set(key, value)
}
