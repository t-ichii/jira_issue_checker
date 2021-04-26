package main

import (
	"os"

	"github.com/jessevdk/go-flags"

	"example.com/t-ichii/jira_issue_checker/jira"
)

type CliOption struct {
	Mode string `short:"m" long:"mode" description:"execute mode (e.g.: rank_checker )"`
	BoardId int `short:"b" long:"board_id" description:"target board id (use in rank_checker mode)"`
	SprintId int `short:"s" long:"sprint_id" description:"target sprint id(use in rank_checker mode)"`
}

func main() {
	var cliOptions CliOption

	if _, err := flags.Parse(&cliOptions); err != nil {
		os.Exit(1)
	}

	// NOTE: mode ごとの振り分けはここで実施。modeのfuncは `jira/mode.go` で書くこと。（ここでロジック書かない）
	if cliOptions.Mode == "rank_checker" {
		jira.RankChecker(cliOptions.BoardId, cliOptions.SprintId)
	} else if cliOptions.Mode == "get_sprint_list" {
		jira.GetSprintList(cliOptions.BoardId)
	}

	os.Exit(0)
}
