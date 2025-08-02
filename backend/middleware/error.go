package middleware

import (
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
)

// ErrorResponse represents a structured error response
type ErrorResponse struct {
	Error      string    `json:"error"`
	Message    string    `json:"message"`
	Timestamp  time.Time `json:"timestamp"`
	RequestID  string    `json:"request_id,omitempty"`
	Path       string    `json:"path,omitempty"`
	Method     string    `json:"method,omitempty"`
	StatusCode int       `json:"status_code"`
}

// LogEntry represents a structured log entry
type LogEntry struct {
	Timestamp  time.Time     `json:"timestamp"`
	Level      string        `json:"level"`
	Message    string        `json:"message"`
	RequestID  string        `json:"request_id,omitempty"`
	ClientIP   string        `json:"client_ip"`
	Method     string        `json:"method"`
	Path       string        `json:"path"`
	StatusCode int           `json:"status_code"`
	Duration   time.Duration `json:"duration"`
	UserAgent  string        `json:"user_agent"`
	Error      string        `json:"error,omitempty"`
	Stack      string        `json:"stack,omitempty"`
}

// ErrorHandler middleware handles panics and errors
func ErrorHandler() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		// Log the panic
		log.Printf("PANIC: %v", recovered)
		log.Printf("Stack: %s", debug.Stack())

		// Create error response
		errorResponse := ErrorResponse{
			Error:      "Internal Server Error",
			Message:    "An unexpected error occurred",
			Timestamp:  time.Now(),
			RequestID:  c.GetString("request_id"),
			Path:       c.Request.URL.Path,
			Method:     c.Request.Method,
			StatusCode: http.StatusInternalServerError,
		}

		// Log structured error
		logEntry := LogEntry{
			Timestamp:  time.Now(),
			Level:      "ERROR",
			Message:    "Panic recovered",
			RequestID:  c.GetString("request_id"),
			ClientIP:   c.ClientIP(),
			Method:     c.Request.Method,
			Path:       c.Request.URL.Path,
			StatusCode: http.StatusInternalServerError,
			UserAgent:  c.Request.UserAgent(),
			Error:      errorResponse.Error,
			Stack:      string(debug.Stack()),
		}

		log.Printf("ERROR_LOG: %+v", logEntry)

		c.JSON(http.StatusInternalServerError, errorResponse)
	})
}

// RequestLogger middleware logs all requests
func RequestLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// Create structured log entry
		logEntry := LogEntry{
			Timestamp:  param.TimeStamp,
			Level:      "INFO",
			Message:    "HTTP Request",
			RequestID:  param.Keys["request_id"].(string),
			ClientIP:   param.ClientIP,
			Method:     param.Method,
			Path:       param.Path,
			StatusCode: param.StatusCode,
			Duration:   param.Latency,
			UserAgent:  param.Request.UserAgent(),
		}

		// Log based on status code
		if param.StatusCode >= 400 {
			logEntry.Level = "ERROR"
			logEntry.Message = "HTTP Error"
		}

		log.Printf("REQUEST_LOG: %+v", logEntry)
		return ""
	})
}

// RequestID middleware adds a unique request ID
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = generateRequestID()
		}
		c.Set("request_id", requestID)
		c.Header("X-Request-ID", requestID)
		c.Next()
	}
}

// generateRequestID generates a unique request ID
func generateRequestID() string {
	return time.Now().Format("20060102150405") + "-" + randomString(8)
}

// randomString generates a random string
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(b)
}

// CustomError creates a custom error response
func CustomError(c *gin.Context, statusCode int, errorType, message string) {
	errorResponse := ErrorResponse{
		Error:      errorType,
		Message:    message,
		Timestamp:  time.Now(),
		RequestID:  c.GetString("request_id"),
		Path:       c.Request.URL.Path,
		Method:     c.Request.Method,
		StatusCode: statusCode,
	}

	// Log the error
	logEntry := LogEntry{
		Timestamp:  time.Now(),
		Level:      "ERROR",
		Message:    "Custom Error",
		RequestID:  c.GetString("request_id"),
		ClientIP:   c.ClientIP(),
		Method:     c.Request.Method,
		Path:       c.Request.URL.Path,
		StatusCode: statusCode,
		UserAgent:  c.Request.UserAgent(),
		Error:      errorType,
	}

	log.Printf("ERROR_LOG: %+v", logEntry)

	c.JSON(statusCode, errorResponse)
}

// ValidationErrorResponse creates a validation error response
func ValidationErrorResponse(c *gin.Context, errors []ValidationError) {
	errorResponse := ErrorResponse{
		Error:      "Validation Error",
		Message:    "Input validation failed",
		Timestamp:  time.Now(),
		RequestID:  c.GetString("request_id"),
		Path:       c.Request.URL.Path,
		Method:     c.Request.Method,
		StatusCode: http.StatusBadRequest,
	}

	// Log the validation error
	logEntry := LogEntry{
		Timestamp:  time.Now(),
		Level:      "WARN",
		Message:    "Validation Error",
		RequestID:  c.GetString("request_id"),
		ClientIP:   c.ClientIP(),
		Method:     c.Request.Method,
		Path:       c.Request.URL.Path,
		StatusCode: http.StatusBadRequest,
		UserAgent:  c.Request.UserAgent(),
		Error:      "Validation failed",
	}

	log.Printf("VALIDATION_LOG: %+v", logEntry)

	c.JSON(http.StatusBadRequest, gin.H{
		"error":      errorResponse.Error,
		"message":    errorResponse.Message,
		"details":    errors,
		"timestamp":  errorResponse.Timestamp,
		"request_id": errorResponse.RequestID,
	})
}
