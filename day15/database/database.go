package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

// MustConnect initialize the database in memory
func MustConnect() {
	db = sqlx.MustConnect("sqlite3", ":memory:")
}

// Get returns the database connection
func Get() *sqlx.DB {
	return db
}
