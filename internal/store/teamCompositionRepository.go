package store

import "obstore/internal/model"

type TeamCompositionRepository interface {
	Create(model.TeamComposition) (model.TeamComposition, error)
	DeleteById(int) error
	Update(model.TeamComposition) (model.TeamComposition, error)
	All() ([]model.TeamComposition, error)
	ByUserId(int) ([]model.TeamComposition, error)
	ByTeamId(int) ([]model.TeamComposition, error)
	ByID(uint) (model.TeamComposition, error)
}
