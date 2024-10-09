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
	return team, r.store.db.Table("teams").Select("*").Where("id=?", id).Scan(team).Error
}

func (r *TeamRepository) GetAll() (teams []model.Team, err error) {
	return teams, r.store.db.Table("teams").Select("*").Scan(&teams).Error
}
