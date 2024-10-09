package model

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	NameTeam string `gorm:"column:nameTeam" json:"name_team" validate:"required"`
}

func (Team) TableName() string {
	return "teams"
}
