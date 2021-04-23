package jira

func RankChecker(boardId int, sprintId int) {
	issues := GetSprintIssues(boardId, sprintId)
	RenderLinkIssueTable(issues)
}
