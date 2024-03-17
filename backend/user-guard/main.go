package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

const (
	host     = "build-nexusdb-1"
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "mydb"
)

func main() {

	for {
		// Wait for the PostgreSQL container to be ready
		time.Sleep(5 * time.Second)

		// Build the connection string
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

		// Open the connection
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Fatalf("Failed to open PostgreSQL connection: %v", err)
		}
		defer db.Close()

		// Create the table
		_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL
		)
	`)
		if err != nil {
			fmt.Println("Failed to create table 'users':", err)
			continue
		}

		log.Println("Table 'users' created successfully")

	}

}
