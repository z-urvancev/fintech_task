package impl

import (
	"github.com/jmoiron/sqlx"
)

type PostgresRepo struct {
	db *sqlx.DB
}

func NewPostgresRepo(db *sqlx.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (pg *PostgresRepo) GetByShort(short string) (string, error) {
	stmt, psErr := pg.db.Prepare(`SELECT url FROM links WHERE short=?`)
	if psErr != nil {
		return "", psErr
	}
	var result string
	row := stmt.QueryRow(short)
	if row.Err() != nil {
		return "", row.Err()
	}

	scanErr := row.Scan(&result)
	if scanErr != nil {
		return "", scanErr
	}
	return result, nil
}

func (pg *PostgresRepo) GetByURL(url string) (string, error) {
	stmt, psErr := pg.db.Prepare(`SELECT short FROM links WHERE url=?`)
	if psErr != nil {
		return "", psErr
	}
	var result string
	row := stmt.QueryRow(url)
	if row.Err() != nil {
		return "", row.Err()
	}

	scanErr := row.Scan(&result)
	if scanErr != nil {
		return "", scanErr
	}
	return result, nil
}

func (pg *PostgresRepo) Insert(url, short string) error {
	insertSchema := `INSERT INTO links (url, short) VALUES (?, ?)`
	_, err := pg.db.Exec(insertSchema, url, short)
	if err != nil {
		return err
	}
	return nil
}
