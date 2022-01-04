package sqlm

import (
	"context"
	"database/sql"
	"time"
)

type DB struct {
	db *sql.DB
}

type DBStats struct {
	stats sql.DBStats
}

func (d *DB) Begin(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	t, err := d.db.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}

	return &Tx{t}, nil
}

func (d *DB) Transaction(ctx context.Context, f func(tx *Tx) (interface{}, error)) (interface{}, error) {
	tx, err := d.Begin(ctx, nil)
	if err != nil {
		return nil, err
	}

	r, err := f(tx)
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return nil, err
		}
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return nil, err
		}
		return nil, err
	}

	return r, nil
}

func (d *DB) Insert(ctx context.Context, table string, field []string, values [][]interface{}) (Result, error) {
	query, args := genInsertParam(table, field, values)
	return d.Exec(ctx, query, args...)
}

func (d *DB) Close() error {
	return d.db.Close()
}

func (d *DB) Exec(ctx context.Context, query string, args ...interface{}) (Result, error) {
	return d.db.ExecContext(ctx, query, args...)
}

func (d *DB) Ping(ctx context.Context) error {
	return d.db.PingContext(ctx)
}

func (d *DB) Prepare(ctx context.Context, query string) (*Stmt, error) {
	s, err := d.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	return &Stmt{stmt: s}, nil
}

func (d *DB) Query(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	rows, err := d.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &Rows{rows: rows}, nil
}

func (d *DB) QueryRow(ctx context.Context, query string, args ...interface{}) *Row {
	row := d.db.QueryRowContext(ctx, query, args...)
	return &Row{row: row}
}

func (d *DB) SetConnMaxIdleTime(t time.Duration) {
	d.db.SetConnMaxIdleTime(t)
}

func (d *DB) SetConnMaxLifetime(t time.Duration) {
	d.db.SetConnMaxLifetime(t)
}

func (d *DB) SetMaxIdleConns(n int) {
	d.db.SetMaxIdleConns(n)
}

func (d *DB) SetMaxOpenConns(n int) {
	d.SetMaxOpenConns(n)
}

func (d *DB) Stats() DBStats {
	s := d.db.Stats()
	return DBStats{stats: s}
}
