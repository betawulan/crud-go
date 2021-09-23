package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // iki namane side import, enek underscore, artinya dia autoload pertama kali, otomatis jalan dulauan sebelum codenya
	// panic ki di nggo setup,
	// ki kan butuh sql.Open ora oleh error, makane yen errror langsung dipateni app ne, kode dibawah e ga dijalankan
)

func main() {
	db, err := sql.Open("mysql", "beta:12345beta@tcp(127.0.0.1:3306)/kuliah")
	defer db.Close() // artinya defer, perintah dilakukan setelah function selesai, brarti setlah main selesei, db.Close dijalankan

	if err != nil {
		log.Fatal(err) // log.Fatal, ingat yen gunake fatal, semisal error, aplikasine langsung mati
	}

	sql := "INSERT INTO mhs(nama, email) VALUES ('tono', 'tono@gmail.com')"

	res, err := db.Exec(sql)
	if err != nil { // biasakan if err, ki ngisor e langsung, ga usah nganggo enter, kan iseh sepaket
		panic(err.Error())
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("the recent inserted row nim: %d\n", lastId)
}
