// jira apiを叩くためのhttp client

package jira

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"example.com/t-ichii/jira_issue_checker/environment"
)

type Client struct {
	http.Client
}

var env = environment.Init()

func initHttpClient() Client {
	var c = Client{}
	c.Timeout = time.Second * 30
	return c
}

func (c Client) request(method string, url string, body io.Reader) http.Response {
	req, _ := http.NewRequest(method, url, body)
	req.SetBasicAuth(env.JiraUsername, env.JiraApiToken)
	resp, err := c.Do(req)

	if err != nil {
		fmt.Printf("error %+v\n", err)
		os.Exit(1)
	}

	return *resp
}
