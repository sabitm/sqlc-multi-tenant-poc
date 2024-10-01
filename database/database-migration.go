package database

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"sort"
	"strings"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

func createMigrationsTable(db *sql.DB) error {
	_, err := db.Exec(`
    CREATE TABLE IF NOT EXISTS migrations (
      id INT AUTO_INCREMENT PRIMARY KEY,
      version VARCHAR(255) NOT NULL UNIQUE,
      applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
	`)
	return err
}

func runMigrations(db *sql.DB) error {
	migrations, err := migrationsFS.ReadDir("migrations")
	if err != nil {
		return err
	}

	var migrationFiles []string
	for _, migration := range migrations {
		if strings.HasSuffix(migration.Name(), ".sql") {
			migrationFiles = append(migrationFiles, migration.Name())
		}
	}

	sort.Strings(migrationFiles)

	for _, filename := range migrationFiles {
		version := strings.TrimSuffix(filename, ".sql")

		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM migrations WHERE version = ?", version).Scan(&count)
		if err != nil {
			return err
		}

		if count > 0 {
			continue
		}

		content, err := migrationsFS.ReadFile(fmt.Sprintf("migrations/%s", filename))
		if err != nil {
			return err
		}

		_, err = db.Exec(string(content))
		if err != nil {
			return fmt.Errorf("error executing migration %s: %v", version, err)
		}

		_, err = db.Exec("INSERT INTO migrations (version) VALUES (?)", version)
		if err != nil {
			return err
		}

		log.Printf("Applied migration: %s\n", version)
	}

	return nil
}
