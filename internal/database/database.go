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
	log.Println(databaseTag, ": connecting to database..")
	connectionString := "postgres://postgres:mysecretpassword@localhost:5432/willweb-db?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(databaseTag, err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(databaseTag, err)
	}

	log.Println(databaseTag, ": database connected")

	log.Println(databaseTag, ": initializing tables")
	createTable(db)
	log.Println(databaseTag, ": tables initialized")

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
	if err := database.database.Close(); err != nil {
		log.Fatal(databaseTag, err)
	}

	log.Println(databaseTag, "database connection closed")
}

func createTable(db *sql.DB) {

	dropTopbar := `-- DROP TABLE IF EXISTS topbar_games;`
	dropUsers := `DROP TABLE IF EXISTS users CASCADE;`

	createUsers := `
	CREATE TABLE users (
		id BIGSERIAL PRIMARY KEY,
		first_name VARCHAR(100) NOT NULL,
		last_name VARCHAR(100) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT NOW()
	);`

	//createTopbar := `
	//CREATE TABLE test_table (
	//	id BIGSERIAL PRIMARY KEY,
	//	user_id BIGINT NOT NULL,
	//	game_name VARCHAR(100) NOT NULL,
	//	created_at TIMESTAMP DEFAULT NOW(),
	//	CONSTRAINT fk_user
	//		FOREIGN KEY (user_id)
	//		REFERENCES users(id)
	//		ON DELETE CASCADE
	//);`

	queries := []string{
		dropTopbar,
		dropUsers,
		createUsers,
		//createTopbar,
	}

	for _, q := range queries {
		if _, err := db.Exec(q); err != nil {
			log.Fatal(err)
		}
	}
}
