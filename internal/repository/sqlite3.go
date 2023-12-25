// go:build !inmemory

package repository

import (
	"database/sql"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/xor111xor/wifi-api/internal/models"
)

const (
	createTableMetrics string = `CREATE TABLE IF NOT EXISTS "metrics" (
	  "id" INTEGER,
		"time_scrape" DATETIME NOT NULL,
		"ssid" TEXT NOT NULL,
		"received" INTEGER DEFAULT 0,
		PRIMARY KEY("id")
		);`
)

type dbRepo struct {
	db *sql.DB
	sync.RWMutex
}

func NewSQLite3Repo(dbfile string) (*dbRepo, error) {
	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetMaxOpenConns(1)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if _, err := db.Exec(createTableMetrics); err != nil {
		return nil, err
	}
	return &dbRepo{
		db: db,
	}, nil
}

func (r *dbRepo) Add(m models.Metric) error {
	// Create entry in the repository
	r.Lock()
	defer r.Unlock()

	// Prepare INSERT statements
	insStmt, err := r.db.Prepare("INSERT INTO metrics VALUES(NULL, ?,?,?)")
	if err != nil {
		return err
	}
	defer insStmt.Close()

	// Exec INSERT statements
	_, err = insStmt.Exec(m.TimeScrape, m.Ssid, m.Received)
	if err != nil {
		return err
	}

	return nil
}

func (r *dbRepo) GetMetricsFromString(day string) ([]models.Metric, error) {
	r.RLock()
	defer r.RUnlock()

	stmt := `SELECT * FROM metrics
	WHERE
	strftime('%Y%m%d', time_scrape, 'localtime')=
	strftime('%Y%m%d', ?, 'localtime')`

	rows, err := r.db.Query(stmt, day)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := []models.Metric{}
	for rows.Next() {
		i := models.Metric{}
		err := rows.Scan(&i.ID, &i.TimeScrape, &i.Ssid, &i.Received)

		if err != nil {
			return nil, err
		}

		data = append(data, i)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return data, nil
}
