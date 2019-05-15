# Development

We welcome any additions, modifications or fixes to the code.  Just make a pull request with your changes.

## Development Environment

1. Pull down the code
1. Install dependencies
1. Build

### Pull down the code
```
go get github.com/4INGS/vows
```

### Dependencies
This project uses the newer vgo dependency management.  (Good talk on it [here](https://www.youtube.com/watch?v=F8nrpe0XWRg) if you are interested)

At the current time (Go 1.12.5), you need to enable a flag in order to take advantage of vgo functionality.  An alternative to this is to [download](https://github.com/golang/vgo) and use vgo directly.
```
export GO111MODULE=on
```
### Build
Must be using go 1.12 or above.
```
go build
```

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
Note: These require extra configuration values as shown below.  This will add and remove branch protection rules on this repo in Github, as well as adjust teams.  To support these tests, you will need to add in the following configuration values into your config.json file.
```
  "GITHUB_TEST_REPOSITORY_ID":"xxxxxxxxx-only needed if running external tests",
  "GITHUB_TEST_REPOSITORY":"xxxxxxxxx-only needed if running external tests",
  "GITHUB_TEST_TEAM_NAME":"xxxxxxxxx-only needed if running external tests",
```
Then just run
```
go test --external
```
Note: Using only a single dash in front "-external" will run the tests but they will fail.  Just use two dashes.  :)  
