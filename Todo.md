# Todo
* Drastically expand configuration.  Make single file with all configs.  Move into it: rules, ignore list, teams and permissions
* Redo command line parameters so testing flags work (-test.timeout) and single dashes work (-integration, -external).  The will involve redoing how int and ext tests are run
* Break external testing configuration into a separate file.
* Add in logging with proper log levels
* Support adding single admin user to all repos

## Completed
* Finsih Main.test
* Move configuration into config file as well as environment variables
* Support integration and external test types
* Support command line parameters for configuration
* Support preview mode (just list tasks, don't execute)
* Implement ignore list
* Check all errors are propertly logged or used
* Support adding teams to all repos
