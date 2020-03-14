package model

import (
	"time"
)

type JiraIssue struct {
	ID          int32     `gorm:"column:ID;primary_key"`
	Assignee    string    `gorm:"column:ASSIGNEE"`
	IssueType   int16     `gorm:"column:issuetype"`
	Summary     string    `gorm:"column:SUMMARY"`
	Resolution  int16     `gorm:"column:RESOLUTION"`
	IssueStatus int16     `gorm:"column:issuestatus"`
	Project     int16     `gorm:"column:PROJECT"`
	DueDate     time.Time `gorm:"column:DUEDATE;type:date"`
}

func (JiraIssue) TableName() string {
	return "jiraissue"
}
