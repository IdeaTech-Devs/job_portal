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
		// Frontend & Web Development
		{
			Position:  "Frontend Developer",
			Company:   "TechCorp Indonesia",
			Location:  "Jakarta",
			SalaryMin: 2000000,
			SalaryMax: 4000000,
		},
		{
			Position:  "React Developer",
			Company:   "WebTech Solutions",
			Location:  "Jakarta",
			SalaryMin: 3000000,
			SalaryMax: 5000000,
		},
		{
			Position:  "Vue.js Developer",
			Company:   "Digital Creative",
			Location:  "Bandung",
			SalaryMin: 2500000,
			SalaryMax: 4500000,
		},
		{
			Position:  "Angular Developer",
			Company:   "Enterprise Web",
			Location:  "Surabaya",
			SalaryMin: 3500000,
			SalaryMax: 5500000,
		},
		{
			Position:  "Svelte Developer",
			Company:   "Modern Web Studio",
			Location:  "Yogyakarta",
			SalaryMin: 2800000,
			SalaryMax: 4800000,
		},

		// Backend Development
		{
			Position:  "Backend Developer",
			Company:   "Digital Solutions",
			Location:  "Surabaya",
			SalaryMin: 3000000,
			SalaryMax: 5000000,
		},
		{
			Position:  "Python Developer",
			Company:   "CodeCraft Indonesia",
			Location:  "Surabaya",
			SalaryMin: 3500000,
			SalaryMax: 6000000,
		},
		{
			Position:  "Java Developer",
			Company:   "Enterprise Solutions",
			Location:  "Bandung",
			SalaryMin: 4000000,
			SalaryMax: 7000000,
		},
		{
			Position:  "Go Developer",
			Company:   "Microservices Inc",
			Location:  "Yogyakarta",
			SalaryMin: 4500000,
			SalaryMax: 7500000,
		},
		{
			Position:  "Node.js Developer",
			Company:   "JavaScript Pro",
			Location:  "Jakarta",
			SalaryMin: 3200000,
			SalaryMax: 5200000,
		},
		{
			Position:  "PHP Developer",
			Company:   "Web Solutions",
			Location:  "Bandung",
			SalaryMin: 2500000,
			SalaryMax: 4500000,
		},

		// Full Stack Development
		{
			Position:  "Full Stack Developer",
			Company:   "Innovation Labs",
			Location:  "Bandung",
			SalaryMin: 4000000,
			SalaryMax: 7000000,
		},
		{
			Position:  "MERN Stack Developer",
			Company:   "Modern Tech",
			Location:  "Jakarta",
			SalaryMin: 4500000,
			SalaryMax: 7500000,
		},
		{
			Position:  "Laravel Developer",
			Company:   "PHP Studio",
			Location:  "Surabaya",
			SalaryMin: 3000000,
			SalaryMax: 5500000,
		},

		// Mobile Development
		{
			Position:  "Mobile Developer",
			Company:   "AppWorks",
			Location:  "Bandung",
			SalaryMin: 3000000,
			SalaryMax: 5500000,
		},
		{
			Position:  "React Native Developer",
			Company:   "Cross Platform Apps",
			Location:  "Jakarta",
			SalaryMin: 3500000,
			SalaryMax: 6000000,
		},
		{
			Position:  "Flutter Developer",
			Company:   "Google Tech Indonesia",
			Location:  "Surabaya",
			SalaryMin: 3200000,
			SalaryMax: 5800000,
		},
		{
			Position:  "iOS Developer",
			Company:   "Apple Developer Studio",
			Location:  "Jakarta",
			SalaryMin: 4000000,
			SalaryMax: 7000000,
		},
		{
			Position:  "Android Developer",
			Company:   "Android Studio Pro",
			Location:  "Bandung",
			SalaryMin: 3800000,
			SalaryMax: 6500000,
		},

		// DevOps & Infrastructure
		{
			Position:  "DevOps Engineer",
			Company:   "CloudTech",
			Location:  "Yogyakarta",
			SalaryMin: 5000000,
			SalaryMax: 8000000,
		},
		{
			Position:  "Site Reliability Engineer",
			Company:   "Reliability First",
			Location:  "Jakarta",
			SalaryMin: 5500000,
			SalaryMax: 8500000,
		},
		{
			Position:  "Cloud Engineer",
			Company:   "AWS Solutions",
			Location:  "Surabaya",
			SalaryMin: 4500000,
			SalaryMax: 7500000,
		},
		{
			Position:  "Kubernetes Engineer",
			Company:   "Container Tech",
			Location:  "Bandung",
			SalaryMin: 4800000,
			SalaryMax: 7800000,
		},

		// Data & Analytics
		{
			Position:  "Data Scientist",
			Company:   "Analytics Pro",
			Location:  "Surabaya",
			SalaryMin: 4000000,
			SalaryMax: 6000000,
		},
		{
			Position:  "Data Engineer",
			Company:   "Big Data Solutions",
			Location:  "Jakarta",
			SalaryMin: 4500000,
			SalaryMax: 7000000,
		},
		{
			Position:  "Machine Learning Engineer",
			Company:   "AI Innovation",
			Location:  "Bandung",
			SalaryMin: 5000000,
			SalaryMax: 8000000,
		},
		{
			Position:  "Business Intelligence Analyst",
			Company:   "BI Solutions",
			Location:  "Yogyakarta",
			SalaryMin: 3500000,
			SalaryMax: 5500000,
		},

		// Design & UX
		{
			Position:  "UI/UX Designer",
			Company:   "Creative Studio",
			Location:  "Jakarta",
			SalaryMin: 2500000,
			SalaryMax: 4500000,
		},
		{
			Position:  "Product Designer",
			Company:   "Design Thinking",
			Location:  "Bandung",
			SalaryMin: 3000000,
			SalaryMax: 5000000,
		},
		{
			Position:  "Graphic Designer",
			Company:   "Visual Arts Studio",
			Location:  "Surabaya",
			SalaryMin: 2000000,
			SalaryMax: 3500000,
		},
		{
			Position:  "UX Researcher",
			Company:   "User Experience Lab",
			Location:  "Jakarta",
			SalaryMin: 2800000,
			SalaryMax: 4800000,
		},

		// Product & Project Management
		{
			Position:  "Product Manager",
			Company:   "StartupHub",
			Location:  "Bali",
			SalaryMin: 6000000,
			SalaryMax: 10000000,
		},
		{
			Position:  "Project Manager",
			Company:   "Project Solutions",
			Location:  "Jakarta",
			SalaryMin: 5000000,
			SalaryMax: 8000000,
		},
		{
			Position:  "Scrum Master",
			Company:   "Agile Solutions",
			Location:  "Bandung",
			SalaryMin: 4500000,
			SalaryMax: 7000000,
		},
		{
			Position:  "Product Owner",
			Company:   "Product First",
			Location:  "Surabaya",
			SalaryMin: 5500000,
			SalaryMax: 8500000,
		},

		// Quality Assurance
		{
			Position:  "QA Engineer",
			Company:   "Quality First",
			Location:  "Yogyakarta",
			SalaryMin: 2000000,
			SalaryMax: 3500000,
		},
		{
			Position:  "Test Automation Engineer",
			Company:   "Automation Pro",
			Location:  "Jakarta",
			SalaryMin: 3500000,
			SalaryMax: 5500000,
		},
		{
			Position:  "Performance Tester",
			Company:   "Performance Lab",
			Location:  "Bandung",
			SalaryMin: 3000000,
			SalaryMax: 5000000,
		},

		// System & Network
		{
			Position:  "System Administrator",
			Company:   "IT Solutions",
			Location:  "Bali",
			SalaryMin: 3500000,
			SalaryMax: 5500000,
		},
		{
			Position:  "Network Engineer",
			Company:   "Network Solutions",
			Location:  "Jakarta",
			SalaryMin: 4000000,
			SalaryMax: 6000000,
		},
		{
			Position:  "Security Engineer",
			Company:   "Cyber Security Pro",
			Location:  "Surabaya",
			SalaryMin: 4500000,
			SalaryMax: 7000000,
		},

		// Emerging Technologies
		{
			Position:  "Blockchain Developer",
			Company:   "Crypto Labs",
			Location:  "Bali",
			SalaryMin: 5000000,
			SalaryMax: 8000000,
		},
		{
			Position:  "IoT Developer",
			Company:   "Internet of Things",
			Location:  "Bandung",
			SalaryMin: 4000000,
			SalaryMax: 6500000,
		},
		{
			Position:  "AR/VR Developer",
			Company:   "Virtual Reality Studio",
			Location:  "Jakarta",
			SalaryMin: 4500000,
			SalaryMax: 7500000,
		},

		// Entry Level Positions
		{
			Position:  "Junior Developer",
			Company:   "Junior Tech",
			Location:  "Jakarta",
			SalaryMin: 1500000,
			SalaryMax: 2500000,
		},
		{
			Position:  "Graduate Developer",
			Company:   "Fresh Graduate Hub",
			Location:  "Bandung",
			SalaryMin: 1200000,
			SalaryMax: 2000000,
		},
		{
			Position:  "Intern Developer",
			Company:   "Internship Program",
			Location:  "Surabaya",
			SalaryMin: 800000,
			SalaryMax: 1500000,
		},

		// Senior & Lead Positions
		{
			Position:  "Senior Developer",
			Company:   "Senior Tech Solutions",
			Location:  "Jakarta",
			SalaryMin: 8000000,
			SalaryMax: 12000000,
		},
		{
			Position:  "Tech Lead",
			Company:   "Leadership Tech",
			Location:  "Bandung",
			SalaryMin: 10000000,
			SalaryMax: 15000000,
		},
		{
			Position:  "Engineering Manager",
			Company:   "Management Solutions",
			Location:  "Surabaya",
			SalaryMin: 12000000,
			SalaryMax: 18000000,
		},

		// Specialized Roles
		{
			Position:  "Game Developer",
			Company:   "Game Studio Indonesia",
			Location:  "Yogyakarta",
			SalaryMin: 3500000,
			SalaryMax: 6000000,
		},
		{
			Position:  "Embedded Systems Engineer",
			Company:   "Hardware Solutions",
			Location:  "Bandung",
			SalaryMin: 4000000,
			SalaryMax: 6500000,
		},
		{
			Position:  "Database Administrator",
			Company:   "Database Solutions",
			Location:  "Jakarta",
			SalaryMin: 4500000,
			SalaryMax: 7000000,
		},
		{
			Position:  "API Developer",
			Company:   "API Solutions",
			Location:  "Surabaya",
			SalaryMin: 3500000,
			SalaryMax: 5500000,
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
