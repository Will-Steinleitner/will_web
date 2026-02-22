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
	homeCtrl := controllers.NewHomeScreenController(
		app.HomeRepo(),
		app.GetRenderer(),
		app.GetPasswordHasher(),
	)

	//can we refactor this ?
	fs := http.FileServer(http.Dir("./ui/static/"))
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.Handle("/", homeCtrl)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("%s: %v", mainTag, err)
	}

	defer app.Database().Close()
}
