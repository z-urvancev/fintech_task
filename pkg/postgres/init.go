package postgres

import (
	"fintech/config"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func NewPostgresDB(cfg *config.DBConfig) (*sqlx.DB, error) {
	db, openDBErr := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if openDBErr != nil {
		return nil, openDBErr
	}

	return db, nil
}
