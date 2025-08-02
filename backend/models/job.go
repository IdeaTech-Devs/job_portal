package models

import (
	"database/sql"
	"job-portal-backend/database"
	"time"
)

// Job represents a job posting
// @Description Job posting information
type Job struct {
	ID        int       `json:"id" example:"1"`
	Position  string    `json:"position" example:"Frontend Developer"`
	Company   string    `json:"company" example:"TechCorp Indonesia"`
	Location  string    `json:"location" example:"Jakarta"`
	SalaryMin int       `json:"salary_min" example:"3000000"`
	SalaryMax int       `json:"salary_max" example:"5000000"`
	CreatedAt time.Time `json:"created_at" example:"2025-01-15T10:30:00Z"`
}

// JobFilter represents filters for job search
// @Description Job search filters
type JobFilter struct {
	Location  string `json:"location" example:"Jakarta"`
	SalaryMin int    `json:"salary_min" example:"2000000"`
	SalaryMax int    `json:"salary_max" example:"8000000"`
	Search    string `json:"search" example:"developer"`
}

func GetAllJobs(filters JobFilter) ([]Job, error) {
	query := "SELECT id, position, company, location, salary_min, salary_max, created_at FROM jobs WHERE 1=1"
	args := []interface{}{}
	argIndex := 1

	if filters.Location != "" {
		query += " AND location = $" + string(rune(argIndex+'0'))
		args = append(args, filters.Location)
		argIndex++
	}

	if filters.Search != "" {
		query += " AND (position ILIKE $" + string(rune(argIndex+'0')) + " OR company ILIKE $" + string(rune(argIndex+1+'0')) + " OR location ILIKE $" + string(rune(argIndex+2+'0')) + ")"
		searchTerm := "%" + filters.Search + "%"
		args = append(args, searchTerm, searchTerm, searchTerm)
		argIndex += 3
	}

	if filters.SalaryMin > 0 {
		query += " AND salary_max >= $" + string(rune(argIndex+'0'))
		args = append(args, filters.SalaryMin)
		argIndex++
	}

	if filters.SalaryMax > 0 {
		query += " AND salary_min <= $" + string(rune(argIndex+'0'))
		args = append(args, filters.SalaryMax)
		argIndex++
	}

	query += " ORDER BY created_at DESC"

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []Job
	for rows.Next() {
		var job Job
		err := rows.Scan(&job.ID, &job.Position, &job.Company, &job.Location, &job.SalaryMin, &job.SalaryMax, &job.CreatedAt)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}

	return jobs, nil
}

func GetJobsWithPagination(filters JobFilter, page, limit int) ([]Job, int, error) {
	// Build base query for counting total
	countQuery := "SELECT COUNT(*) FROM jobs WHERE 1=1"
	args := []interface{}{}
	argIndex := 1

	if filters.Location != "" {
		countQuery += " AND location = $" + string(rune(argIndex+'0'))
		args = append(args, filters.Location)
		argIndex++
	}

	if filters.Search != "" {
		countQuery += " AND (position ILIKE $" + string(rune(argIndex+'0')) + " OR company ILIKE $" + string(rune(argIndex+1+'0')) + " OR location ILIKE $" + string(rune(argIndex+2+'0')) + ")"
		searchTerm := "%" + filters.Search + "%"
		args = append(args, searchTerm, searchTerm, searchTerm)
		argIndex += 3
	}

	if filters.SalaryMin > 0 {
		countQuery += " AND salary_max >= $" + string(rune(argIndex+'0'))
		args = append(args, filters.SalaryMin)
		argIndex++
	}

	if filters.SalaryMax > 0 {
		countQuery += " AND salary_min <= $" + string(rune(argIndex+'0'))
		args = append(args, filters.SalaryMax)
		argIndex++
	}

	// Get total count
	var total int
	err := database.DB.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Build query for getting jobs with pagination
	query := "SELECT id, position, company, location, salary_min, salary_max, created_at FROM jobs WHERE 1=1"
	queryArgs := []interface{}{}
	queryArgIndex := 1

	if filters.Location != "" {
		query += " AND location = $" + string(rune(queryArgIndex+'0'))
		queryArgs = append(queryArgs, filters.Location)
		queryArgIndex++
	}

	if filters.Search != "" {
		query += " AND (position ILIKE $" + string(rune(queryArgIndex+'0')) + " OR company ILIKE $" + string(rune(queryArgIndex+1+'0')) + " OR location ILIKE $" + string(rune(queryArgIndex+2+'0')) + ")"
		searchTerm := "%" + filters.Search + "%"
		queryArgs = append(queryArgs, searchTerm, searchTerm, searchTerm)
		queryArgIndex += 3
	}

	if filters.SalaryMin > 0 {
		query += " AND salary_max >= $" + string(rune(queryArgIndex+'0'))
		queryArgs = append(queryArgs, filters.SalaryMin)
		queryArgIndex++
	}

	if filters.SalaryMax > 0 {
		query += " AND salary_min <= $" + string(rune(queryArgIndex+'0'))
		queryArgs = append(queryArgs, filters.SalaryMax)
		queryArgIndex++
	}

	query += " ORDER BY created_at DESC LIMIT $" + string(rune(queryArgIndex+'0')) + " OFFSET $" + string(rune(queryArgIndex+1+'0'))
	queryArgs = append(queryArgs, limit, (page-1)*limit)

	rows, err := database.DB.Query(query, queryArgs...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var jobs []Job
	for rows.Next() {
		var job Job
		err := rows.Scan(&job.ID, &job.Position, &job.Company, &job.Location, &job.SalaryMin, &job.SalaryMax, &job.CreatedAt)
		if err != nil {
			return nil, 0, err
		}
		jobs = append(jobs, job)
	}

	return jobs, total, nil
}

func GetJobByID(id int) (*Job, error) {
	var job Job
	err := database.DB.QueryRow("SELECT id, position, company, location, salary_min, salary_max, created_at FROM jobs WHERE id = $1", id).
		Scan(&job.ID, &job.Position, &job.Company, &job.Location, &job.SalaryMin, &job.SalaryMax, &job.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &job, nil
}

func CreateJob(job *Job) error {
	query := `INSERT INTO jobs (position, company, location, salary_min, salary_max) 
			  VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`

	return database.DB.QueryRow(query, job.Position, job.Company, job.Location, job.SalaryMin, job.SalaryMax).
		Scan(&job.ID, &job.CreatedAt)
}

func GetLocations() ([]string, error) {
	rows, err := database.DB.Query("SELECT DISTINCT location FROM jobs ORDER BY location")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locations []string
	for rows.Next() {
		var location string
		err := rows.Scan(&location)
		if err != nil {
			return nil, err
		}
		locations = append(locations, location)
	}

	return locations, nil
}
