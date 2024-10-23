package store

import "obstore/internal/model"

type Dashboard interface {
	StatsCollectors(string, string) ([]model.Online, error)
	StatsOnline() (model.Statistic, error)
}
