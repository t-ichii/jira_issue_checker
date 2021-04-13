package environment

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"os"
)

type Environment struct {
	JiraUsername string `required:"true" split_words:"true"` // JIRA_USER_NAME
	JiraApiToken string `required:"true" split_words:"true"` // JIRA_API_TOKEN
	JiraHost string `required:"true" split_words:"true"` // JIRA_HOST
}

func Init() Environment{
	var env = Environment{}
	if err := envconfig.Process("", &env); err != nil {
		fmt.Printf("error %+v\n", err)
		os.Exit(1)
	}
	return env
}
