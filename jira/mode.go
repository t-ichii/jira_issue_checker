// main.go から modeで呼び出す各funcを配置する。

package jira

import "github.com/thoas/go-funk"

func RankChecker(boardId int, sprintId int) {
	issues := GetSprintIssues(boardId, sprintId)
	RenderLinkIssueTable(issues)
}

func GetSprintList(boardId int) {
	sprints := GetSprints(boardId)

	renderData := funk.Map(sprints, func(sprint Sprint) SprintReportRenderData {
		report := GetSprintReport(boardId, sprint.Id)

		users := funk.Map(report.Contents.CompletedIssues, func(issue SprintReportIssue) string {
			return issue.Assignee
		}).([]string)

		epics := funk.Map(report.Contents.CompletedIssues, func(issue SprintReportIssue) string {
			return issue.EpicField.EpicName + "(" + issue.EpicField.EpicKey + ")"
		}).([]string)

		return SprintReportRenderData{
			Id: sprint.Id,
			Name: sprint.Name,
			State: sprint.State,
			StoryPoints: struct{
				Completed float64
			}{
				Completed: report.Contents.CompletedIssuesEstimateSum.Value,
			},
			Date: struct {
				Start string
				End   string
			}{
				Start: report.Sprint.StartDate,
				End: report.Sprint.EndDate,
			},
			UserNames: funk.UniqString(users),
			Epics: funk.UniqString(epics),
		}
	}).([]SprintReportRenderData)

	RenderSprintReports(renderData)
}
