package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// GithubToken is the security token used by the app.  It must be for an account that has branch protection write permissions.  More details at https://help.github.com/en/articles/creating-a-personal-access-token-for-the-command-line
const GithubToken = "GITHUB_TOKEN"

// GithubOrganization is the organization name in Github that rules are enforced against.
const GithubOrganization = "GITHUB_ORG"

// GithubTestRepositoryID is only used for integraion tests.  Branch rules will be created and removed on this repo.
const GithubTestRepositoryID = "GITHUB_TEST_REPOSITORY_ID"

var keys = []string{
	GithubToken,
	GithubOrganization,
	GithubTestRepositoryID,
}

func configInit() {
	// Setup environment
	viper.SetEnvPrefix("vows")
	for _, k := range keys {
		viper.BindEnv(k)
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
