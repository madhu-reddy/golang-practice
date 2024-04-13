package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:welcome123@tcp(127.0.0.1:3306)/testmadhu")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//fmt.Println(err)
	//fmt.Println(db.Ping())

	createTableQuery := `
        CREATE TABLE IF NOT EXISTS users (
            id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(50),
            email VARCHAR(100)
        )
    `
	_, err = db.Exec(createTableQuery)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Table created successfully")

	// Insert data into the table
	insertQuery := "INSERT INTO users (name, email) VALUES (?, ?)"
	_, err = db.Exec(insertQuery, "John Doe", "john@example.com")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Data inserted successfully")

}


/*
go get -u github.com/go-sql-driver/mysql

func Open(driverName, dataSourceName string) (*DB, error)

type DB struct {
	// contains filtered or unexported fields
}

func (db *DB) Exec(query string, args ...any) (Result, error)

*/
