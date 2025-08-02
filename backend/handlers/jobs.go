package handlers

import (
	"job-portal-backend/cache"
	"job-portal-backend/middleware"
	"job-portal-backend/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaginatedResponse struct {
	Jobs       []models.Job `json:"jobs"`
	Pagination struct {
		Page       int  `json:"page"`
		Limit      int  `json:"limit"`
		Total      int  `json:"total"`
		TotalPages int  `json:"total_pages"`
		HasNext    bool `json:"has_next"`
		HasPrev    bool `json:"has_prev"`
	} `json:"pagination"`
}

// GetJobs godoc
// @Summary Get all jobs with pagination and filters
// @Description Retrieve a list of jobs with optional filtering and pagination
// @Tags jobs
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param limit query int false "Number of items per page (default: 12, max: 50)"
// @Param location query string false "Filter by location"
// @Param salary_min query int false "Minimum salary filter"
// @Param salary_max query int false "Maximum salary filter"
// @Success 200 {object} PaginatedResponse
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /jobs [get]
func GetJobs(c *gin.Context) {
	// Parse query parameters untuk filter
	location := c.Query("location")
	salaryMinStr := c.Query("salary_min")
	salaryMaxStr := c.Query("salary_max")

	// Parse pagination parameters
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "12")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 12
	}

	// Set maximum limit
	if limit > 50 {
		limit = 50
	}

	filters := models.JobFilter{
		Location: location,
	}

	// Parse salary filters
	if salaryMinStr != "" {
		if salaryMin, err := strconv.Atoi(salaryMinStr); err == nil {
			filters.SalaryMin = salaryMin
		}
	}

	if salaryMaxStr != "" {
		if salaryMax, err := strconv.Atoi(salaryMaxStr); err == nil {
			filters.SalaryMax = salaryMax
		}
	}

	// Try to get from cache first (only for unfiltered requests)
	var response PaginatedResponse
	if location == "" && salaryMinStr == "" && salaryMaxStr == "" && page == 1 {
		if err := cache.GetCachedJobs(&response); err == nil {
			c.JSON(http.StatusOK, response)
			return
		}
	}

	// Get jobs with pagination
	jobs, total, err := models.GetJobsWithPagination(filters, page, limit)
	if err != nil {
		middleware.CustomError(c, http.StatusInternalServerError, "Database Error", "Failed to fetch jobs")
		return
	}

	// Calculate pagination info
	totalPages := (total + limit - 1) / limit
	hasNext := page < totalPages
	hasPrev := page > 1

	response = PaginatedResponse{
		Jobs: jobs,
	}
	response.Pagination.Page = page
	response.Pagination.Limit = limit
	response.Pagination.Total = total
	response.Pagination.TotalPages = totalPages
	response.Pagination.HasNext = hasNext
	response.Pagination.HasPrev = hasPrev

	// Cache the result if it's the first page without filters
	if location == "" && salaryMinStr == "" && salaryMaxStr == "" && page == 1 {
		cache.CacheJobs(response)
	}

	c.JSON(http.StatusOK, response)
}

// GetJobByID godoc
// @Summary Get a job by ID
// @Description Retrieve a specific job by its ID
// @Tags jobs
// @Accept json
// @Produce json
// @Param id path int true "Job ID"
// @Success 200 {object} models.Job
// @Failure 400 {object} map[string]interface{} "Invalid job ID"
// @Failure 404 {object} map[string]interface{} "Job not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /jobs/{id} [get]
func GetJobByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		middleware.CustomError(c, http.StatusBadRequest, "Invalid Input", "Invalid job ID")
		return
	}

	// Try to get from cache first
	var job *models.Job
	if err := cache.GetCachedJob(id, &job); err == nil {
		c.JSON(http.StatusOK, job)
		return
	}

	job, err = models.GetJobByID(id)
	if err != nil {
		middleware.CustomError(c, http.StatusInternalServerError, "Database Error", "Failed to fetch job")
		return
	}

	if job == nil {
		middleware.CustomError(c, http.StatusNotFound, "Not Found", "Job not found")
		return
	}

	// Cache the job
	cache.CacheJob(id, job)

	c.JSON(http.StatusOK, job)
}

