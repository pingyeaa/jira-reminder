package repository

import (
	"jira-reminder/model"
)

type IssueTypeRepository struct {
}

func (IssueTypeRepository) GetByPNames(pName []string) []model.IssueType {
	var issueTypes []model.IssueType
	db.Where("pname in (?)", pName).
		Find(&issueTypes)
	return issueTypes
}
