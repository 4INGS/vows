# Github Management Ops

Apply a standard set of rules to all Github repositories in an organization

## Building
```
go build
```

## Running
You will need to set some configuration before running the program
```
export GITHUB_TOKEN={Github Token here}
export GITHUB_ORG={Organization name here}
./github-management-ops
```

## Testing
Run unit tests
```
go test
```
Run integration tests
```
go test -tags=integration
```