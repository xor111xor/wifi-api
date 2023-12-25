//go:build boltdb
package cmd

import (
	"github.com/xor111xor/wifi-api/internal/models"
	"github.com/xor111xor/wifi-api/internal/repository"
)

func GetRepo() (models.Repo, error) {
	repo, err := repository.NewBoltRepo("bolt.db")
	if err != nil {
		return nil, err
	}
	return repo, nil
}
