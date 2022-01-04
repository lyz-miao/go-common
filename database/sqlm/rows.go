package sqlm

import "database/sql"

type Rows struct {
	rows *sql.Rows
}

type ColumnType struct {
	columnType *sql.ColumnType
}

func (r *Rows) Close() error {
	return r.rows.Close()
}
func (r *Rows) ColumnTypes() ([]*ColumnType, error) {
	c, err := r.rows.ColumnTypes()
	if err != nil {
		return nil, err
	}
	ct := make([]*ColumnType, 0, len(c))
	for _, i := range c {
		ct = append(ct, &ColumnType{columnType: i})
	}
	return ct, nil
}
func (r *Rows) Columns() ([]string, error) {
	return r.rows.Columns()
}
func (r *Rows) Err() error {
	return r.rows.Err()
}
func (r *Rows) Next() bool {
	return r.rows.Next()
}
func (r *Rows) NextResultSet() bool {
	return r.rows.NextResultSet()
}
func (r *Rows) Scan(dest ...interface{}) error {
	return r.rows.Scan(dest...)
}
