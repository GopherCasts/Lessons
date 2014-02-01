package main

import (
	"database/sql"
	"fmt"
	"github.com/codegangsta/martini"
	_ "github.com/lib/pq"
	"net/http"
)

func SetupDB() *sql.DB {
	db, err := sql.Open("postgres", "dbname=lesson4 sslmode=disable")
	PanicIf(err)

	return db
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
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
                           OR description ILIKE $1`, search)
		PanicIf(err)
		defer rows.Close()

		var title, author, description string
		for rows.Next() {
			err := rows.Scan(&title, &author, &description)
			PanicIf(err)
			fmt.Fprintf(rw, "Title: %s\nAuthor: %s\nDescription: %s\n\n", title, author, description)
		}
	})

	m.Run()
}
