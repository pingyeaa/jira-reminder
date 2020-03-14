package model

type IssueStatus struct {
	ID    int32  `gorm:"column:ID;primary_key"`
	PName string `gorm:"column:pname"`
}

func (IssueStatus) TableName() string {
	return "issuestatus"
}
