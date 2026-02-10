package main

import (
	"will_web"
	"will_web/internal/controllers"
)

func main() {
	app := will_web.NewApplication()
	homeCtrl := controllers.NewHomeScreenController(app.HomeRepo())
}
