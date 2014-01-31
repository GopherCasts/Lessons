package main

import (
	"database/sql"
	"github.com/codegangsta/martini"
	_ "github.com/lib/pq"
	"log"
)

func SetupDB() *sql.DB {
	db, err := sql.Open("postgres", "dbname=lesson4 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// Does the Database Exist?
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func main() {
	m := martini.Classic()

	m.Map(SetupDB())

	m.Get("/", func(db *sql.DB) string {
		var (
			id int
			title string
			author string
			description string
			output string
		)

		rows, err := db.Query("SELECT id, title, author, description FROM books")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&id, &title, &author, &description)
			if err != nil {
				log.Fatal(err)
			}
			output += title + "\n"
		}

		return output
	})

	m.Run()

}
