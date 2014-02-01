package main

import (
	"database/sql"
	"fmt"
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

func FormatBook(title, author, description string) string {

	return fmt.Sprintf("Title: %s\nAuthor: %s\nDescription: %s\n\n", title, author, description)
}

func main() {
	m := martini.Classic()

	m.Map(SetupDB())

	m.Get("/", func(db *sql.DB) string {

		var title, author, description, output string

		rows, err := db.Query("SELECT title, author, description FROM books")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&title, &author, &description)
			if err != nil {
				log.Fatal(err)
			}
			output += FormatBook(title, author, description)
		}

		if output == "" {
			output = "No Books Found!"
		}

		return output
	})

	m.Get("/:query", func(db *sql.DB, params martini.Params) string {

		var title, author, description, output string

		queryString := "SELECT title, author, description "
		queryString += "FROM books "
		queryString += "WHERE title ILIKE '%" + params["query"] + "%' "
		queryString += "OR author ILIKE '%" + params["query"] + "%' "
		queryString += "OR description ILIKE '%" + params["query"] + "%'"

		rows, err := db.Query(queryString)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&title, &author, &description)
			if err != nil {
				log.Fatal(err)
			}
			output += FormatBook(title, author, description)
		}

		if output == "" {
			output = "No Books Found!"
		}

		return output
	})

	m.Run()

}
