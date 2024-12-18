package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	dsn := "host=localhost port=5432 user=wahyu password=postgres dbname=golangcrud sslmode=disable TimeZone=Asia/Jakarta"
	db, err := sql.Open("postgres", dsn)
	
	if err != nil {
		log.Fatalf("Failed to connect to the database %v", err)
	}

	log.Println("Database connected")
	DB = db
}