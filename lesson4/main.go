package main

import (
	"database/sql"
	//"github.com/codegangsta/martini"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func main() {
	//m := martini.Classic()

	//m.Run()

	var err error
	db, err = sql.Open("postgres", "dbname=lesson4 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// Does the Database Exist?
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	var (
		id int
		title string
		output string
	)

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &title)
		if err != nil {
			log.Fatal(err)
		}
		output += title
	}

	log.Println(output)
}
