package main

import (
	"database/sql"
	"net/http"

	"github.com/codegangsta/martini"
	_ "github.com/lib/pq"
	"github.com/martini-contrib/render"
)

func SetupDB() *sql.DB {
	db, err := sql.Open("postgres", "dbname=lesson7 sslmode=disable")
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
	m.Use(render.Renderer())
	m.Map(SetupDB())

	m.Get("/login", GetLogin)
	m.Post("/login", PostLogin)

	m.Run()
}

func GetLogin(r render.Render) {
	r.HTML(200, "login", nil)
}

func PostLogin(req *http.Request, db *sql.DB) (int, string) {
	var id string

	email, password := req.FormValue("email"), req.FormValue("password")
	err := db.QueryRow("select id from users where email=$1 and password=$2", email, password).Scan(&id)

	if err != nil {
		return 401, "Unauthorized"
	}

	return 200, "User id is " + id
}
