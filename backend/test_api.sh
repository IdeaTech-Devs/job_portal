#!/bin/bash

# Job Portal API Testing Script
# Usage: ./test_api.sh

BASE_URL="http://localhost:8082/api"

echo "🚀 Job Portal API Testing"
echo "=========================="
echo ""

# Test 1: Get all jobs
echo "📋 Test 1: Get all jobs"
curl -s "$BASE_URL/jobs" | jq '.' 2>/dev/null || curl -s "$BASE_URL/jobs"
echo ""
echo ""

# Test 2: Get jobs with pagination
echo "📋 Test 2: Get jobs with pagination"
curl -s "$BASE_URL/jobs?page=1&limit=5" | jq '.' 2>/dev/null || curl -s "$BASE_URL/jobs?page=1&limit=5"
echo ""
echo ""

# Test 3: Get jobs with filters
echo "📋 Test 3: Get jobs with filters (Jakarta, salary min 3M)"
curl -s "$BASE_URL/jobs?location=Jakarta&salary_min=3000000" | jq '.' 2>/dev/null || curl -s "$BASE_URL/jobs?location=Jakarta&salary_min=3000000"
echo ""
echo ""

# Test 4: Get specific job
echo "📋 Test 4: Get job by ID (ID: 1)"
curl -s "$BASE_URL/jobs/1" | jq '.' 2>/dev/null || curl -s "$BASE_URL/jobs/1"
echo ""
echo ""

# Test 5: Get locations
echo "📋 Test 5: Get all locations"
curl -s "$BASE_URL/locations" | jq '.' 2>/dev/null || curl -s "$BASE_URL/locations"
echo ""
echo ""

# Test 6: Get applications
echo "📋 Test 6: Get all applications"
curl -s "$BASE_URL/applications" | jq '.' 2>/dev/null || curl -s "$BASE_URL/applications"
echo ""
echo ""

# Test 7: Create a new job
echo "📋 Test 7: Create a new job"
curl -s -X POST "$BASE_URL/jobs" \
  -H "Content-Type: application/json" \
  -d '{
    "position": "Test Developer",
    "company": "Test Company",
    "location": "Jakarta",
    "salary_min": 5000000,
    "salary_max": 8000000
  }' | jq '.' 2>/dev/null || curl -s -X POST "$BASE_URL/jobs" \
  -H "Content-Type: application/json" \
  -d '{
    "position": "Test Developer",
    "company": "Test Company",
    "location": "Jakarta",
    "salary_min": 5000000,
    "salary_max": 8000000
  }'
echo ""
echo ""

echo "✅ API Testing completed!"
echo ""
echo "🌐 Swagger UI: http://localhost:8082/swagger/index.html"
echo "📚 API Documentation: http://localhost:8082/swagger/doc.json" 