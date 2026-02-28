package store

import (
	"database/sql"
	"fmt"
	"io/fs"
	"training/config"

	"github.com/pressly/goose/v3"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func Open() (*sql.DB, error) {
	str := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.DB_HOST(),
		config.DB_USER(),
		config.DB_PASSWORD(),
		config.DB_NAME(),
		config.DB_PORT(),
		"disable",
	)
	db, err := sql.Open("pgx", str)
	if err != nil {
		return nil, fmt.Errorf("db: open %w", err)
	}
	fmt.Println("Connected to database...")
	return db, nil
}

func MigrateFS(db *sql.DB, migrationsFS fs.FS, dir string) error {
	goose.SetBaseFS(migrationsFS)
	defer func() {
		goose.SetBaseFS(nil)
	}()
	return Migrate(db, dir)
}

func Migrate(db *sql.DB, dir string) error {
	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("migrate: %w", err)
	}

	err = goose.Up(db, dir)
	if err != nil {
		return fmt.Errorf("goose: %w", err)
	}
	return nil
}
