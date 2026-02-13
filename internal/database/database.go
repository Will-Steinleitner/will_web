package database

import (
	"database/sql"
	"log"
)

type Database struct {
	database *sql.DB
}

const databaseTag = "Database"

func NewDatabase() *Database {
	connectionString := "postgres://postgres:mysecretpassword@localhost:5432/willweb-db?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)

	log.Println(databaseTag, db.Stats())
	if err != nil {
		log.Fatal(databaseTag, err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(databaseTag, err)
	}

	createTable(db)

	return &Database{
		db,
	}
}

func (database *Database) GetDatabase() *sql.DB {
	return database.database
}

// Close The deferred function will be executed when functions returns.
// However, main() will not return while http.ListenAndServe is running,because it blocks and keeps the server alive.
// Therefore, the deferred function will only run if the server stops or an error occurs.
func (database *Database) Close() {
	err := database.database.Close()
	//we need our own Close func, because
	log.Println(databaseTag, "Database connection closed")
	if err != nil {
		log.Fatal(databaseTag, err)
	}
}

func createTable(db *sql.DB) {
	query := `
	DROP TABLE IF EXISTS users;

	CREATE TABLE users (
		id BIGSERIAL PRIMARY KEY,
		first_name VARCHAR(100) NOT NULL,
		last_name VARCHAR(100) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT NOW()
	);
`
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}
