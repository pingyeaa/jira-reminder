package model

type IssueType struct {
	ID    int32  `gorm:"column:ID;primary_key"`
	PName string `gorm:"column:pname"`
}

func (IssueType) TableName() string {
	return "issuetype"
}
