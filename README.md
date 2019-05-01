# Vows
Apply a standard set of rules to all Github repositories in an organization

## Building
```
go build
```
## Configure
You can supply a list of repos that should be ignored.
```
RepoName1
RepoName2
```

## Running
You will need to set some configuration before running the program.  Your token can be created by following [these instructions](https://help.github.com/en/articles/creating-a-personal-access-token-for-the-command-line).
```
export VOWS_GITHUB_TOKEN={Github Token here}
export VOWS_GITHUB_ORG={Organization name here}
./vows
```

## Config file
You can also configure the application using a json configuration file
```
{
    "GITHUB_ORG":"Your_Org_Here",
    "GITHUB_TOKEN":"Your_Token_Here"
}
```

## Testing
Run unit tests
```
go test
```
Run integration tests
Note: This will attempt to add and remove branch protection rules on this repo in Github
```
export GITHUB_TEST_REPOSITORY_ID={RepoIDYouDoNotCareAbout}
go test -tags=integration
```