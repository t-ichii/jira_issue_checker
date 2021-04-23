// main.go から modeで呼び出す各funcを配置する。

package jira

func RankChecker(boardId int, sprintId int) {
	issues := GetSprintIssues(boardId, sprintId)
	RenderLinkIssueTable(issues)
}
