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
	fullCache := app.TemplateCache()

	defer app.Database().Close()

	templates := map[string]*template.Template{
		"home.html": fullCache["home.html"],
	}

	homeCtrl := controllers.NewHomeScreenController(templates, app.HomeRepo())

	http.Handle("/", homeCtrl)

	log.Println("Server startet auf http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(mainTag, err)
	}
}
