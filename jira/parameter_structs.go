// endpoints.goでbuildEndpointsで使うParameterのStructsを定義

package jira

type GetIssueParams struct {
	IssueId string
}

type GetSprintsParams struct {
	BoardId string
}

type GetSprintIssuesParams struct {
	BoardId string
	SprintId string
}

type GetSprintReportParams struct {
	BoardId string
	SprintId string
}
