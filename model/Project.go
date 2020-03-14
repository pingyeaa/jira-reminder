package model

type Project struct {
	ID   int32  `gorm:"column:ID;primary_key"`
	PKey string `gorm:"column:pkey"`
}

func (Project) TableName() string {
	return "project"
}
