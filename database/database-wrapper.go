package database

import (
	"context"
	"database/sql"
)

type WrapperDB struct {
	db *sql.DB
}

func (wdb *WrapperDB) Ping() error {
	return wdb.db.Ping()
}

func (wdb *WrapperDB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return wdb.db.ExecContext(ctx, query, args)
}

func (wdb *WrapperDB) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return wdb.db.PrepareContext(ctx, query)
}

func (wdb *WrapperDB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return wdb.db.QueryContext(ctx, query, args)
}

func (wdb *WrapperDB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return wdb.db.QueryRowContext(ctx, query, args)
}
