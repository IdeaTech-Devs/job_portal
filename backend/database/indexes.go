package database

import (
	"log"
)

// CreateIndexes creates database indexes for better performance
func CreateIndexes() {
	// Index for jobs table
	createJobIndexes := []string{
		// Index on location for filtering
		"CREATE INDEX IF NOT EXISTS idx_jobs_location ON jobs(location)",

		// Index on salary range for filtering
		"CREATE INDEX IF NOT EXISTS idx_jobs_salary_range ON jobs(salary_min, salary_max)",

		// Index on created_at for sorting
		"CREATE INDEX IF NOT EXISTS idx_jobs_created_at ON jobs(created_at DESC)",

		// Composite index for location and salary filtering
		"CREATE INDEX IF NOT EXISTS idx_jobs_location_salary ON jobs(location, salary_min, salary_max)",

		// Index on company for searching
		"CREATE INDEX IF NOT EXISTS idx_jobs_company ON jobs(company)",

		// Index on position for searching
		"CREATE INDEX IF NOT EXISTS idx_jobs_position ON jobs(position)",
	}

	// Index for applications table
	createApplicationIndexes := []string{
		// Index on job_id for foreign key lookups
		"CREATE INDEX IF NOT EXISTS idx_applications_job_id ON applications(job_id)",

		// Index on applied_at for sorting
		"CREATE INDEX IF NOT EXISTS idx_applications_applied_at ON applications(applied_at DESC)",

		// Index on email for searching
		"CREATE INDEX IF NOT EXISTS idx_applications_email ON applications(email)",

		// Composite index for job_id and applied_at
		"CREATE INDEX IF NOT EXISTS idx_applications_job_applied ON applications(job_id, applied_at DESC)",
	}

	// Create job indexes
	for _, query := range createJobIndexes {
		_, err := DB.Exec(query)
		if err != nil {
			log.Printf("Error creating job index: %v", err)
		} else {
			log.Printf("Job index created successfully")
		}
	}

	// Create application indexes
	for _, query := range createApplicationIndexes {
		_, err := DB.Exec(query)
		if err != nil {
			log.Printf("Error creating application index: %v", err)
		} else {
			log.Printf("Application index created successfully")
		}
	}

	log.Println("Database indexes created successfully")
}

// AnalyzeTables runs ANALYZE on tables for query optimization
func AnalyzeTables() {
	tables := []string{"jobs", "applications"}

	for _, table := range tables {
		query := "ANALYZE " + table
		_, err := DB.Exec(query)
		if err != nil {
			log.Printf("Error analyzing table %s: %v", table, err)
		} else {
			log.Printf("Table %s analyzed successfully", table)
		}
	}
}

// VacuumTables runs VACUUM on tables to reclaim storage
func VacuumTables() {
	tables := []string{"jobs", "applications"}

	for _, table := range tables {
		query := "VACUUM " + table
		_, err := DB.Exec(query)
		if err != nil {
			log.Printf("Error vacuuming table %s: %v", table, err)
		} else {
			log.Printf("Table %s vacuumed successfully", table)
		}
	}
}

// GetTableStats returns table statistics for monitoring
func GetTableStats() map[string]interface{} {
	stats := make(map[string]interface{})

	// Get jobs count
	var jobsCount int
	err := DB.QueryRow("SELECT COUNT(*) FROM jobs").Scan(&jobsCount)
	if err != nil {
		log.Printf("Error getting jobs count: %v", err)
	} else {
		stats["jobs_count"] = jobsCount
	}

	// Get applications count
	var applicationsCount int
	err = DB.QueryRow("SELECT COUNT(*) FROM applications").Scan(&applicationsCount)
	if err != nil {
		log.Printf("Error getting applications count: %v", err)
	} else {
		stats["applications_count"] = applicationsCount
	}

	// Get table sizes
	var jobsSize, applicationsSize int64
	err = DB.QueryRow("SELECT pg_total_relation_size('jobs')").Scan(&jobsSize)
	if err == nil {
		stats["jobs_table_size_bytes"] = jobsSize
	}

	err = DB.QueryRow("SELECT pg_total_relation_size('applications')").Scan(&applicationsSize)
	if err == nil {
		stats["applications_table_size_bytes"] = applicationsSize
	}

	return stats
}
