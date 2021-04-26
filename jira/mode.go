// main.go から modeで呼び出す各funcを配置する。

package jira

import "github.com/thoas/go-funk"

func RankChecker(boardId int, sprintId int) {
	issues := GetSprintIssues(boardId, sprintId)
	RenderLinkIssueTable(issues)
}

func GetSprintList(boardId int) {
	sprints := GetSprints(boardId)
	renderData := funk.Map(sprints, func(sprint Sprint) SprintReportData {
		report := GetSprintReport(boardId, sprint.Id)
		return SprintReportData{
			Id: sprint.Id,
			Name: sprint.Name,
			State: sprint.State,
			StoryPoints: struct{ Completed string }{Completed: report.Contents.CompletedIssuesEstimateSum.Value},
		}
	}).([]SprintReportData)
	RenderSprintReports(renderData)
}
