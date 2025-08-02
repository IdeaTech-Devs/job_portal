package handlers

import (
	"fmt"
	"job-portal-backend/models"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateApplication godoc
// @Summary Submit a job application
// @Description Submit a job application with CV file upload
// @Tags applications
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "Applicant name"
// @Param email formData string true "Applicant email"
// @Param job_id formData int true "Job ID"
// @Param cv formData file true "CV file (PDF only, max 5MB)"
// @Success 201 {object} map[string]interface{} "Application submitted successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request data"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /applications [post]
func CreateApplication(c *gin.Context) {
	// Parse form data
	name := c.PostForm("name")
	email := c.PostForm("email")
	jobIDStr := c.PostForm("job_id")

	// Validasi input
	if name == "" || email == "" || jobIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name, email, and job_id are required"})
		return
	}

	// Parse job ID
	jobID, err := strconv.Atoi(jobIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}

	// Validasi email format sederhana
	if !strings.Contains(email, "@") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	// Handle file upload
	file, err := c.FormFile("cv")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CV file is required"})
		return
	}

	// Validasi file type
	if !isValidPDF(file) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only PDF files are allowed"})
		return
	}

	// Validasi file size (max 5MB)
	if file.Size > 5*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size must be less than 5MB"})
		return
	}

	// Generate unique filename
	filename := generateUniqueFilename(file.Filename)

	// Save file - Use /tmp for Railway deployment
	uploadPath := fmt.Sprintf("/tmp/%s", filename)
	if err := c.SaveUploadedFile(file, uploadPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Create application
	application := &models.Application{
		JobID:      jobID,
		Name:       name,
		Email:      email,
		CVFilename: filename,
	}

	err = models.CreateApplication(application)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create application"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":     "Application submitted successfully",
		"application": application,
	})
}

// GetApplications godoc
// @Summary Get all applications
// @Description Retrieve a list of all job applications
// @Tags applications
// @Accept json
// @Produce json
// @Success 200 {array} models.Application
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /applications [get]
func GetApplications(c *gin.Context) {
	applications, err := models.GetApplications()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch applications"})
		return
	}

	c.JSON(http.StatusOK, applications)
}

// GetApplicationByID godoc
// @Summary Get application by ID
// @Description Retrieve a specific application by its ID
// @Tags applications
// @Accept json
// @Produce json
// @Param id path int true "Application ID"
// @Success 200 {object} models.Application
// @Failure 400 {object} map[string]interface{} "Invalid application ID"
// @Failure 404 {object} map[string]interface{} "Application not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /applications/{id} [get]
func GetApplicationByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid application ID"})
		return
	}

	application, err := models.GetApplicationByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch application"})
		return
	}

	if application == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
		return
	}

	c.JSON(http.StatusOK, application)
}

// Helper functions
func isValidPDF(file *multipart.FileHeader) bool {
	ext := strings.ToLower(filepath.Ext(file.Filename))
	return ext == ".pdf"
}

func generateUniqueFilename(originalName string) string {
	ext := filepath.Ext(originalName)
	timestamp := time.Now().Unix()
	return fmt.Sprintf("cv_%d%s", timestamp, ext)
}
