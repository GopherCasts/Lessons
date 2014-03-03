package main

import (
	"database/sql"
	"net/http"

	"github.com/codegangsta/martini"
	_ "github.com/lib/pq"
	"github.com/martini-contrib/sessions"
)

type User struct {
	Name  string
	Email string
}

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
	m.Map(SetupDB())

	// Sessions
	store := sessions.NewCookieStore([]byte("secret123"))
	m.Use(sessions.Sessions("lesson8", store))

	m.Get("/", RequireLogin, SecretPath)
	m.Post("/login", PostLogin)
	m.Get("/logout", Logout)

	m.Run()
}

func Logout(rw http.ResponseWriter, req *http.Request, s sessions.Session) {
	s.Delete("userId")
	http.Redirect(rw, req, "/", http.StatusFound)
}

func RequireLogin(rw http.ResponseWriter, req *http.Request,
	s sessions.Session, db *sql.DB, c martini.Context) {
	user := &User{}
	err := db.QueryRow("select name, email from users where id=$1", s.Get("userId")).Scan(&user.Name, &user.Email)

	if err != nil {
		http.Redirect(rw, req, "/login", http.StatusFound)
		return
	}

	c.Map(user)
}

func SecretPath(user *User) string {
	return "Great Job " + user.Name
}

func PostLogin(req *http.Request, db *sql.DB, s sessions.Session) (int, string) {
	var id string

	email, password := req.FormValue("email"), req.FormValue("password")
	err := db.QueryRow("select id from users where email=$1 and password=$2", email, password).Scan(&id)

	if err != nil {
		return 401, "Unauthorized"
	}

	s.Set("userId", id)

	return 200, "User id is " + id
}
