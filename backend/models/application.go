package models

import (
	"database/sql"
	"job-portal-backend/database"
	"time"
)

// Application represents a job application
// @Description Job application information
type Application struct {
	ID         int       `json:"id" example:"1"`
	JobID      int       `json:"job_id" example:"1"`
	Name       string    `json:"name" example:"John Doe"`
	Email      string    `json:"email" example:"john.doe@example.com"`
	CVFilename string    `json:"cv_filename" example:"cv_1234567890.pdf"`
	AppliedAt  time.Time `json:"applied_at" example:"2025-01-15T10:30:00Z"`
	Job        *Job      `json:"job,omitempty"`
}

func CreateApplication(app *Application) error {
	query := `INSERT INTO applications (job_id, name, email, cv_filename) 
			  VALUES ($1, $2, $3, $4) RETURNING id, applied_at`

	return database.DB.QueryRow(query, app.JobID, app.Name, app.Email, app.CVFilename).
		Scan(&app.ID, &app.AppliedAt)
}

func GetApplications() ([]Application, error) {
	query := `
		SELECT a.id, a.job_id, a.name, a.email, a.cv_filename, a.applied_at,
			   j.position, j.company, j.location, j.salary_min, j.salary_max, j.created_at
		FROM applications a
		JOIN jobs j ON a.job_id = j.id
		ORDER BY a.applied_at DESC
	`

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var applications []Application
	for rows.Next() {
		var app Application
		var job Job
		err := rows.Scan(
			&app.ID, &app.JobID, &app.Name, &app.Email, &app.CVFilename, &app.AppliedAt,
			&job.Position, &job.Company, &job.Location, &job.SalaryMin, &job.SalaryMax, &job.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		app.Job = &job
		applications = append(applications, app)
	}

	return applications, nil
}

func GetApplicationByID(id int) (*Application, error) {
	query := `
		SELECT a.id, a.job_id, a.name, a.email, a.cv_filename, a.applied_at,
			   j.position, j.company, j.location, j.salary_min, j.salary_max, j.created_at
		FROM applications a
		JOIN jobs j ON a.job_id = j.id
		WHERE a.id = $1
	`

	var app Application
	var job Job
	err := database.DB.QueryRow(query, id).Scan(
		&app.ID, &app.JobID, &app.Name, &app.Email, &app.CVFilename, &app.AppliedAt,
		&job.Position, &job.Company, &job.Location, &job.SalaryMin, &job.SalaryMax, &job.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	app.Job = &job
	return &app, nil
}
