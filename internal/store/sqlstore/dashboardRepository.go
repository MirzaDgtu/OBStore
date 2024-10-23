package sqlstore

import "obstore/internal/model"

type DashboardRepository struct {
	store *Store
}

func (r *DashboardRepository) StatsCollectors(string, string) (u []model.Online, err error) {
	return u, nil
}
