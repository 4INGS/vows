# Vows
Apply a standard set of rules to all Github repositories in an organization.

Loops over all the repositories in your organization and sets branch protections on the master branch.  These branch protections are currently hard coded, but a future enhancement will allow these to be customized to your needs.

## Building
```
go build
```

## Running
Super simple, just run:
```
./vows
```

## Ignore List
You can supply a list of repos that should be ignored.  Create a file named "ignorelist.txt" with a single line per repo to ignore.
```
RepoName1
RepoName2
```

## Configuration
| Key | Example Value | Details |
| --- | ------------- | ------- |
|github_token|xxxxxxxxxxxxx|Github access token, Your token can be created by following [these instructions](https://help.github.com/en/articles/creating-a-personal-access-token-for-the-command-line)|
|github_org|MyAwesomeOrg|The Github organization to work against|
|preview|false|Do not take any action, only print a list of actions that would be taken|
|debug|true|If the program should print out debugging information|
|github_test_repository_id|xxxxxxxxxxxxxxxxxxxx|The internal github repository id to use for external tests|

These can be configured through either
* Environment Variables
* Configuration file
* Command line parameters

You can specify configuration in any or all of the methods listed.  For example, Github tokens can be sent in through environment variables, while the debug option is listed on the command line.

### Environment Variables
All keys should be prefixed with "VOWS_" when setting through an environment variable.  Environment variables should be all UPPER CASE.
```
export VOWS_GITHUB_TOKEN={Github Token here}
export VOWS_GITHUB_ORG={Organization name here}
```

### Config file
You can also configure the application using a json configuration file
```
{
    "GITHUB_ORG":"Your_Org_Here",
    "GITHUB_TOKEN":"Your_Token_Here"
}
```
### Command line configuration
```
./vows --github_org=myorg --debug=true --preview=true
```

### Token Permissions
This app needs **repo** and **organizations** permissions
![token permissions](assets/repo-permissions.png)

You can create the token at [this link](https://help.github.com/en/articles/creating-a-personal-access-token-for-the-command-line)

## Testing

### Unit tests
These are very fast tests (less then 1 second) to verify internal logic 
```
go test
```

### Integration tests
These tests take a bit longer, but verify the system bounderies are correct
```
go test --integration
```

### External tests
Slower tests that exercise external systems
Note: These require a configuration value for GITHUB_TEST_REPOSITORY_ID. This will attempt to add and remove branch protection rules on this repo in Github
```
export GITHUB_TEST_REPOSITORY_ID={RepoIDYouDoNotCareAbout}
go test --external
```
Note: Using only a single dash in front "-external" will run the tests but they will fail.  Just use two dashes.  :)  
