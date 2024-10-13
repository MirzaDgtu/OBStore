package sqlstore

import "obstore/internal/model"

type TeamCompositionRepository struct {
	store *Store
}

func (r *TeamCompositionRepository) Create(teamComposition model.TeamComposition) (model.TeamComposition, error) {
	return teamComposition, r.store.db.Create(&teamComposition).Error
}

func (r *TeamCompositionRepository) DeleteById(id int) error {
	var tc model.TeamComposition
	result := r.store.db.Table("teamcompositions").Where("id=,", id)
	err := result.First(&tc)
	if err != nil {
		return nil
	}

	return r.store.db.Delete(&tc).Error
}

func (r *TeamCompositionRepository) Update(teamComposition model.TeamComposition) (model.TeamComposition, error) {
	return teamComposition, r.store.db.Save(&teamComposition).Error
}

func (r *TeamCompositionRepository) GetAll() (tcs []model.TeamComposition, err error) {
	return tcs, r.store.db.Find(&tcs).Error
}

func (r *TeamCompositionRepository) GetByUserId(userId int) (tcs []model.TeamComposition, err error) {
	return tcs, r.store.db.Where(&model.TeamComposition{UserId: userId}).Find(&tcs).Error
}

func (r *TeamCompositionRepository) GetByTeamId(idTeam int) (tcs []model.TeamComposition, err error) {
	return tcs, r.store.db.Where(&model.TeamComposition{TeamId: idTeam}).Find(&tcs).Error
}
