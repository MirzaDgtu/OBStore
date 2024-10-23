package sqlstore

import "obstore/internal/model"

type DashboardRepository struct {
	store *Store
}

func (r *DashboardRepository) StatsCollectors(string, string) ([]model.Online, error) {
	return []&model.Online, nil
}
