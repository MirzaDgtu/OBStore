package model

import (
	"gorm.io/gorm"
)

type TeamComposition struct {
	gorm.Model
	TeamId int `gorm:"column:teamId" json:"team_id"`
	UserId int `gorm:"userId" json:"users"`
}

func (TeamComposition) TableName() string {
	return "teamcompositions"
}
