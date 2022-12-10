package impl

import (
	"fintech/config"
	"fintech/internal/repository"
	"fintech/pkg/errors"
	"fintech/pkg/postgres"
)

func InitRepository(storeType string, config *config.Config) (repository.Repository, error) {
	if storeType == "inMemory" {
		imStore := NewInMemoryRepo(make(map[string]string))
		return imStore, nil
	}
	if storeType == "postgres" {
		db, dbErr := postgres.NewPostgresDB(&config.DB)
		if dbErr != nil {
			return nil, dbErr
		}
		pgStore := NewPostgresRepo(db)
		return pgStore, nil
	}
	return nil, errors.ErrIncorrectStoreType
}
