package repository

import (
	"jira-reminder/model"
)

type JiraIssueRepository struct {
}

func (JiraIssueRepository) WithProjectInAndIssueTypeInAndIssueStatusIn(project []int32, issueType []int32, issuestatus []int32) []model.JiraIssue {
	var jiraIssues []model.JiraIssue
	db.Where("PROJECT in (?)", project).
		Where("issuetype in (?)", issueType).
		Where("issuestatus in (?)", issuestatus).
		Find(&jiraIssues)
	return jiraIssues
}
