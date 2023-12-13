package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/xor111xor/wifi-api/internal/collect"
)

var (
	ErrNoMetric = fmt.Errorf("Not found metric")
)

type Metric struct {
	TimeScrape time.Time `json:"time"`
	Ssid       string    `json:"ssid"`
	Received   uint      `json:"received_mb"`
}

func (m *Metric) MarshalJSON() ([]byte, error) {
	type Alias Metric
	return json.MarshalIndent(&struct {
		*Alias
		TimeScrape string `json:"time"`
	}{
		Alias:      (*Alias)(m),
		TimeScrape: m.TimeScrape.Format(time.DateTime),
	}, "", "    ")
}

func (m *Metric) SetMetric() error {
	var err error

	m.Ssid, err = collect.GetSssid()
	if err != nil {
		return err
	}
	m.Received, err = collect.GetTraffic()
	if err != nil {
		return err
	}

	m.TimeScrape = time.Now()

	return nil
}

func (m *Metric) GetMetric() []string {
	setTimeFormat := m.TimeScrape.Format(time.DateTime)
	return []string{setTimeFormat, m.Ssid, strconv.FormatUint(uint64(m.Received), 10)}
}

type Repo interface {
	Add(Metric) error
	GetMetricsFromString(string) ([]Metric, error)
}
