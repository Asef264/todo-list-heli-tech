package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"todo-list/config"

	_ "github.com/lib/pq"
)

func NewDB(cfg *config.Config) (*sql.DB, error) {
	dbInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.DBName,
		cfg.DB.SSLMode,
	)

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func MigrateUp(db *sql.DB, filesPath, dbName string) {
	migrationFiles, err := getMigrationFiles("../migrations")
	if err != nil {
		log.Fatalf("error reading migration files: %v", err)
	}

	for _, file := range migrationFiles {
		err := applyMigration(db, file)
		if err != nil {
			log.Fatalf("error applying migration %s: %v", file, err)
		}
		log.Printf("Migration %s applied successfully.", file)
	}

	log.Println("Migration applied successfully.")
}

func getMigrationFiles(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".sql" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, nil
}

func applyMigration(db *sql.DB, file string) error {
	migrationContent, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("error reading migration file %s: %v", file, err)
	}

	_, err = db.Exec(string(migrationContent))
	if err != nil {
		return fmt.Errorf("error executing migration %s: %v", file, err)
	}

	return nil
}
