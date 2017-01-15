package main

import (
	"fmt"
	"log"

	"github.com/cassiobotaro/60-days-of-go/day15/database"
)

// This example is based on https://www.goinggo.net/2013/07/singleton-design-pattern-in-go.html
// also https://github.com/mattn/go-sqlite3/blob/master/_example/simple/simple.go

// CreateTable creates table foo without data
func CreateTable() {
	// Get an instance from db, not a connection
	// instance should be unique
	db := database.Get()
	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`
	// get a connection from pool and execute some statement
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

// PopulateTable insert some records in table foo
func PopulateTable() {
	// Get an instance from db, not a connection
	// instance should be unique
	db := database.Get()
	// start a trasanction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	// prepare the query that should be executed
	stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	// insert 100 records on database
	for i := 0; i < 100; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("Hello World %03d", i))
		if err != nil {
			log.Fatal(err)
		}
	}
	// finally commit
	tx.Commit()
}

// ListFoo list all records from foo
func ListFoo() {
	// Get an instance from db, not a connection
	// instance should be unique
	db := database.Get()
	// start a transaction
	tx, err := db.Begin()
	// execute a query
	rows, err := tx.Query("select id, name from foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	// read the rows returned
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// initialize the database
	database.MustConnect()
	CreateTable()
	PopulateTable()
	ListFoo()
}
