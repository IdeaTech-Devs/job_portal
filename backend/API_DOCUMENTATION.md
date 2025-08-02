# Job Portal API Documentation

## Overview

Job Portal API adalah RESTful API untuk platform lowongan kerja yang menyediakan endpoint untuk mengelola jobs dan applications.

**Base URL**: `http://localhost:8082/api`

## Authentication

Saat ini API tidak memerlukan authentication, namun sudah disiapkan untuk implementasi JWT di masa depan.

## Response Format

Semua response menggunakan format JSON dengan struktur:

```json
{
  "data": {...},
  "message": "Success message",
  "error": "Error message (if any)"
}
```

## Endpoints

### 1. Jobs

#### Get All Jobs
```
GET /api/jobs
```

**Query Parameters:**
- `page` (int, optional): Nomor halaman (default: 1)
- `limit` (int, optional): Jumlah item per halaman (default: 12, max: 50)
- `location` (string, optional): Filter berdasarkan lokasi
- `salary_min` (int, optional): Filter gaji minimum
- `salary_max` (int, optional): Filter gaji maksimum

**Response:**
```json
{
  "jobs": [
    {
      "id": 1,
      "position": "Frontend Developer",
      "company": "TechCorp Indonesia",
      "location": "Jakarta",
      "salary_min": 3000000,
      "salary_max": 5000000,
      "created_at": "2025-01-15T10:30:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 12,
    "total": 50,
    "total_pages": 5,
    "has_next": true,
    "has_prev": false
  }
}
```

**Contoh Request:**
```bash
curl "http://localhost:8082/api/jobs?page=1&limit=12&location=Jakarta&salary_min=3000000"
```

#### Get Job by ID
```
GET /api/jobs/{id}
```

**Response:**
```json
{
  "id": 1,
  "position": "Frontend Developer",
  "company": "TechCorp Indonesia",
  "location": "Jakarta",
  "salary_min": 3000000,
  "salary_max": 5000000,
  "created_at": "2025-01-15T10:30:00Z"
}
```

#### Create Job
```
POST /api/jobs
```

**Request Body:**
```json
{
  "position": "Backend Developer",
  "company": "Digital Solutions",
  "location": "Surabaya",
  "salary_min": 4000000,
  "salary_max": 7000000
}
```

**Response:**
```json
{
  "id": 2,
  "position": "Backend Developer",
  "company": "Digital Solutions",
  "location": "Surabaya",
  "salary_min": 4000000,
  "salary_max": 7000000,
  "created_at": "2025-01-15T11:00:00Z"
}
```

#### Get Locations
```
GET /api/locations
```

**Response:**
```json
[
  "Jakarta",
  "Surabaya",
  "Bandung",
  "Yogyakarta",
  "Bali"
]
```

### 2. Applications

#### Submit Application
```
POST /api/applications
```

**Content-Type:** `multipart/form-data`

**Form Data:**
- `name` (string, required): Nama pelamar
- `email` (string, required): Email pelamar
- `job_id` (int, required): ID job yang dilamar
- `cv` (file, required): File CV dalam format PDF (max 5MB)

**Response:**
```json
{
  "message": "Application submitted successfully",
  "application": {
    "id": 1,
    "job_id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com",
    "cv_filename": "cv_1234567890.pdf",
    "applied_at": "2025-01-15T10:30:00Z"
  }
}
```

**Contoh Request:**
```bash
curl -X POST "http://localhost:8082/api/applications" \
  -F "name=John Doe" \
  -F "email=john@example.com" \
  -F "job_id=1" \
  -F "cv=@/path/to/cv.pdf"
```

#### Get All Applications
```
GET /api/applications
```

**Response:**
```json
[
  {
    "id": 1,
    "job_id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com",
    "cv_filename": "cv_1234567890.pdf",
    "applied_at": "2025-01-15T10:30:00Z",
    "job": {
      "id": 1,
      "position": "Frontend Developer",
      "company": "TechCorp Indonesia",
      "location": "Jakarta",
      "salary_min": 3000000,
      "salary_max": 5000000,
      "created_at": "2025-01-15T10:30:00Z"
    }
  }
]
```

#### Get Application by ID
```
GET /api/applications/{id}
```

**Response:**
```json
{
  "id": 1,
  "job_id": 1,
  "name": "John Doe",
  "email": "john.doe@example.com",
  "cv_filename": "cv_1234567890.pdf",
  "applied_at": "2025-01-15T10:30:00Z",
  "job": {
    "id": 1,
    "position": "Frontend Developer",
    "company": "TechCorp Indonesia",
    "location": "Jakarta",
    "salary_min": 3000000,
    "salary_max": 5000000,
    "created_at": "2025-01-15T10:30:00Z"
  }
}
```

## Error Responses

### 400 Bad Request
```json
{
  "error": "Invalid request data"
}
```

### 404 Not Found
```json
{
  "error": "Job not found"
}
```

### 500 Internal Server Error
```json
{
  "error": "Failed to fetch jobs"
}
```

## Data Models

### Job
```json
{
  "id": "integer",
  "position": "string",
  "company": "string",
  "location": "string",
  "salary_min": "integer",
  "salary_max": "integer",
  "created_at": "datetime"
}
```

### Application
```json
{
  "id": "integer",
  "job_id": "integer",
  "name": "string",
  "email": "string",
  "cv_filename": "string",
  "applied_at": "datetime",
  "job": "Job object (optional)"
}
```

### Pagination
```json
{
  "page": "integer",
  "limit": "integer",
  "total": "integer",
  "total_pages": "integer",
  "has_next": "boolean",
  "has_prev": "boolean"
}
```

## Rate Limiting

Saat ini tidak ada rate limiting, namun direkomendasikan untuk implementasi di production.

## File Upload

- **Supported formats**: PDF only
- **Maximum size**: 5MB
- **Storage**: Local filesystem (`uploads/` directory)
- **Naming**: Auto-generated unique filename

## CORS

API mendukung CORS untuk cross-origin requests dari frontend.

## Testing

### Menggunakan curl

```bash
# Get jobs
curl "http://localhost:8082/api/jobs"

# Get jobs with filters
curl "http://localhost:8082/api/jobs?location=Jakarta&salary_min=3000000"

# Submit application
curl -X POST "http://localhost:8082/api/applications" \
  -F "name=Test User" \
  -F "email=test@example.com" \
  -F "job_id=1" \
  -F "cv=@test.pdf"
```

### Menggunakan Swagger UI

Akses http://localhost:8082/swagger/index.html untuk testing interaktif.

## Production Considerations

1. **Authentication**: Implementasi JWT atau OAuth2
2. **Rate Limiting**: Implementasi rate limiting
3. **File Storage**: Gunakan cloud storage (AWS S3, Google Cloud Storage)
4. **Database**: Optimasi query dan indexing
5. **Logging**: Implementasi structured logging
6. **Monitoring**: Health checks dan metrics
7. **Security**: HTTPS, input validation, SQL injection protection 