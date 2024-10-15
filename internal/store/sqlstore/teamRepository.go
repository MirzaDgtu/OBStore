package sqlstore

import "obstore/internal/model"

type TeamRepository struct {
	store *Store
}

func (r *TeamRepository) Create(u model.Team) (model.Team, error) {
	err := r.store.db.Create(&u).Error
	return u, err
}

func (r *TeamRepository) Update(u model.Team) (model.Team, error) {
	return u, r.store.db.Table("teams").Save(&u).Error
}

func (r *TeamRepository) DeleteById(id int) error {
	var team model.Team
	result := r.store.db.Table("teams").Where("id=?", id)
	err := result.First(&team).Error
	if err != nil {
		return err
	}
	return r.store.db.Delete(&team).Error
}

func (r *TeamRepository) GetById(id int) (team model.Team, err error) {
	team.ID = uint(id)
	return team, r.store.db.First(&team).Error
}

func (r *TeamRepository) GetAll() (teams []model.Team, err error) {
	return teams, r.store.db.Table("teams").Select("*").Scan(&teams).Error
}

func (r *TeamRepository) TeamComposition(id uint) (tc model.Team, err error) {
	return tc, r.store.db.Table("teams").Select("teams.id as teamId, teams.nameTeam, u.id as userId, CONCAT(u.firstname, ' ', u.lastname) as userName, u.inn").
		Joins("left join teamcompositions as tc on tc.teamId = teams.id").
		Joins("left join users as u on u.id = tc.userId").
		Where("teams.id=?", id).Scan(&tc).Error

}
