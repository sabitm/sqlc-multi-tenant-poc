package database

import (
	"database/sql"
	_ "embed"
	"fmt"
	"os"
	"project/compiled"
	"time"
)

func (db *DBStruct) ConnectMySQL() (*sql.DB, error) {
	if _, err := os.Stat("./data"); os.IsNotExist(err) {
		_ = os.Mkdir("./data", 0755)
	}

	source := "root:%s@/core?parseTime=true"
	source = fmt.Sprintf(source, os.Getenv("MYSQL_ROOT_PASSWORD"))

	dbCurrent, err := sql.Open("mysql", source)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}
	db.Conn = dbCurrent
	if err := db.Conn.Ping(); err != nil {
		return nil, fmt.Errorf("unable to ping to database: %w", err)
	}

	dbCurrent.SetConnMaxLifetime(time.Minute * 3)
	dbCurrent.SetMaxOpenConns(10)
	dbCurrent.SetMaxIdleConns(10)

	_, err = dbCurrent.Exec("USE core;")
	if err != nil {
		return nil, fmt.Errorf("unable to use database: %w", err)
	}

	err = createMigrationsTable(dbCurrent)
	if err != nil {
		panic(err)
	}
	err = runMigrations(dbCurrent)
	if err != nil {
		panic(err)
	}
	db.Query = compiled.New(dbCurrent)
	return db.Conn, nil
}
