package repository

import (
	"jira-reminder/model"
)

type ProjectRepository struct {
}

func (ProjectRepository) GetByPKeys(pKey []string) []model.Project {
	var types []model.Project
	db.Where("pkey in (?)", pKey).
		Find(&types)
	return types
}
