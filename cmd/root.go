package cmd

import (
	"github.com/go-co-op/gocron"
	"github.com/xor111xor/wifi-api/internal/api"
	"github.com/xor111xor/wifi-api/internal/models"
	"time"
)

func Run() error {
	repo, err := GetRepo()
	if err != nil {
		return err
	}
	metric := new(models.Metric)

	// Get metrics every minute
	s := gocron.NewScheduler(time.UTC)
	_, err = s.Cron("*/1 * * * *").Do(func() error {
		err := metric.SetMetric()
		if err != nil {
			return err
		}
		err = repo.Add(*metric)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	s.StartAsync()
	err = api.RunApi(repo)
	if err != nil {
		return err
	}

	return nil
}
