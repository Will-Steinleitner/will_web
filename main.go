package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"will_web/internal/controllers"

	_ "github.com/lib/pq"
)

func main() {
	app := NewApplication()
	fullCache := app.TemplateCache()

	connectionString := "postgres://postgres:mysecretpassword@localhost:5432/gopg?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	createTable(db)
	user := User{"Test", "123"}
	id := insertUser(db, user)
	fmt.Printf("ID = %d", id)

	homeTemplate := map[string]*template.Template{
		"home.html": fullCache["home.html"],
	}
	homeCtrl := controllers.NewHomeScreenController(homeTemplate, app.HomeRepo())

	//Handle force you to implement ServHTTP
	http.Handle("/", homeCtrl)

	log.Println("Server startet auf http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func createTable(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        password VARCHAR(255) NOT NULL,
        created TIMESTAMP DEFAULT NOW()
    );
    `
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

type User struct {
	Name string
	Pass string
}

func insertUser(db *sql.DB, user User) int {
	//($1, $2) sind die datenn von unserem User - siehe struct
	query := `INSERT INTO users (name, password) 
			VALUES ($1, $2) RETURNING id`
	var id int
	err := db.QueryRow(query, user.Name, user.Pass).Scan(&id) // hier returnt es uns die id von der Row (siehe query)
	if err != nil {
		log.Fatal(err)
	}
	return id
}
