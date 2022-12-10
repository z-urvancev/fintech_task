package impl

import (
	"database/sql"
	"fintech/pkg/errors"
	"github.com/jmoiron/sqlx"
)

type PostgresRepo struct {
	db *sqlx.DB
}

func NewPostgresRepo(db *sqlx.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (pg *PostgresRepo) GetByShort(short string) (string, error) {
	query := `SELECT short FROM links WHERE url=$1`
	row := pg.db.QueryRow(query, short)
	var result string
	err := row.Scan(&result)
	if err == sql.ErrNoRows {
		return "", errors.ErrURLNotFound
	}
	return result, err
}

func (pg *PostgresRepo) GetByURL(url string) (string, error) {
	query := `SELECT short FROM links WHERE url=$1`
	row := pg.db.QueryRow(query, url)
	var short string
	err := row.Scan(&short)
	if err == sql.ErrNoRows {
		return "", errors.ErrURLNotFound
	}
	return short, err
}

func (pg *PostgresRepo) Insert(url, short string) error {
	query := `INSERT INTO links (short, url) values ($1, $2) RETURNING id`
	_, err := pg.db.Exec(query, short, url)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}
