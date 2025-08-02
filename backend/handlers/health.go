package handlers

import (
	"job-portal-backend/database"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

// HealthResponse represents health check response
type HealthResponse struct {
	Status    string         `json:"status"`
	Timestamp time.Time      `json:"timestamp"`
	Version   string         `json:"version"`
	Uptime    string         `json:"uptime"`
	Database  DatabaseHealth `json:"database"`
	System    SystemHealth   `json:"system"`
	Cache     CacheHealth    `json:"cache"`
}

// DatabaseHealth represents database health status
type DatabaseHealth struct {
	Status    string                 `json:"status"`
	Connected bool                   `json:"connected"`
	Stats     map[string]interface{} `json:"stats,omitempty"`
	Error     string                 `json:"error,omitempty"`
}

// SystemHealth represents system health status
type SystemHealth struct {
	MemoryUsage MemoryUsage `json:"memory_usage"`
	Goroutines  int         `json:"goroutines"`
	CPUUsage    float64     `json:"cpu_usage"`
	DiskUsage   DiskUsage   `json:"disk_usage"`
}

// MemoryUsage represents memory usage statistics
type MemoryUsage struct {
	Alloc      uint64 `json:"alloc"`
	TotalAlloc uint64 `json:"total_alloc"`
	Sys        uint64 `json:"sys"`
	NumGC      uint32 `json:"num_gc"`
	HeapAlloc  uint64 `json:"heap_alloc"`
	HeapSys    uint64 `json:"heap_sys"`
}

// DiskUsage represents disk usage statistics
type DiskUsage struct {
	Total   uint64  `json:"total"`
	Used    uint64  `json:"used"`
	Free    uint64  `json:"free"`
	Percent float64 `json:"percent"`
}

// CacheHealth represents cache health status
type CacheHealth struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

var startTime = time.Now()

// HealthCheck godoc
// @Summary Health check endpoint
// @Description Comprehensive health check for load balancers
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	// Check database health
	dbHealth := checkDatabaseHealth()

	// Check system health
	sysHealth := checkSystemHealth()

	// Check cache health
	cacheHealth := checkCacheHealth()

	// Determine overall status
	status := "healthy"
	if dbHealth.Status != "healthy" || cacheHealth.Status != "healthy" {
		status = "unhealthy"
	}

	response := HealthResponse{
		Status:    status,
		Timestamp: time.Now(),
		Version:   "1.0.0",
		Uptime:    time.Since(startTime).String(),
		Database:  dbHealth,
		System:    sysHealth,
		Cache:     cacheHealth,
	}

	// Set appropriate status code
	if status == "healthy" {
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(http.StatusServiceUnavailable, response)
	}
}

// LivenessCheck godoc
// @Summary Liveness check endpoint
// @Description Simple liveness check for Kubernetes
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health/live [get]
func LivenessCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "alive",
		"timestamp": time.Now(),
		"uptime":    time.Since(startTime).String(),
	})
}

// ReadinessCheck godoc
// @Summary Readiness check endpoint
// @Description Readiness check for Kubernetes
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health/ready [get]
func ReadinessCheck(c *gin.Context) {
	// Check if database is ready
	dbHealth := checkDatabaseHealth()

	status := "ready"
	if dbHealth.Status != "healthy" {
		status = "not_ready"
	}

	response := gin.H{
		"status":    status,
		"timestamp": time.Now(),
		"database":  dbHealth.Status,
	}

	if status == "ready" {
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(http.StatusServiceUnavailable, response)
	}
}

// checkDatabaseHealth checks database connection and returns health status
func checkDatabaseHealth() DatabaseHealth {
	// Test database connection
	err := database.DB.Ping()
	if err != nil {
		return DatabaseHealth{
			Status:    "unhealthy",
			Connected: false,
			Error:     err.Error(),
		}
	}

	// Get database statistics
	stats := database.GetDBStats()

	return DatabaseHealth{
		Status:    "healthy",
		Connected: true,
		Stats:     stats,
	}
}

// checkSystemHealth checks system resources and returns health status
func checkSystemHealth() SystemHealth {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return SystemHealth{
		MemoryUsage: MemoryUsage{
			Alloc:      m.Alloc,
			TotalAlloc: m.TotalAlloc,
			Sys:        m.Sys,
			NumGC:      m.NumGC,
			HeapAlloc:  m.HeapAlloc,
			HeapSys:    m.HeapSys,
		},
		Goroutines: runtime.NumGoroutine(),
		CPUUsage:   0.0, // Would need external library for CPU usage
		DiskUsage: DiskUsage{
			Total:   0, // Would need external library for disk usage
			Used:    0,
			Free:    0,
			Percent: 0.0,
		},
	}
}

// checkCacheHealth checks cache status
func checkCacheHealth() CacheHealth {
	// For now, return healthy status
	// In production, this would check Redis or other cache
	return CacheHealth{
		Status: "healthy",
	}
}

// Metrics godoc
// @Summary Metrics endpoint
// @Description Application metrics for monitoring
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /metrics [get]
func Metrics(c *gin.Context) {
	// Get database stats
	dbStats := database.GetDBStats()

	// Get system stats
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	metrics := gin.H{
		"timestamp": time.Now(),
		"database":  dbStats,
		"system": gin.H{
			"goroutines":     runtime.NumGoroutine(),
			"memory_alloc":   m.Alloc,
			"memory_sys":     m.Sys,
			"memory_heap":    m.HeapAlloc,
			"gc_num":         m.NumGC,
			"uptime_seconds": time.Since(startTime).Seconds(),
		},
		"application": gin.H{
			"version": "1.0.0",
			"status":  "running",
		},
	}

	c.JSON(http.StatusOK, metrics)
}
