package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Config struct {
	DSN         string
	Debug       bool
	MaxOpenConn int
}

func New(c *Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", c.DSN)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(c.MaxOpenConn)
	db.SetMaxIdleConns(c.MaxOpenConn)

	return db, nil
}
