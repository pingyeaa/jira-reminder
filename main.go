package main

import (
	"jira-reminder/service"
)

func main() {
	jiraService := new(service.JiraService)
	jiraService.Remind()
}
