// jira apiのendpoints

package jira

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

var EndpointTemplates = struct {
	GetIssue string
	GetSprints string
	GetSprintIssues string
}{
	GetIssue: "/rest/api/2/issue/{{.IssueId}}",
	GetSprints: "/rest/agile/1.0/board/{{.BoardId}}/sprint",
	GetSprintIssues: "/rest/agile/1.0/board/{{.BoardId}}/sprint/{{.SprintId}}/issue",
}

func buildEndpoint(urlTemplate string, parameters interface{}) string {
	t, _ := template.New("url").Parse(urlTemplate)
	var url strings.Builder

	if err := t.Execute(&url, parameters); err != nil {
		fmt.Printf("error %+v\n", err)
		os.Exit(1)
	}

	return env.JiraHost + url.String()
}
