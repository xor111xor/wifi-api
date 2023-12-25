//go:build sqlite

package cmd

import (
	"github.com/xor111xor/wifi-api/internal/models"
	"github.com/xor111xor/wifi-api/internal/repository"
)

func GetRepo() (models.Repo, error) {
	repo, err := repository.NewSQLite3Repo("sqlite.db")
	if err != nil {
		return nil, err
	}
	return repo, nil
}
