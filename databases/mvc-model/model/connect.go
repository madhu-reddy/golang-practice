package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB // variable starts with small "c", so only the files in the package "model" has access to this global variable, if it starts with capital "C", then all the packages would have access to this variable.

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:welcome123@tcp(localhost:3306)/testmadhu")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")
	con = db
	return db
}
