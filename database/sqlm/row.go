package sqlm

import (
    "database/sql"
    "errors"
)

type Row struct {
	row *sql.Row
}

func (r *Row) Err() error {
    err := r.row.Err()
    if errors.Is(err, sql.ErrNoRows){
        return ErrNoRows
    }
	return err
}
func (r *Row) Scan(dest ...interface{}) error {
	return r.row.Scan(dest...)
}
