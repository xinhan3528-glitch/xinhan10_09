package store

import (
	"database/sql"
	"fmt"
	"io/fs"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
)

func Open() (*sql.DB, error) {
	connectionString := "host=localhost user=user password=pass dbname=postgres port=5432 sslmode=disable"

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, fmt.Errorf("db: open error %w", err)
	}
	fmt.Println("已连接Database")

	return db, err
}

func MigrateFS(db *sql.DB, migrateFS fs.FS, dir string) error {
	goose.SetBaseFS(migrateFS)
	defer func() {
		goose.SetBaseFS(nil)
	}()
	return Migrate(db, dir)
}

func Migrate(db *sql.DB, dir string) error {
	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("db-migrate: error %w", err)
	}

	err = goose.Up(db, dir)
	if err != nil {
		return fmt.Errorf("goose-up: error %w", err)
	}
	return nil
}
