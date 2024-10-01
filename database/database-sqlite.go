package database

import (
	"database/sql"
	_ "embed"
	"fmt"
	"os"
	"project/compiled"
)

var DB DBStruct

type DBStruct struct {
	Conn  *sql.DB
	Query *compiled.Queries
}

func (db *DBStruct) ConnectSqlite() (*sql.DB, error) {
	if _, err := os.Stat("./data"); os.IsNotExist(err) {
		_ = os.Mkdir("./data", 0755)
	}
	dbCurrent, err := sql.Open("sqlite", "./data/database.db")
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}
	db.Conn = dbCurrent
	if err := db.Conn.Ping(); err != nil {
		return nil, fmt.Errorf("unable to ping to database: %w", err)
	}

	_ = createMigrationsTable(dbCurrent)
	err = runMigrations(dbCurrent)
	if err != nil {
		panic(err)
	}
	db.Query = compiled.New(dbCurrent)
	return db.Conn, nil
}
