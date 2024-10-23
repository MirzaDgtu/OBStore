package store

import "obstore/internal/model"

type DashboardRepository interface {
	StatsCollectors(string, string) ([]model.Online, error)
	StatsOnline() (model.Statistic, error)
}
