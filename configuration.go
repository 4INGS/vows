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
		Name:         "preview",
		Abbreviation: "p",
		HelpText:     "Enable debug logging",
		Default:      "false",
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
	// Ignore unknown flags
	et := pflag.ParseErrorsWhitelist{UnknownFlags: true}
	pflag.CommandLine.ParseErrorsWhitelist = et
	pflag.BoolP("debug", "d", false, "Enable Debug Logging")
	for _, config := range configs {
		pflag.StringP(strings.ToLower(config.Name), config.Abbreviation, config.Default, config.HelpText)
	}
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
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

func isDebug() bool {
	value, _ := getConfigValue("debug")
	return value == "true"
}
func isPreview() bool {
	value, err := getConfigValue("preview")
	// This one is a special case.  If we can't determine if we are in preview mode, we immediately
	// stop to ensure nothing is written if the user wanted to be in preview mode
	if err != nil {
		panic("Unable to determine if in preview mode " + err.Error())
	}
	return value == "true"
}

func setConfigValue(key string, value string) {
	viper.Set(key, value)
}

func printConfiguration() {
	fmt.Println("Printing Debug Configuration")
	for _, key := range viper.AllKeys() {
		fmt.Printf("Configuration %s=%s\n", key, viper.Get(key))
	}
}
