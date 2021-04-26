// cli上でrenderするfuncを配置

package jira

import (
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/thoas/go-funk"
)

type LinkIssueRenderData struct {
	Key string
	IsBlockedBy struct{
		Issues []IssueLinks
		ValidRank bool

	}
	Blocks []IssueLinks
}

type SprintReportRenderData struct {
	Id int
	Name string
	State string
	StoryPoints struct{
		Completed float64
	}
	Date struct{
		Start string
		End string
	}
	UserNames []string
	Epics []string
}

func RenderLinkIssueTable(issues []IssueResult) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetColWidth(30)
	table.SetHeader([]string{
		"Key",
		"Is Blocked By",
		"Blocks",
	})

	currentSprintIssueMap := make( map[string]*IssueResult)

	for _, currentIssue := range issues {
		currentSprintIssueMap[currentIssue.Key] = &currentIssue
	}

	for index, currentIssue := range issues {
		if index != 0 {
			table.SetRowLine(true)
		}

		inwardIssues := funk.Filter(currentIssue.Fields.IssueLinks, func(issue IssueLinks) bool {
			// is blocked byの判定
			return issue.InwardIssue.Key != "" // TODO: ここを outwardIssueがnilで有ることで判定したい
		})

		outwardIssues := funk.Filter(currentIssue.Fields.IssueLinks, func(issue IssueLinks) bool {
			// blocksの判定
			return issue.OutwardIssue.Key != "" // TODO: ここを outwardIssueがnilで有ることで判定したい
		})

		inwardIssueStrings := funk.Reduce(inwardIssues, func(acc string, issue IssueLinks) string {
			var isExistIssueCurrentSprint bool = currentSprintIssueMap[issue.InwardIssue.Key] != nil
			var isValidRank bool = false

			if !isExistIssueCurrentSprint {
				// TODO: state futureのsprintに入っているか検査する
			} else {
				// rankがcurrentIssue前に入っているか検査する
				// inwardIssueがsprint 外の場合はこのループに引っかからないので対応しない
				for _, targetIssue := range issues {
					// note: issuesがrankの順番に並んでることを前提とする
					if targetIssue.Key == currentIssue.Key {
						// この場合、currentIssue < InwardIssueが先にくる
						break
					}
					if targetIssue.Key == issue.InwardIssue.Key {
						// この場合、 rankは inwardIssue < currentIssue
						isValidRank = true
						break
					}
				}
			}

			var validMessage string = "in sprint: " + strconv.FormatBool(isExistIssueCurrentSprint)
			if isExistIssueCurrentSprint {
				validMessage = validMessage + " / rank valid: " + strconv.FormatBool(isValidRank)
			}

			return acc + "\n\n" +
				issue.InwardIssue.Key +
				"(" +
				issue.InwardIssue.Fields.Status.Name +
				" / " + validMessage +
				")\n" +
				issue.InwardIssue.Fields.Summary
		}, "").(string)

		outwardIssueStrings := funk.Reduce(outwardIssues, func(acc string, issue IssueLinks) string {
			var isExistIssueCurrentSprint bool = currentSprintIssueMap[issue.OutwardIssue.Key] != nil
			var isValidRank bool = false

			if !isExistIssueCurrentSprint {
				// TODO: state futureのsprintに入っているか検査する
			} else {
				// rankがcurrentIssueあとに入っているか検査する
				// inwardIssueがsprint 外の場合はこのループに引っかからないので対応しない
				for _, targetIssue := range issues {
					// note: issuesがrankの順番に並んでることを前提とする
					if targetIssue.Key == currentIssue.Key {
						// この場合、currentIssue < InwardIssueが先にくる
						isValidRank = true
						break
					}
					if targetIssue.Key == issue.InwardIssue.Key {
						// この場合、 rankは inwardIssue < currentIssue
						break
					}
				}
			}

			var validMessage string = "in sprint: " + strconv.FormatBool(isExistIssueCurrentSprint)
			if isExistIssueCurrentSprint {
				validMessage = validMessage + " / rank valid: " + strconv.FormatBool(isValidRank)
			}

			return acc + "\n\n" +
				issue.OutwardIssue.Key +
				"(" +
				issue.OutwardIssue.Fields.Status.Name +
				" / " + validMessage +
				")\n" +
				issue.OutwardIssue.Fields.Summary
		}, "").(string)


		if len(outwardIssueStrings) > 2 {
			outwardIssueStrings = outwardIssueStrings[2:]
		}

		if len(inwardIssueStrings) > 2 {
			inwardIssueStrings = inwardIssueStrings[2:]
		}

		table.Append([]string{
			currentIssue.Key + "\n" + currentIssue.Fields.Summary,
			inwardIssueStrings,
			outwardIssueStrings,
		})
	}
	table.Render()
}

func RenderSprintReports(sprints []SprintReportRenderData) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetColWidth(30)
	table.SetHeader([]string{
		"Id",
		"Name",
		"State",
		"Completed",
		"Start Date",
		"End Date",
		"Assignee",
		"Epics",
	})

	for _, sprint := range sprints {
		table.Append([]string {
			strconv.Itoa(sprint.Id),
			sprint.Name,
			sprint.State,
			strconv.FormatFloat(sprint.StoryPoints.Completed, 'f', 2, 32),
			sprint.Date.Start,
			sprint.Date.End,
			strings.Join(sprint.UserNames, `, `),
			strings.Join(sprint.Epics, `, `),
		})
	}

	table.Render()
}
