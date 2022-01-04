package sqlm

import (
	"context"
	"database/sql"
    "errors"
    "strconv"
)

type Tx struct {
	tx *sql.Tx
}

type TxOptions struct {
	txOptions *sql.TxOptions
}

const (
	LevelDefault IsolationLevel = iota
	LevelReadUncommitted
	LevelReadCommitted
	LevelWriteCommitted
	LevelRepeatableRead
	LevelSnapshot
	LevelSerializable
	LevelLinearizable
)

type IsolationLevel int

func (i IsolationLevel) String() string {
	switch i {
	case LevelDefault:
		return "Default"
	case LevelReadUncommitted:
		return "Read Uncommitted"
	case LevelReadCommitted:
		return "Read Committed"
	case LevelWriteCommitted:
		return "Write Committed"
	case LevelRepeatableRead:
		return "Repeatable Read"
	case LevelSnapshot:
		return "Snapshot"
	case LevelSerializable:
		return "Serializable"
	case LevelLinearizable:
		return "Linearizable"
	default:
		return "IsolationLevel(" + strconv.Itoa(int(i)) + ")"
	}
}

func (t *Tx) Insert(ctx context.Context, table string, field []string, values [][]interface{}) (Result, error) {
	query, args := genInsertParam(table, field, values)
	return t.Exec(ctx, query, args...)
}

func (t *Tx) Commit() error {
	err := t.tx.Commit()
    if errors.Is(err, sql.ErrTxDone){
        return ErrTxDone
    }
    return err
}

func (t *Tx) Exec(ctx context.Context, query string, args ...interface{}) (Result, error) {
	return t.tx.ExecContext(ctx, query, args...)
}
func (t *Tx) Prepare(ctx context.Context, query string) (*Stmt, error) {
	s, err := t.tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	return &Stmt{stmt: s}, nil
}

func (t *Tx) Query(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	rows, err := t.tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &Rows{rows: rows}, nil
}
func (t *Tx) QueryRow(ctx context.Context, query string, args ...interface{}) *Row {
	return &Row{t.tx.QueryRowContext(ctx, query, args...)}
}
func (t *Tx) Rollback() error {
	err := t.tx.Rollback()
    if errors.Is(err, sql.ErrTxDone){
        return ErrTxDone
    }
    return err
}
func (t *Tx) Stmt(ctx context.Context, stmt *Stmt) *Stmt {
	return &Stmt{t.tx.StmtContext(ctx, stmt.stmt)}
}
