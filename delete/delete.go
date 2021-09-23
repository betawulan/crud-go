package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	id    int
	nama  string
	email string
}

func main() {
	db, e := sql.Open("mysql", "beta:12345beta@tcp(127.0.0.1:3306)/kuliah")
	ErrorCheck(e)

	defer db.Close()

	PingDB(db)

	// delete data
	stmt, e := db.Prepare("DELETE FROM mhs WHERE id=?")
	ErrorCheck(e)

	// delete 4st student
	res, e := stmt.Exec("4")
	ErrorCheck(e)

	// affected rows
	a, e := res.RowsAffected()
	ErrorCheck(e)

	fmt.Println(a)
}

func ErrorCheck(err error) {
	if err != nil {
		panic((err.Error()))
	}
}

func PingDB(db *sql.DB) {
	err := db.Ping()
	ErrorCheck(err)
}
