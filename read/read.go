package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	Id    int
	Nama  string
	Email string
}

func main() {
	db, err := sql.Open("mysql", "beta:12345beta@tcp(127.0.0.1:3306)/kuliah")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}
	res, err := db.Query("SELECT * FROM mhs")

	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		var student Student
		err := res.Scan(&student.Id, &student.Nama, &student.Email)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", student)
	}
}
