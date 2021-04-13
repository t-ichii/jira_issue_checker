build:
	go build -o ./bin/jira_issue_checker

build_debug:
	go build -gcflags="-N -l" -o ./bin/jira_issue_checker_debug
