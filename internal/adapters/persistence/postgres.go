package persistence

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

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// MigrateUp is a function to migrate the database
func MigrateUp(db *sql.DB, filesPath, dbName string) {
	// Read all migration files from the /migrations directory
	migrationFiles, err := getMigrationFiles("./migrations")
	if err != nil {
		log.Fatalf("error reading migration files: %v", err)
	}

	// Apply each migration file
	for _, file := range migrationFiles {
		err := applyMigration(db, file)
		if err != nil {
			log.Fatalf("error applying migration %s: %v", file, err)
		}
		log.Printf("Migration %s applied successfully.", file)
	}

	log.Println("Migration applied successfully.")
}

// getMigrationFiles returns a sorted list of migration file paths from the migrations directory.
func getMigrationFiles(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Check if it's a .sql file
		if filepath.Ext(path) == ".sql" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Optionally, sort files in a way you want (e.g., lexicographically)
	// sort.Strings(files)

	return files, nil
}

// applyMigration reads a SQL migration file and executes its contents.
func applyMigration(db *sql.DB, file string) error {
	// Read the migration file content
	migrationContent, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("error reading migration file %s: %v", file, err)
	}

	// Execute the migration SQL query
	_, err = db.Exec(string(migrationContent))
	if err != nil {
		return fmt.Errorf("error executing migration %s: %v", file, err)
	}

	return nil
}
