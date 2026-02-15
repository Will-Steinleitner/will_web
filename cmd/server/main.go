package main

import (
	"html/template"
	"log"
	"net/http"
	application "will_web"
	"will_web/internal/controllers"

	_ "github.com/lib/pq"
)

const mainTag = "Main"

func main() {
	app := application.NewApplication()
	defer app.Database().Close()
	fullCache := app.TemplateCache()

	templates := map[string]*template.Template{
		"home.html": fullCache["home.html"],
	}

	//bonus6725
	homeCtrl := controllers.NewHomeScreenController(templates, app.HomeRepo())

	fs := http.FileServer(http.Dir("./ui/static/"))
	// http.Handle registers a handler for a specific URL pattern,
	// e.g. it routes requests like GET /static/home.css to the file server.
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// handle force you to implement ServeHTTP
	http.Handle("/", homeCtrl)

	log.Println("Server startet auf http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(mainTag, err)
	}
}
