package store

import (
	"obstore/internal/model"
)

type TeamRepository interface {
	Create(model.Team) (model.Team, error)
	Update(model.Team) (model.Team, error)
	DeleteById(int) error
	ById(int) (model.Team, error)
	All() ([]model.Team, error)
	TeamComposition(uint) (model.Team, error)
}
