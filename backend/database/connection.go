package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	// Konfigurasi database PostgreSQL
	host := "localhost"
	port := 5050
	user := "postgres"
	password := "postgres"
	dbname := "job_portal"

	// String koneksi PostgreSQL
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Buka koneksi database
	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error opening database:", err)
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
