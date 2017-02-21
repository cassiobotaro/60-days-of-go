package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jmoiron/sqlx"
	// Used pg drive on sqlx
	_ "github.com/lib/pq"
)

var (
	db   *sqlx.DB
	once sync.Once
)

// Singleton
// In software engineering, the singleton pattern is a software design pattern that restricts the instantiation of a class to one object.
// This is useful when exactly one object is needed to coordinate actions across the system.
// https://en.wikipedia.org/wiki/Singleton_pattern
// http://marcio.io/2015/07/singleton-pattern-in-go/

// Set database config
// export PGUSER=postgres
// export PGDB=postgres
// export PGHOST=localhost
// export PGPORT=5432
// Run postgresql inside a container
// docker run -d postgresql:latest

// MustGetConnection returns database connection
func MustGetConnection() *sqlx.DB {
	once.Do(func() {
		pguser := os.Getenv("PGUSER")
		pgdb := os.Getenv("PGDB")
		pghost := os.Getenv("PGHOST")
		pgport := os.Getenv("PGPORT")
		pgpass := os.Getenv("PGPASS")
		dbURI := fmt.Sprintf("user=%s dbname=%s host=%s port=%v sslmode=disable", pguser, pgdb, pghost, pgport)
		if pgpass != "" {
			dbURI += " password=" + pgpass
		}
		var err error
		db, err = sqlx.Connect("postgres", dbURI)
		if err != nil {
			panic(fmt.Sprintf("Unable to connection to database: %v\n", err))
		}
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(10)
	})
	return db
}

func main() {
	// Verify if connection is ok
	conn := MustGetConnection()
	err := conn.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected ✓")
}
