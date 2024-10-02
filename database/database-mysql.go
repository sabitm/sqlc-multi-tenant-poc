package database

import (
	"database/sql"
	_ "embed"
	"fmt"
	"os"
	"project/compiled"
	"time"
)

var DB DBStruct

type DBStruct struct {
	Conn  *WrapperDB
	Query *compiled.Queries
}

func (db *DBStruct) ConnectMySQL() (*WrapperDB, error) {
	if _, err := os.Stat("./data"); os.IsNotExist(err) {
		_ = os.Mkdir("./data", 0755)
	}

	source := "root:%s@/xxxx?parseTime=true"
	source = fmt.Sprintf(source, os.Getenv("MYSQL_ROOT_PASSWORD"))

	dbCurrent, err := sql.Open("mysql", source)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}
	db.Conn = &WrapperDB{
		db: dbCurrent,
	}
	if err := db.Conn.Ping(); err != nil {
		return nil, fmt.Errorf("unable to ping to database: %w", err)
	}

	dbCurrent.SetConnMaxLifetime(time.Minute * 3)
	dbCurrent.SetMaxOpenConns(10)
	dbCurrent.SetMaxIdleConns(10)

	err = createMigrationsTable(dbCurrent, "xxxx")
	if err != nil {
		panic(err)
	}
	err = runMigrations(dbCurrent, "xxxx")
	if err != nil {
		panic(err)
	}
	db.Query = compiled.New(db.Conn)
	return db.Conn, nil
}
