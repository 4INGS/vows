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
export GITHUB_TOKEN={Github Token here}
export GITHUB_ORG={Organization name here}
./vows
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