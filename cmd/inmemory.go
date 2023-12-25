//go:build inmemory
package cmd

import (
	"github.com/xor111xor/wifi-api/internal/models"
	"github.com/xor111xor/wifi-api/internal/repository"
)

func GetRepo() (models.Repo, error) {
	repo, err := repository.NewInmemoryRepo()
	if err != nil {
		return nil, err
	}
	return repo, nil
}
