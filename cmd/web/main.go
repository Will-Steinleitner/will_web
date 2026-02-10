package main

import (
	"html/template"
	"will_web"
	"will_web/internal/controllers"
)

func main() {
	app := will_web.NewApplication()
	fullCache := app.TemplateCache()

	homeTemplate := map[string]*template.Template{
		"home.html": fullCache["home.html"],
	}
	homeCtrl := controllers.NewHomeScreenController(homeTemplate, app.HomeRepo())
}
