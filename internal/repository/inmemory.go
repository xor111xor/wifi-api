package repository

import (
	"github.com/xor111xor/wifi-api/internal/models"
	"sync"
)

type InMemoryRepo struct {
	sync.RWMutex
	metrics []models.Metric
}

func NewInmemoryRepo() (*InMemoryRepo, error) {
	return &InMemoryRepo{
		metrics: []models.Metric{},
	}, nil
}

func (in *InMemoryRepo) Add(m models.Metric) error {
	in.Lock()
	m.ID = len(in.metrics) + 1

	in.metrics = append(in.metrics, m)
	in.Unlock()
	return nil
}

func (in *InMemoryRepo) GetMetricsFromString(day string) ([]models.Metric, error) {
	var metric []models.Metric
	in.Lock()
	for _, y := range in.metrics {
		if y.TimeScrape.Format("2006-01-02") == day {
			metric = append(metric, y)
		}
	}
	in.Unlock()
	return metric, nil
}
