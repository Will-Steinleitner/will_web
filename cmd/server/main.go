package main

import (
	"log"
	"net/http"
	application "will_web"
	"will_web/internal/controllers"

	_ "github.com/lib/pq"
)

const mainTag = "Main"

func main() {
	app := application.NewApplication()
	homeCtrl := controllers.NewHomeScreenController(app.HomeRepo(), app.GetRenderer())

	//can we refactor this ?
	fs := http.FileServer(http.Dir("./ui/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/", homeCtrl)

	log.Println("Server startet auf http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(mainTag, err)
	}

	defer app.Database().Close()
}
