# Swagger Documentation Guide

## Overview

Job Portal API menggunakan Swagger/OpenAPI untuk dokumentasi API yang interaktif dan lengkap.

## Akses Swagger UI

Setelah server berjalan, akses Swagger UI di:
- **URL**: http://localhost:8082/swagger/index.html
- **JSON**: http://localhost:8082/swagger/doc.json
- **YAML**: http://localhost:8082/swagger/swagger.yaml

## Fitur Swagger UI

### 1. Interactive Documentation
- Dokumentasi API yang interaktif
- Testing endpoint langsung dari browser
- Contoh request dan response
- Schema validation

### 2. Endpoint Testing
- Try it out functionality
- Parameter input yang mudah
- Response preview
- Error handling

### 3. Model Documentation
- Job model
- Application model
- Pagination model
- Error responses

## Cara Menggunakan Swagger UI

### 1. Buka Swagger UI
```
http://localhost:8082/swagger/index.html
```

### 2. Explore Endpoints
- Jobs endpoints (GET, POST)
- Applications endpoints (GET, POST)
- Locations endpoint (GET)

### 3. Test Endpoints

#### Test Get Jobs
1. Klik endpoint `GET /api/jobs`
2. Klik "Try it out"
3. Isi parameter (opsional):
   - `page`: 1
   - `limit`: 12
   - `location`: Jakarta
   - `salary_min`: 3000000
4. Klik "Execute"
5. Lihat response

#### Test Create Job
1. Klik endpoint `POST /api/jobs`
2. Klik "Try it out"
3. Isi request body:
```json
{
  "position": "Test Developer",
  "company": "Test Company",
  "location": "Jakarta",
  "salary_min": 5000000,
  "salary_max": 8000000
}
```
4. Klik "Execute"
5. Lihat response

#### Test Submit Application
1. Klik endpoint `POST /api/applications`
2. Klik "Try it out"
3. Isi form data:
   - `name`: John Doe
   - `email`: john@example.com
   - `job_id`: 1
   - `cv`: Upload file PDF
4. Klik "Execute"
5. Lihat response

## Generate Swagger Documentation

### Install Swag CLI
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### Generate Documentation
```bash
~/go/bin/swag init
```

### Regenerate After Changes
Setelah menambah atau mengubah endpoint, regenerate dokumentasi:
```bash
~/go/bin/swag init
```

## Swagger Annotations

### Main API Info
```go
// @title Job Portal API
// @version 1.0
// @description API untuk Job Portal - Platform lowongan kerja terbaik di Indonesia
// @host localhost:8082
// @BasePath /api
```

### Endpoint Documentation
```go
// @Summary Get all jobs with pagination and filters
// @Description Retrieve a list of jobs with optional filtering and pagination
// @Tags jobs
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param limit query int false "Number of items per page (default: 12, max: 50)"
// @Success 200 {object} PaginatedResponse
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /jobs [get]
```

### Model Documentation
```go
// Job represents a job posting
// @Description Job posting information
type Job struct {
    ID        int       `json:"id" example:"1"`
    Position  string    `json:"position" example:"Frontend Developer"`
    Company   string    `json:"company" example:"TechCorp Indonesia"`
    Location  string    `json:"location" example:"Jakarta"`
    SalaryMin int       `json:"salary_min" example:"3000000"`
    SalaryMax int       `json:"salary_max" example:"5000000"`
    CreatedAt time.Time `json:"created_at" example:"2025-01-15T10:30:00Z"`
}
```

## Best Practices

### 1. Documentation Standards
- Gunakan deskripsi yang jelas dan lengkap
- Sertakan contoh request dan response
- Dokumentasikan semua parameter
- Jelaskan error responses

### 2. Model Documentation
- Tambahkan `@Description` untuk setiap model
- Gunakan `example` tag untuk contoh nilai
- Dokumentasikan semua field

### 3. Endpoint Documentation
- Kelompokkan endpoint dengan `@Tags`
- Jelaskan parameter dengan detail
- Sertakan response schema
- Dokumentasikan error cases

### 4. File Upload
- Gunakan `multipart/form-data` untuk file upload
- Dokumentasikan file size limits
- Jelaskan supported file types

## Troubleshooting

### Common Issues

#### 1. Swagger UI tidak muncul
- Pastikan server berjalan di port 8082
- Cek apakah endpoint `/swagger/*any` terdaftar
- Pastikan import `_ "job-portal-backend/docs"` ada di main.go

#### 2. Documentation tidak update
- Jalankan `swag init` setelah perubahan
- Restart server
- Clear browser cache

#### 3. Model tidak muncul di Swagger
- Pastikan struct memiliki tag `json`
- Tambahkan `@Description` untuk model
- Gunakan `example` tag untuk contoh nilai

### Debug Commands

```bash
# Check if swag is installed
which swag

# Generate docs with verbose output
swag init --parseDependency --parseInternal

# Check generated files
ls -la docs/

# Validate swagger.json
curl -s http://localhost:8082/swagger/doc.json | jq '.'
```

## Integration with CI/CD

### GitHub Actions Example
```yaml
name: Generate Swagger Docs
on: [push]
jobs:
  generate-docs:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2
    - uses: actions/setup-go@v2
      with:
        go-version: '1.21'
    - run: go install github.com/swaggo/swag/cmd/swag@latest
    - run: swag init
    - run: git add docs/
    - run: git commit -m "Update swagger docs" || true
```

## Security Considerations

### 1. Production Environment
- Disable Swagger UI di production
- Gunakan environment variable untuk control
- Implementasi authentication untuk docs

### 2. API Security
- Validasi semua input
- Implementasi rate limiting
- Gunakan HTTPS
- Sanitasi file uploads

## Advanced Features

### 1. Custom Swagger UI
```go
import (
    "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
)

// Custom Swagger UI configuration
r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
    ginSwagger.URL("http://localhost:8082/swagger/doc.json"),
    ginSwagger.DefaultModelsExpandDepth(-1),
))
```

### 2. Multiple API Versions
```go
// @title Job Portal API v1
// @version 1.0
// @BasePath /api/v1

// @title Job Portal API v2
// @version 2.0
// @BasePath /api/v2
```

### 3. Authentication Documentation
```go
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
```

## Resources

- [Swag Documentation](https://github.com/swaggo/swag)
- [Gin Swagger](https://github.com/swaggo/gin-swagger)
- [OpenAPI Specification](https://swagger.io/specification/)
- [Swagger UI](https://swagger.io/tools/swagger-ui/) 