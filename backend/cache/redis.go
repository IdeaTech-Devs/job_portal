package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	redisClient *redis.Client
	ctx         = context.Background()
)

// CacheConfig represents cache configuration
type CacheConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

// InitRedis initializes Redis connection
func InitRedis() {
	// Get Redis configuration from environment
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("REDIS_PORT")
	if port == "" {
		port = "6379"
	}

	password := os.Getenv("REDIS_PASSWORD")
	db := 0

	// Create Redis client
	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       db,
		PoolSize: 10,
	})

	// Test connection
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Printf("Warning: Redis connection failed: %v", err)
		log.Println("Continuing without Redis cache")
		redisClient = nil
		return
	}

	log.Println("Redis cache initialized successfully")
}

// Set sets a key-value pair in cache
func Set(key string, value interface{}, expiration time.Duration) error {
	if redisClient == nil {
		return fmt.Errorf("Redis not available")
	}

	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return redisClient.Set(ctx, key, jsonValue, expiration).Err()
}

// Get retrieves a value from cache
func Get(key string, dest interface{}) error {
	if redisClient == nil {
		return fmt.Errorf("Redis not available")
	}

	val, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(val), dest)
}

// Delete removes a key from cache
func Delete(key string) error {
	if redisClient == nil {
		return fmt.Errorf("Redis not available")
	}

	return redisClient.Del(ctx, key).Err()
}

// Exists checks if a key exists in cache
func Exists(key string) bool {
	if redisClient == nil {
		return false
	}

	exists, err := redisClient.Exists(ctx, key).Result()
	return err == nil && exists > 0
}

// FlushAll clears all cache
func FlushAll() error {
	if redisClient == nil {
		return fmt.Errorf("Redis not available")
	}

	return redisClient.FlushAll(ctx).Err()
}

// GetStats returns Redis statistics
func GetStats() map[string]interface{} {
	if redisClient == nil {
		return map[string]interface{}{
			"status": "unavailable",
		}
	}

	info, err := redisClient.Info(ctx).Result()
	if err != nil {
		return map[string]interface{}{
			"status": "error",
			"error":  err.Error(),
		}
	}

	return map[string]interface{}{
		"status": "available",
		"info":   info,
	}
}

// Cache keys for different data types
const (
	// Job cache keys
	JobsCacheKey      = "jobs:all"
	JobCacheKey       = "job:%d"
	LocationsCacheKey = "locations:all"

	// Application cache keys
	ApplicationsCacheKey = "applications:all"
	ApplicationCacheKey  = "application:%d"

	// Cache expiration times
	JobsCacheExpiration         = 5 * time.Minute
	JobCacheExpiration          = 10 * time.Minute
	LocationsCacheExpiration    = 30 * time.Minute
	ApplicationsCacheExpiration = 2 * time.Minute
	ApplicationCacheExpiration  = 5 * time.Minute
)

// CacheJobs caches jobs data
func CacheJobs(jobs interface{}) error {
	return Set(JobsCacheKey, jobs, JobsCacheExpiration)
}

// GetCachedJobs retrieves cached jobs
func GetCachedJobs(dest interface{}) error {
	return Get(JobsCacheKey, dest)
}

// CacheJob caches individual job data
func CacheJob(jobID int, job interface{}) error {
	key := fmt.Sprintf(JobCacheKey, jobID)
	return Set(key, job, JobCacheExpiration)
}

// GetCachedJob retrieves cached job
func GetCachedJob(jobID int, dest interface{}) error {
	key := fmt.Sprintf(JobCacheKey, jobID)
	return Get(key, dest)
}

// CacheLocations caches locations data
func CacheLocations(locations interface{}) error {
	return Set(LocationsCacheKey, locations, LocationsCacheExpiration)
}

// GetCachedLocations retrieves cached locations
func GetCachedLocations(dest interface{}) error {
	return Get(LocationsCacheKey, dest)
}

// CacheApplications caches applications data
func CacheApplications(applications interface{}) error {
	return Set(ApplicationsCacheKey, applications, ApplicationsCacheExpiration)
}

// GetCachedApplications retrieves cached applications
func GetCachedApplications(dest interface{}) error {
	return Get(ApplicationsCacheKey, dest)
}

// CacheApplication caches individual application data
func CacheApplication(appID int, application interface{}) error {
	key := fmt.Sprintf(ApplicationCacheKey, appID)
	return Set(key, application, ApplicationCacheExpiration)
}

// GetCachedApplication retrieves cached application
func GetCachedApplication(appID int, dest interface{}) error {
	key := fmt.Sprintf(ApplicationCacheKey, appID)
	return Get(key, dest)
}

// InvalidateJobsCache invalidates jobs cache
func InvalidateJobsCache() error {
	return Delete(JobsCacheKey)
}

// InvalidateJobCache invalidates specific job cache
func InvalidateJobCache(jobID int) error {
	key := fmt.Sprintf(JobCacheKey, jobID)
	return Delete(key)
}

// InvalidateLocationsCache invalidates locations cache
func InvalidateLocationsCache() error {
	return Delete(LocationsCacheKey)
}

// InvalidateApplicationsCache invalidates applications cache
func InvalidateApplicationsCache() error {
	return Delete(ApplicationsCacheKey)
}

// InvalidateApplicationCache invalidates specific application cache
func InvalidateApplicationCache(appID int) error {
	key := fmt.Sprintf(ApplicationCacheKey, appID)
	return Delete(key)
}
