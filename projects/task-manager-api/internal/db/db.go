package db


import (

	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

//database connection logic

func ConnectDB() *sql.DB {

	connStr := "host=localhost port=5432 user=taskuser password=taskpassword dbname=taskdb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Database is not reachable: %v", err)
	}

	fmt.Println("Connected to the database successfully!")
	return db
}

