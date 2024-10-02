package database

import (
	"context"
	"database/sql"
	"log"
	"strings"
)

type TenantContextKey struct{}

type WrapperDB struct {
	db *sql.DB
}

func (wdb *WrapperDB) Ping() error {
	return wdb.db.Ping()
}

func (wdb *WrapperDB) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	tenant := ctx.Value(TenantContextKey{}).(string)
	query = strings.ReplaceAll(query, "xxxx", tenant)
	log.Println("Query: ", query)
	return wdb.db.ExecContext(ctx, query, args...)
}

func (wdb *WrapperDB) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	tenant := ctx.Value(TenantContextKey{}).(string)
	query = strings.ReplaceAll(query, "xxxx", tenant)
	log.Println("Query: ", query)
	return wdb.db.PrepareContext(ctx, query)
}

func (wdb *WrapperDB) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	tenant := ctx.Value(TenantContextKey{}).(string)
	query = strings.ReplaceAll(query, "xxxx", tenant)
	log.Println("Query: ", query)
	return wdb.db.QueryContext(ctx, query, args...)
}

func (wdb *WrapperDB) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	tenant := ctx.Value(TenantContextKey{}).(string)
	query = strings.ReplaceAll(query, "xxxx", tenant)
	log.Println("Query: ", query)
	return wdb.db.QueryRowContext(ctx, query, args...)
}
