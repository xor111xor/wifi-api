package repository_test

import (
	"os"
	"testing"
	"time"

	"github.com/xor111xor/wifi-api/internal/models"
	"github.com/xor111xor/wifi-api/internal/repository"
)

func getRepoSqlite3(t *testing.T) (models.Repo, func()) {
	t.Helper()

	tf, err := os.CreateTemp("/tmp", "wifi-sql")
	if err != nil {
		t.Fatal(err)
	}
	// tf.Close()

	repo, err := repository.NewSQLite3Repo(tf.Name())
	if err != nil {
		t.Error(err)
	}
	return repo, func() {}
}

func TestAddSqlite3(t *testing.T) {
	repo, cleanup := getRepoSqlite3(t)
	defer cleanup()

	testCases := []struct {
		name       string
		formatDate string
		set        models.Metric
		exp        []models.Metric
	}{
		{
			name:       "Add_manual",
			formatDate: "2022-11-17",
			set: models.Metric{
				TimeScrape: time.Date(2022, 11, 17, 20, 34, 58, 651387237, time.UTC),
				Ssid:       "Wopen",
				Received:   50,
			},
			exp: []models.Metric{models.Metric{
				ID:         1,
				TimeScrape: time.Date(2022, 11, 17, 20, 34, 58, 651387237, time.UTC),
				Ssid:       "Wopen",
				Received:   50,
			}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := repo.Add(tc.set)
			if err != nil {
				t.Fatal(err)
			}
			got, err := repo.GetMetricsFromString(tc.formatDate)
			if err != nil {
				t.Fatal(err)
			}

			if len(got) != len(tc.exp) {
				t.Errorf("Expected length slice %q, got %q.\n", tc.exp, got)
			}
			for i := range got {
				if got[i] != tc.exp[i] {
					t.Errorf("Expected entity of slice %q, got %q.\n", tc.exp[i], got[i])
				}
			}
		})
	}
}
