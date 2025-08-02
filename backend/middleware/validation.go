package middleware

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidationResponse represents validation error response
type ValidationResponse struct {
	Error   string            `json:"error"`
	Details []ValidationError `json:"details"`
}

// Validation rules
var (
	emailRegex    = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	phoneRegex    = regexp.MustCompile(`^[0-9+\-\s()]{10,15}$`)
	companyRegex  = regexp.MustCompile(`^[a-zA-Z0-9\s\-&.,()]{2,100}$`)
	positionRegex = regexp.MustCompile(`^[a-zA-Z0-9\s\-&.,()]{2,100}$`)
	locationRegex = regexp.MustCompile(`^[a-zA-Z\s\-.,()]{2,50}$`)
	nameRegex     = regexp.MustCompile(`^[a-zA-Z\s]{2,50}$`)
)

// ValidateJobInput validates job creation/update input
func ValidateJobInput() gin.HandlerFunc {
	return func(c *gin.Context) {
		var errors []ValidationError

		// Validate position
		position := c.PostForm("position")
		if position == "" {
			errors = append(errors, ValidationError{Field: "position", Message: "Position is required"})
		} else if !positionRegex.MatchString(position) {
			errors = append(errors, ValidationError{Field: "position", Message: "Position contains invalid characters"})
		}

		// Validate company
		company := c.PostForm("company")
		if company == "" {
			errors = append(errors, ValidationError{Field: "company", Message: "Company is required"})
		} else if !companyRegex.MatchString(company) {
			errors = append(errors, ValidationError{Field: "company", Message: "Company contains invalid characters"})
		}

		// Validate location
		location := c.PostForm("location")
		if location == "" {
			errors = append(errors, ValidationError{Field: "location", Message: "Location is required"})
		} else if !locationRegex.MatchString(location) {
			errors = append(errors, ValidationError{Field: "location", Message: "Location contains invalid characters"})
		}

		// Validate salary_min
		salaryMinStr := c.PostForm("salary_min")
		if salaryMinStr == "" {
			errors = append(errors, ValidationError{Field: "salary_min", Message: "Minimum salary is required"})
		} else {
			salaryMin, err := strconv.Atoi(salaryMinStr)
			if err != nil {
				errors = append(errors, ValidationError{Field: "salary_min", Message: "Minimum salary must be a valid number"})
			} else if salaryMin < 0 {
				errors = append(errors, ValidationError{Field: "salary_min", Message: "Minimum salary must be positive"})
			}
		}

		// Validate salary_max
		salaryMaxStr := c.PostForm("salary_max")
		if salaryMaxStr == "" {
			errors = append(errors, ValidationError{Field: "salary_max", Message: "Maximum salary is required"})
		} else {
			salaryMax, err := strconv.Atoi(salaryMaxStr)
			if err != nil {
				errors = append(errors, ValidationError{Field: "salary_max", Message: "Maximum salary must be a valid number"})
			} else if salaryMax < 0 {
				errors = append(errors, ValidationError{Field: "salary_max", Message: "Maximum salary must be positive"})
			}
		}

		// Validate salary range
		if len(errors) == 0 {
			salaryMin, _ := strconv.Atoi(salaryMinStr)
			salaryMax, _ := strconv.Atoi(salaryMaxStr)
			if salaryMin > salaryMax {
				errors = append(errors, ValidationError{Field: "salary_range", Message: "Minimum salary cannot be greater than maximum salary"})
			}
		}

		if len(errors) > 0 {
			c.JSON(http.StatusBadRequest, ValidationResponse{
				Error:   "Validation failed",
				Details: errors,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// ValidateApplicationInput validates application submission input
func ValidateApplicationInput() gin.HandlerFunc {
	return func(c *gin.Context) {
		var errors []ValidationError

		// Validate name
		name := c.PostForm("name")
		if name == "" {
			errors = append(errors, ValidationError{Field: "name", Message: "Name is required"})
		} else if !nameRegex.MatchString(name) {
			errors = append(errors, ValidationError{Field: "name", Message: "Name contains invalid characters"})
		}

		// Validate email
		email := c.PostForm("email")
		if email == "" {
			errors = append(errors, ValidationError{Field: "email", Message: "Email is required"})
		} else if !emailRegex.MatchString(email) {
			errors = append(errors, ValidationError{Field: "email", Message: "Invalid email format"})
		}

		// Validate job_id
		jobIDStr := c.PostForm("job_id")
		if jobIDStr == "" {
			errors = append(errors, ValidationError{Field: "job_id", Message: "Job ID is required"})
		} else {
			jobID, err := strconv.Atoi(jobIDStr)
			if err != nil {
				errors = append(errors, ValidationError{Field: "job_id", Message: "Job ID must be a valid number"})
			} else if jobID <= 0 {
				errors = append(errors, ValidationError{Field: "job_id", Message: "Job ID must be positive"})
			}
		}

		// Validate CV file
		file, err := c.FormFile("cv")
		if err != nil {
			errors = append(errors, ValidationError{Field: "cv", Message: "CV file is required"})
		} else {
			// Validate file type
			if !strings.HasSuffix(strings.ToLower(file.Filename), ".pdf") {
				errors = append(errors, ValidationError{Field: "cv", Message: "Only PDF files are allowed"})
			}

			// Validate file size (max 5MB)
			if file.Size > 5*1024*1024 {
				errors = append(errors, ValidationError{Field: "cv", Message: "File size must be less than 5MB"})
			}

			// Validate filename
			if len(file.Filename) > 255 {
				errors = append(errors, ValidationError{Field: "cv", Message: "Filename too long"})
			}
		}

		if len(errors) > 0 {
			c.JSON(http.StatusBadRequest, ValidationResponse{
				Error:   "Validation failed",
				Details: errors,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// ValidateQueryParams validates query parameters
func ValidateQueryParams() gin.HandlerFunc {
	return func(c *gin.Context) {
		var errors []ValidationError

		// Validate page parameter
		pageStr := c.Query("page")
		if pageStr != "" {
			page, err := strconv.Atoi(pageStr)
			if err != nil || page < 1 {
				errors = append(errors, ValidationError{Field: "page", Message: "Page must be a positive number"})
			}
		}

		// Validate limit parameter
		limitStr := c.Query("limit")
		if limitStr != "" {
			limit, err := strconv.Atoi(limitStr)
			if err != nil || limit < 1 || limit > 100 {
				errors = append(errors, ValidationError{Field: "limit", Message: "Limit must be between 1 and 100"})
			}
		}

		// Validate salary_min parameter
		salaryMinStr := c.Query("salary_min")
		if salaryMinStr != "" {
			salaryMin, err := strconv.Atoi(salaryMinStr)
			if err != nil || salaryMin < 0 {
				errors = append(errors, ValidationError{Field: "salary_min", Message: "Minimum salary must be a positive number"})
			}
		}

		// Validate salary_max parameter
		salaryMaxStr := c.Query("salary_max")
		if salaryMaxStr != "" {
			salaryMax, err := strconv.Atoi(salaryMaxStr)
			if err != nil || salaryMax < 0 {
				errors = append(errors, ValidationError{Field: "salary_max", Message: "Maximum salary must be a positive number"})
			}
		}

		// Validate location parameter
		location := c.Query("location")
		if location != "" && !locationRegex.MatchString(location) {
			errors = append(errors, ValidationError{Field: "location", Message: "Location contains invalid characters"})
		}

		if len(errors) > 0 {
			c.JSON(http.StatusBadRequest, ValidationResponse{
				Error:   "Invalid query parameters",
				Details: errors,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// SanitizeInput sanitizes input to prevent XSS and injection attacks
func SanitizeInput() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Sanitize form data
		for key, values := range c.Request.PostForm {
			for i, value := range values {
				// Remove potentially dangerous characters
				sanitized := strings.TrimSpace(value)
				sanitized = strings.ReplaceAll(sanitized, "<script>", "")
				sanitized = strings.ReplaceAll(sanitized, "</script>", "")
				sanitized = strings.ReplaceAll(sanitized, "javascript:", "")
				sanitized = strings.ReplaceAll(sanitized, "onload=", "")
				sanitized = strings.ReplaceAll(sanitized, "onerror=", "")
				c.Request.PostForm[key][i] = sanitized
			}
		}

		// Sanitize query parameters
		for key, values := range c.Request.URL.Query() {
			for i, value := range values {
				sanitized := strings.TrimSpace(value)
				sanitized = strings.ReplaceAll(sanitized, "<script>", "")
				sanitized = strings.ReplaceAll(sanitized, "</script>", "")
				sanitized = strings.ReplaceAll(sanitized, "javascript:", "")
				c.Request.URL.Query()[key][i] = sanitized
			}
		}

		c.Next()
	}
}
