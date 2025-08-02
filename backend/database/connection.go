package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	var err error

	// Check if DATABASE_URL is set (Railway deployment)
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL != "" {
		// Use DATABASE_URL for Railway deployment
		DB, err = sql.Open("postgres", databaseURL)
		if err != nil {
			log.Fatal("Error opening database with DATABASE_URL:", err)
		}
	} else {
		// Use .env file variables for local development
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")

		// Validate required environment variables
		if host == "" || port == "" || user == "" || password == "" || dbname == "" {
			log.Fatal("Missing required database environment variables in .env file")
		}

		// String koneksi PostgreSQL
		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)

		// Buka koneksi database
		DB, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Fatal("Error opening database:", err)
		}
	}

	// Test koneksi
	err = DB.Ping()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	log.Println("Successfully connected to database")

	// Buat tabel jika belum ada
	createTables()
}

func createTables() {
	// Tabel jobs
	createJobsTable := `
	CREATE TABLE IF NOT EXISTS jobs (
		id SERIAL PRIMARY KEY,
		position VARCHAR(255) NOT NULL,
		company VARCHAR(255) NOT NULL,
		location VARCHAR(255) NOT NULL,
		salary_min INTEGER NOT NULL,
		salary_max INTEGER NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	// Tabel applications
	createApplicationsTable := `
	CREATE TABLE IF NOT EXISTS applications (
		id SERIAL PRIMARY KEY,
		job_id INTEGER REFERENCES jobs(id),
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		cv_filename VARCHAR(255) NOT NULL,
		applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(createJobsTable)
	if err != nil {
		log.Fatal("Error creating jobs table:", err)
	}

	_, err = DB.Exec(createApplicationsTable)
	if err != nil {
		log.Fatal("Error creating applications table:", err)
	}

	log.Println("Tables created successfully")
}
