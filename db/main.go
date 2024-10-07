package db

import (
	"database/sql"

	migrate "github.com/rubenv/sql-migrate"
)

func NewDB(
	path string,
	migrationsPath string,
) (*sql.DB, error) {
	dbi, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	migrations := &migrate.FileMigrationSource{
		Dir: migrationsPath,
	}
	_, err = migrate.Exec(dbi, "sqlite3", migrations, migrate.Up)
	if err != nil {
		return nil, err
	}

	return dbi, nil
}
