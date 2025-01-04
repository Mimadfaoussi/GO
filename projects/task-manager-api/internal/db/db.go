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

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

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

