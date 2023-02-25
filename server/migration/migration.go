package migration

import (
	"codeview/config"
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func createMigrationInstance() (*migrate.Migrate, error) {
	cfg := config.Init()

	cfg.LoadFromEnv()

	dataSourceUrl := cfg.Postgres.GetDatabaseURL()
	db, err := sql.Open("postgres", dataSourceUrl)
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		cfg.Migration.FilesPath,
		"postgres", driver)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func MigrateUp() {
	m, err := createMigrationInstance()
	if err != nil {
		log.Printf("Error creating postgres instance: %v\n", err)
		return
	}

	if err := m.Up(); err != nil {
		log.Printf("Error migrating: %v\n", err)
		return
	}

	log.Printf("Success migrating up database [UP-ALL]\n")
}

func MigrateDown() {
	m, err := createMigrationInstance()
	if err != nil {
		log.Printf("Error creating postgres instance: %v\n", err)
		return
	}

	if err := m.Down(); err != nil {
		log.Printf("Error migrating: %v\n", err)
		return
	}

	log.Printf("Success migrating down database [DOWN-ALL]\n")
}

func MigrateStepDown() {
	m, err := createMigrationInstance()
	if err != nil {
		log.Printf("Error creating postgres instance: %v\n", err)
		return
	}

	if err := m.Steps(-1); err != nil {
		log.Printf("Error migrating: %v\n", err)
		return
	}

	log.Printf("Success migrating down database [DOWN-1]\n")
}

func MigrateStepUp() {
	m, err := createMigrationInstance()
	if err != nil {
		log.Printf("Error creating postgres instance: %v\n", err)
		return
	}

	if err := m.Steps(1); err != nil {
		log.Printf("Error migrating: %v\n", err)
		return
	}

	log.Printf("Success migrating up database [UP-1]\n")

}
