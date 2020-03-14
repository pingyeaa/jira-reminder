package repository

import (
	"jira-reminder/model"
)

type IssueStatusRepository struct {
}

func (IssueStatusRepository) GetByPNames(pName []string) []model.IssueStatus {
	var types []model.IssueStatus
	db.Where("pname in (?)", pName).
		Find(&types)
	return types
}
