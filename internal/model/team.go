package model

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Id       int    `gorm:"column:id" json:"id"`
	NameTeam string `gorm:"column:nameteam" json:"name_team"`
}

func (Team) TableName() string {
	return "teams"
}
