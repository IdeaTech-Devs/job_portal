# Job Portal Backend API

Backend API untuk Job Portal - Platform lowongan kerja terbaik di Indonesia.

## Fitur

- ✅ CRUD Jobs (Create, Read, Update, Delete)
- ✅ Job Applications dengan file upload
- ✅ Pagination untuk job listings
- ✅ Filter berdasarkan lokasi dan gaji
- ✅ Swagger/OpenAPI Documentation
- ✅ CORS support
- ✅ PostgreSQL database

## Teknologi

- **Framework**: Gin (Go)
- **Database**: PostgreSQL
- **Documentation**: Swagger/OpenAPI
- **File Upload**: Multipart form data

## Instalasi

### Prerequisites

- Go 1.21+
- PostgreSQL

### Setup Database

1. Buat database PostgreSQL:
```sql
CREATE DATABASE job_portal;
```

2. Update konfigurasi database di `database/database.go`

### Install Dependencies

```bash
go mod tidy
```

### Seed Data (Opsional)

Untuk mengisi database dengan data sample:

```bash
go run main.go -seed
```

### Menjalankan Server

```bash
go run main.go
```

Server akan berjalan di `http://localhost:8082`

## API Documentation

### Swagger UI

Setelah server berjalan, Anda dapat mengakses dokumentasi API di:

- **Swagger UI**: http://localhost:8082/swagger/index.html
- **Swagger JSON**: http://localhost:8082/swagger/doc.json

### Endpoints

#### Jobs

- `GET /api/jobs` - Get all jobs dengan pagination dan filter
- `GET /api/jobs/{id}` - Get job by ID
- `POST /api/jobs` - Create new job
- `GET /api/locations` - Get all available locations

#### Applications

- `GET /api/applications` - Get all applications
- `GET /api/applications/{id}` - Get application by ID
- `POST /api/applications` - Submit job application dengan CV upload

### Query Parameters untuk Jobs

- `page` (int, default: 1) - Nomor halaman
- `limit` (int, default: 12, max: 50) - Jumlah item per halaman
- `location` (string) - Filter berdasarkan lokasi
- `salary_min` (int) - Filter gaji minimum
- `salary_max` (int) - Filter gaji maksimum

### Contoh Request

#### Get Jobs dengan Filter
```bash
curl "http://localhost:8082/api/jobs?page=1&limit=12&location=Jakarta&salary_min=3000000"
```

#### Submit Application
```bash
curl -X POST "http://localhost:8082/api/applications" \
  -F "name=John Doe" \
  -F "email=john@example.com" \
  -F "job_id=1" \
  -F "cv=@/path/to/cv.pdf"
```

## Struktur Database

### Jobs Table
```sql
CREATE TABLE jobs (
    id SERIAL PRIMARY KEY,
    position VARCHAR(255) NOT NULL,
    company VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL,
    salary_min INTEGER NOT NULL,
    salary_max INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Applications Table
```sql
CREATE TABLE applications (
    id SERIAL PRIMARY KEY,
    job_id INTEGER REFERENCES jobs(id),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    cv_filename VARCHAR(255) NOT NULL,
    applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Development

### Generate Swagger Documentation

Setelah menambah atau mengubah endpoint, regenerate dokumentasi:

```bash
~/go/bin/swag init
```

### Testing

```bash
go test ./...
```

## License

MIT License 