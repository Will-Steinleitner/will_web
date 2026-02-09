package main

import (
	"will_web"
	"will_web/internal/controllers"
)

func main() {
	application := will_web.Application{}
	homeCtrl := &controllers.HomeController{HomeScreenMR: application.HomeScreenModelRepo}
}
