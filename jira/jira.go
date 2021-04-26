// jira api種類別のfuncを配置する。関連helper funcもおいておく。

package jira

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

var httpClient = initHttpClient()

func GetIssue(issueId string) IssueResult{
	var jsonBody IssueResult

	parameters := GetIssueParams{ IssueId: issueId }
	resp := httpClient.request("GET", buildEndpoint(EndpointTemplates.GetIssue, parameters), nil)

	if err := json.NewDecoder(resp.Body).Decode(&jsonBody); err != nil {
		fmt.Printf("error %+v\n", err)
		os.Exit(1)
	}

	return jsonBody
}

func GetSprints(boardId int) []Sprint {
	var jsonBody Sprints

	parameters := GetSprintsParams{ BoardId: strconv.Itoa(boardId) }
	resp := httpClient.request("GET", buildEndpoint(EndpointTemplates.GetSprints, parameters), nil)

	if err := json.NewDecoder(resp.Body).Decode(&jsonBody); err != nil {
		fmt.Printf("error %+v\n", err)
		os.Exit(1)
	}

	return jsonBody.Values
}

func GetSprintIssues(boardId int, sprintId int) []IssueResult {
	var jsonBody SprintIssues

	parameters := GetSprintIssuesParams{
		BoardId: strconv.Itoa(boardId),
		SprintId: strconv.Itoa(sprintId),
	}
	resp := httpClient.request("GET", buildEndpoint(EndpointTemplates.GetSprintIssues, parameters), nil)

	if err := json.NewDecoder(resp.Body).Decode(&jsonBody); err != nil {
		fmt.Printf("error %+v\n", err)
		os.Exit(1)
	}

	return jsonBody.Issues
}

func GetSprintReport(boardId int, sprintId int) SprintReportResult {
	var jsonBody SprintReportResult

	parameters := GetSprintReportParams{
		BoardId: strconv.Itoa(boardId),
		SprintId: strconv.Itoa(sprintId),
	}

	resp := httpClient.request("GET", buildEndpoint(EndpointTemplates.GetSprintReport, parameters), nil)

	if err := json.NewDecoder(resp.Body).Decode(&jsonBody); err != nil {
		fmt.Printf("error %+v\n", err)
		os.Exit(1)
	}

	return jsonBody
}
