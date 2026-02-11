package main

import (
	"html/template"
	"log"
	"net/http"
	"will_web/internal/controllers"
)

func main() {
	app := NewApplication()
	fullCache := app.TemplateCache()

	homeTemplate := map[string]*template.Template{
		"home.html": fullCache["home.html"],
	}
	homeCtrl := controllers.NewHomeScreenController(homeTemplate, app.HomeRepo())

	http.Handle("/", homeCtrl)

	log.Println("Server startet auf http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
