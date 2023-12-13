package main

import (
	"github.com/go-co-op/gocron"
	"github.com/xor111xor/wifi-api/internal/api"
	"github.com/xor111xor/wifi-api/internal/models"
	"github.com/xor111xor/wifi-api/internal/repository"
	"time"
)

func main() {
	r := repository.NewInmemoryRepo()
	t := new(models.Metric)

	// Get metrics every minute
	s := gocron.NewScheduler(time.UTC)
	_, err := s.Cron("*/1 * * * *").Do(func() {
		err := t.SetMetric()
		if err != nil {
			panic(err)
		}
		err = r.Add(*t)
		if err != nil {
			panic(err)
		}
	})
	if err != nil {
		panic(err)
	}
	s.StartAsync()

	api.RunApi(r)
}
