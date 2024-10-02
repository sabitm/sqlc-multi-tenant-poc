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

func createMigrationsTable(db *sql.DB, tenantID string) error {
	_, err := db.Exec(fmt.Sprintf(`
    CREATE DATABASE IF NOT EXISTS %s;
	`, tenantID))
	if err != nil {
		return err
	}

	_, err = db.Exec(fmt.Sprintf(`
    CREATE TABLE IF NOT EXISTS %s.migrations (
      id INT AUTO_INCREMENT PRIMARY KEY,
      version VARCHAR(255) NOT NULL UNIQUE,
      applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
	`, tenantID))
	return err
}

func runMigrations(db *sql.DB, tenantID string) error {
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
		countQuery := "SELECT COUNT(*) FROM %s.migrations WHERE version = ?"
		err := db.QueryRow(fmt.Sprintf(countQuery, tenantID), version).Scan(&count)
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

		insertQuery := "INSERT INTO %s.migrations (version) VALUES (?)"
		_, err = db.Exec(fmt.Sprintf(insertQuery, tenantID), version)
		if err != nil {
			return err
		}

		log.Printf("Applied migration: %s\n", version)
	}

	return nil
}
