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
	for k := range fullCache {
		log.Println("CACHE KEY:", k)
	}

	templates := map[string]*template.Template{
		"base.gohtml":     fullCache["base.gohtml"],
		"register.gohtml": fullCache["register.gohtml"],
		"index.gohtml":    fullCache["index.gohtml"],
	}

	homeCtrl := controllers.NewHomeScreenController(templates, app.HomeRepo(), app.GetRenderer())

	fs := http.FileServer(http.Dir("./ui/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/", homeCtrl)

	log.Println("Server startet auf http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(mainTag, err)
	}
}
