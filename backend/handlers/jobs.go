package handlers

import (
	"job-portal-backend/models"
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

	// Get jobs with pagination
	jobs, total, err := models.GetJobsWithPagination(filters, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch jobs"})
		return
	}

	// Calculate pagination info
	totalPages := (total + limit - 1) / limit
	hasNext := page < totalPages
	hasPrev := page > 1

	response := PaginatedResponse{
		Jobs: jobs,
	}
	response.Pagination.Page = page
	response.Pagination.Limit = limit
	response.Pagination.Total = total
	response.Pagination.TotalPages = totalPages
	response.Pagination.HasNext = hasNext
	response.Pagination.HasPrev = hasPrev

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}

	job, err := models.GetJobByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch job"})
		return
	}

	if job == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
		return
	}

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Validasi data
	if job.Position == "" || job.Company == "" || job.Location == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Position, company, and location are required"})
		return
	}

	if job.SalaryMin <= 0 || job.SalaryMax <= 0 || job.SalaryMin > job.SalaryMax {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid salary range"})
		return
	}

	err := models.CreateJob(&job)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create job"})
		return
	}

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
	locations, err := models.GetLocations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch locations"})
		return
	}

	c.JSON(http.StatusOK, locations)
}
