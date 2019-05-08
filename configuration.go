package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/pflag"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type teamConfig struct {
	Name       string
	Permission teamPermission
	ID         int64
}

type teamPermission string

const (
	push  teamPermission = "push"
	pull  teamPermission = "pull"
	admin teamPermission = "admin"
)

type branchProtectionRulesConfig struct {
	Pattern                      string
	DissmissesStaleReview        bool
	IsAdminEnforced              bool
	RequiresApprovingReviews     bool
	RequiredApprovingReviewCount int
	RequiresStatusChecks         bool
	RequiredStatusCheckContexts  []string
}

type configuration struct {
	Name         string
	Abbreviation string
	Default      string
	HelpText     string
}

var configs = []configuration{
	configuration{
		Name:         "AccessToken",
		Abbreviation: "t",
		HelpText:     "The security token used by the app.  It must be for an account that has branch protection write permissions.  More details at https://help.github.com/en/articles/creating-a-personal-access-token-for-the-command-line",
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
	pflag.BoolP("debug", "d", false, "Enable debug logging")
	pflag.BoolP("preview", "p", false, "No changes, will just print actions it would have taken")
	for _, config := range configs {
		pflag.StringP(strings.ToLower(config.Name), config.Abbreviation, config.Default, config.HelpText)
	}
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	if isDebug() {
		printConfiguration()
	}
}

func getConfigValue(key string) string {
	return viper.GetString(key)
}

func fetchAccessToken() string {
	return viper.GetString("AccessToken")
}

func fetchOrganization() string {
	return viper.GetString("Organization")
}

func fetchTestRepositoryID() string {
	return viper.GetString("GITHUB_TEST_REPOSITORY_ID")
}
func fetchTestRepository() string {
	return viper.GetString("GITHUB_TEST_REPOSITORY")
}
func fetchTestTeamName() string {
	return viper.GetString("GITHUB_TEST_TEAM_NAME")
}

func fetchTeams() []teamConfig {
	var cfgs []teamConfig
	value := viper.Get("Teams")

	// Viper does not have a good way to get an array of structs out
	if value == nil {
		fmt.Println("no teams found.")
		var empty []teamConfig
		return empty
	}
	switch reflect.TypeOf(value).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(value)

		for i := 0; i < s.Len(); i++ {
			i := s.Index(i).Interface()
			team := i.(map[string]interface{})

			var tc teamConfig
			tc.Name = fmt.Sprintf("%v", team["name"])
			perm := fmt.Sprintf("%v", team["permissions"])
			switch perm {
			case string(push):
				tc.Permission = push
			case string(pull):
				tc.Permission = pull
			case string(admin):
				tc.Permission = admin
			default:
				panic(fmt.Sprintf("Unknown team permission %s.  Should be 'push', 'pull', or 'admin'", perm))
			}
			cfgs = append(cfgs, tc)
		}
	default:
		panic(fmt.Sprintf("Unknown to parse teams configuration.  Should be an array of teams.  Read: %s", value))
	}
	return cfgs
}

func fetchBranchProtectionRules() branchProtectionRulesConfig {
	var config branchProtectionRulesConfig

	config.DissmissesStaleReview = viper.GetBool("BranchProtectionRules.DissmissesStaleReview")
	config.IsAdminEnforced = viper.GetBool("BranchProtectionRules.IsAdminEnforced")
	config.Pattern = viper.GetString("BranchProtectionRules.Pattern")
	config.RequiredApprovingReviewCount = viper.GetInt("BranchProtectionRules.RequiredApprovingReviewCount")
	config.RequiredStatusCheckContexts = viper.GetStringSlice("BranchProtectionRules.RequiredStatusCheckContexts")
	config.RequiresApprovingReviews = viper.GetBool("BranchProtectionRules.RequiresApprovingReviews")
	config.RequiresStatusChecks = viper.GetBool("BranchProtectionRules.RequiresStatusChecks")
	return config
}

func fetchIgnoreRepositories() []string {
	return viper.GetStringSlice("IgnoreRepositories")
}

func isDebug() bool {
	return viper.GetBool("debug")
}

func isPreview() bool {
	return viper.GetBool("preview")
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
