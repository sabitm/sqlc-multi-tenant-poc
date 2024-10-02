package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type TenantContextKey struct{}

func ReplaceQueryTenant(ctx context.Context, query string) string {
	tenant := ctx.Value(TenantContextKey{}).(string)
	tenant = fmt.Sprintf("`%s`", tenant)
	query = strings.ReplaceAll(query, "xxxx", tenant)
	log.Println("Query: ", query)
	return query
}

type WrapperDB struct {
	db *sql.DB
}

func (wdb *WrapperDB) Ping() error {
	return wdb.db.Ping()
}

func (wdb *WrapperDB) Exec(query string, args ...any) (sql.Result, error) {
	return wdb.db.Exec(query, args...)
}

func (wdb *WrapperDB) QueryRow(query string, args ...any) *sql.Row {
	return wdb.db.QueryRow(query, args...)
}

func (wdb *WrapperDB) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	query = ReplaceQueryTenant(ctx, query)
	return wdb.db.ExecContext(ctx, query, args...)
}

func (wdb *WrapperDB) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	query = ReplaceQueryTenant(ctx, query)
	return wdb.db.PrepareContext(ctx, query)
}

func (wdb *WrapperDB) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	query = ReplaceQueryTenant(ctx, query)
	return wdb.db.QueryContext(ctx, query, args...)
}

func (wdb *WrapperDB) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	query = ReplaceQueryTenant(ctx, query)
	return wdb.db.QueryRowContext(ctx, query, args...)
}
