package repository

import (
	"encoding/binary"
	"encoding/json"
	"github.com/boltdb/bolt"
	"github.com/xor111xor/wifi-api/internal/models"
	"sync"
)

const Bucket string = "metric"

type btRepo struct {
	db *bolt.DB
	sync.RWMutex
}

func NewBoltRepo(dbfile string) (*btRepo, error) {
	db, err := bolt.Open(dbfile, 0600, nil)
	if err != nil {
		return nil, err
	}
	return &btRepo{
		db: db,
	}, nil
}

func (bt *btRepo) Add(m models.Metric) error {
	bt.Lock()
	// Store the user model in the user bucket using the ID as the key.
	err := bt.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(Bucket))
		if err != nil {
			return err
		}
		id, _ := b.NextSequence()
		m.ID = int(id)

		encoded, err := json.Marshal(m)
		if err != nil {
			return err
		}
		return b.Put(itob(m.ID), encoded)
	})
	bt.Unlock()
	return err
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func (bt *btRepo) GetMetricsFromString(day string) ([]models.Metric, error) {
	var metrics []models.Metric
	var current_metric models.Metric

	err := bt.db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(Bucket))

		err := b.ForEach(func(k, v []byte) error {
			err := json.Unmarshal(v, &current_metric)
			if err != nil {
				return err
			}
			if current_metric.TimeScrape.Format("2006-01-02") == day {
				metrics = append(metrics, current_metric)
			}
			return nil
		})

		return err
	})
	return metrics, err
}
