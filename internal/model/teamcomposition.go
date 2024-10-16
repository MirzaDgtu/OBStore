package model

import (
	"gorm.io/gorm"
)

type TeamComposition struct {
	gorm.Model
	TeamId int `gorm:"column:team_id" json:"team_id"`
	UserId int `gorm:"column:user_id" json:"user_id"`
}

func (TeamComposition) TableName() string {
	return "teamcompositions"
}
