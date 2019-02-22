package main

import (
	"database/sql"
	"fmt"
)

func database() {
	fmt.Println("Go MySQL!!")

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println("Successfully Connected to MySQL database")
}
