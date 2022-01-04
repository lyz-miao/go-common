package sqlm

import (
	"context"
	"database/sql"
)

type Stmt struct {
	stmt *sql.Stmt
}

func (s *Stmt) Close() error {
	return s.stmt.Close()
}

func (s *Stmt) Exec(ctx context.Context, args ...interface{}) (Result, error) {
	return s.stmt.ExecContext(ctx, args...)
}

func (s *Stmt) Query(ctx context.Context, args ...interface{}) (*Rows, error) {
	rows, err := s.stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	return &Rows{rows: rows}, nil
}

func (s *Stmt) QueryRow(ctx context.Context, args ...interface{}) *Row {
	row := s.stmt.QueryRowContext(ctx, args...)
	return &Row{row: row}
}
