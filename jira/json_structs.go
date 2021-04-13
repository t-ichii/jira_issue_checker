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

type IssueResult struct {
	Key string `json:"key"`
	Fields struct {
		Summary string `json:"summary"`
		IssueLinks []IssueLinks `json:"issuelinks"`
		AssignedSprints []IssueSprint `json:"customfield_10007"`
	} `json:"fields"`
}

type IssueLinks struct {
	Id string `json:"id"`
	OutwardIssue LinkIssue `json:"outwardIssue"`
	InwardIssue LinkIssue `json:"inwardIssue"`
}

type LinkIssue struct {
	Key string `json:"key"`
	Fields struct{
		Summary string `json:"summary"`
		Status struct{
			Name string `json:"name"`
		} `json:"status"`
	} `json:"fields"`
}

type Sprints struct {
	Values []Sprint `json:"values"`
}

type Sprint struct {
	Id int `json:"id"`
	Self string `json:"self"`
	State string `json:"state"`
	Name string `json:"name"`
	StartDate string `json:"startDate"`
	EndDate string `json:"endDate"`
	CompleteDate string `json:"completeDate"`
	OriginBoardId int `json:"originBoardId"`
	Goal string `json:"goal"`
}

type IssueSprint struct {
	Id int `json:"id"`
	Name string `json:"name"`
	State string `json:"state"`
	BoardId int `json:"boardId"`
}

type SprintIssues struct {
	Issues []IssueResult `json:"issues"`
}
