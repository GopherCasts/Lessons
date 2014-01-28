package main

import (
	"github.com/codegangsta/martini"
	"github.com/russross/blackfriday"
	"net/http"
)

func main() {
	m := martini.Classic()

	m.Post("/generate", func(r *http.Request) []byte {
		body := r.FormValue("body")
		return blackfriday.MarkdownBasic([]byte(body))
	})

	m.Run()
}