// CreateJob godoc
// @Summary Create a new job
// @Description Create a new job posting
// @Tags jobs
// @Accept json
// @Produce json
// @Param job body models.Job true "Job object"
// @Success 201 {object} models.Job
// @Failure 400 {object} map[string]interface{} "Invalid request data"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /jobs [post]
func CreateJob(c *gin.Context) {
	var job models.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		middleware.CustomError(c, http.StatusBadRequest, "Invalid Input", "Invalid request data")
		return
	}

	// Validasi data
	if job.Position == "" || job.Company == "" || job.Location == "" {
		middleware.CustomError(c, http.StatusBadRequest, "Validation Error", "Position, company, and location are required")
		return
	}

	if job.SalaryMin <= 0 || job.SalaryMax <= 0 || job.SalaryMin > job.SalaryMax {
		middleware.CustomError(c, http.StatusBadRequest, "Validation Error", "Invalid salary range")
		return
	}

	err := models.CreateJob(&job)
	if err != nil {
		middleware.CustomError(c, http.StatusInternalServerError, "Database Error", "Failed to create job")
		return
	}

	// Invalidate cache
	cache.InvalidateJobsCache()

	c.JSON(http.StatusCreated, job)
}

// GetLocations godoc
// @Summary Get all available locations
// @Description Retrieve a list of all unique locations where jobs are available
// @Tags jobs
// @Accept json
// @Produce json
// @Success 200 {array} string
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /locations [get]
func GetLocations(c *gin.Context) {
	// Try to get from cache first
	var locations []string
	if err := cache.GetCachedLocations(&locations); err == nil {
		c.JSON(http.StatusOK, locations)
		return
	}

	locations, err := models.GetLocations()
	if err != nil {
		middleware.CustomError(c, http.StatusInternalServerError, "Database Error", "Failed to fetch locations")
		return
	}

	// Cache the locations
	cache.CacheLocations(locations)

	c.JSON(http.StatusOK, locations)
}

// SeedData godoc
// @Summary Seed database with sample data
// @Description Add sample jobs to the database
// @Tags jobs
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Seed completed"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /seed [post]
func SeedData(c *gin.Context) {
	// Check if database is empty
	count, err := models.GetJobCount()
	if err != nil {
		middleware.CustomError(c, http.StatusInternalServerError, "Database Error", "Failed to check job count")
		return
	}

	if count > 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Database already has data",
			"count":   count,
		})
		return
	}

	// Sample jobs data
	sampleJobs := []models.Job{
		{
			Position:  "Frontend Developer",
			Company:   "TechCorp Indonesia",
			Location:  "Jakarta",
			SalaryMin: 8000000,
			SalaryMax: 15000000,
		},
		{
			Position:  "Backend Developer",
			Company:   "Digital Solutions",
			Location:  "Bandung",
			SalaryMin: 10000000,
			SalaryMax: 18000000,
		},
		{
			Position:  "Full Stack Developer",
			Company:   "Innovation Labs",
			Location:  "Surabaya",
			SalaryMin: 12000000,
			SalaryMax: 20000000,
		},
		{
			Position:  "Mobile Developer",
			Company:   "App Studio",
			Location:  "Yogyakarta",
			SalaryMin: 9000000,
			SalaryMax: 16000000,
		},
		{
			Position:  "DevOps Engineer",
			Company:   "Cloud Systems",
			Location:  "Jakarta",
			SalaryMin: 11000000,
			SalaryMax: 19000000,
		},
		{
			Position:  "UI/UX Designer",
			Company:   "Creative Studio",
			Location:  "Bandung",
			SalaryMin: 7000000,
			SalaryMax: 13000000,
		},
		{
			Position:  "Data Scientist",
			Company:   "Analytics Pro",
			Location:  "Jakarta",
			SalaryMin: 12000000,
			SalaryMax: 22000000,
		},
		{
			Position:  "Product Manager",
			Company:   "Product Hub",
			Location:  "Surabaya",
			SalaryMin: 15000000,
			SalaryMax: 25000000,
		},
		{
			Position:  "QA Engineer",
			Company:   "Quality First",
			Location:  "Yogyakarta",
			SalaryMin: 6000000,
			SalaryMax: 11000000,
		},
		{
			Position:  "System Administrator",
			Company:   "IT Solutions",
			Location:  "Jakarta",
			SalaryMin: 8000000,
			SalaryMax: 14000000,
		},
	}

	// Insert sample jobs
	createdCount := 0
	for _, job := range sampleJobs {
		err := models.CreateJob(&job)
		if err != nil {
			log.Printf("Error creating job %s: %v", job.Position, err)
		} else {
			createdCount++
		}
	}

	// Invalidate cache
	cache.InvalidateJobsCache()

	c.JSON(http.StatusOK, gin.H{
		"message": "Sample data seeded successfully",
		"created": createdCount,
	})
}
