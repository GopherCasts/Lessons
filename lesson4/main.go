package main

import (
	"fmt"
	"database/sql"
	"net/http"
	"github.com/codegangsta/martini"
	_ "github.com/lib/pq"
)

func SetupDB() *sql.DB {
	db, err := sql.Open("postgres", "dbname=lesson4 sslmode=disable")
	if err != nil {
		panic(err)
	}

	return db
}

func main() {
	m := martini.Classic()
	m.Map(SetupDB())

	m.Get("/", func(db *sql.DB, r *http.Request, rw http.ResponseWriter) {
		search := "%" + r.URL.Query().Get("search") + "%"
		rows, err := db.Query(`SELECT title, author, description 
                           FROM books 
                           WHERE title ILIKE $1
                           OR author ILIKE $1
                           OR description ILIKE $1;`, search)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		var title, author, description string
		for rows.Next() {
			err := rows.Scan(&title, &author, &description)
			if err != nil {
				panic(err)
			}
			fmt.Fprintf(rw, "Title: %s\nAuthor: %s\nDescription: %s\n\n", title, author, description)
		}
	})

	m.Run()

}
