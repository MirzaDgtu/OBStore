package store

import "obstore/internal/model"

type TeamCompositionRepository interface {
	Create(model.TeamComposition) (model.TeamComposition, error)
	DeleteById(int) error
	Update(model.TeamComposition) (model.TeamComposition, error)
	GetAll() ([]model.TeamComposition, error)
	GetByUserId(int) ([]model.TeamComposition, error)
	GetByTeamId(int) ([]model.TeamComposition, error)
	GetByID(uint) (model.TeamComposition, error)
}
