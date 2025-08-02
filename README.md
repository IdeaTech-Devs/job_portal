# 🚀 Job Portal - Platform Lowongan Kerja Terbaik

> **📋 Test Technical Fullstack Developer - Roomah**  
> Project ini dibuat sebagai test technical untuk posisi Fullstack Developer di Roomah. Aplikasi ini menampilkan kemampuan dalam pengembangan frontend (SvelteKit) dan backend (Go/Gin) dengan fitur-fitur modern dan optimasi untuk skalabilitas.

Job Portal adalah aplikasi web modern untuk platform lowongan kerja yang menyediakan fitur pencarian, filter, dan aplikasi lowongan dengan antarmuka yang clean dan responsif, serta optimasi performa untuk menangani traffic tinggi.

## 🎯 Test Technical Requirements

### Frontend Requirements ✅
- [x] **Modern Framework** - Menggunakan SvelteKit dengan TypeScript
- [x] **Responsive Design** - Optimal di semua device (desktop, tablet, mobile)
- [x] **Clean UI/UX** - Interface minimalis dan user-friendly
- [x] **Real-time Search** - Pencarian instan dengan filter
- [x] **Pagination** - Navigasi halaman yang smooth
- [x] **Component Architecture** - Reusable components yang modular
- [x] **State Management** - Menggunakan Svelte stores
- [x] **API Integration** - Custom API client dengan error handling

### Backend Requirements ✅
- [x] **Modern Framework** - Menggunakan Gin (Go) dengan struktur yang clean
- [x] **Database Integration** - PostgreSQL dengan ORM-like patterns
- [x] **RESTful API** - Endpoints yang terstruktur dan konsisten
- [x] **Pagination** - Efficient pagination untuk data besar
- [x] **File Upload** - CV upload dengan validasi dan security
- [x] **API Documentation** - Swagger/OpenAPI documentation
- [x] **Error Handling** - Proper error responses dan logging
- [x] **CORS Support** - Cross-origin requests untuk frontend

### Security & Stability Features ✅
- [x] **Input Validation** - Comprehensive validation untuk semua endpoints
- [x] **Rate Limiting** - Protect API dari abuse dengan multiple tiers
- [x] **Error Handling** - Structured logging dan error tracking
- [x] **Database Indexing** - Optimasi query performance dengan indexes
- [x] **XSS Protection** - Input sanitization untuk mencegah serangan
- [x] **Request ID Tracking** - Unique request ID untuk debugging

### Scaling & Performance Features ✅
- [x] **Database Connection Pool** - Configure proper connection limits
- [x] **Health Checks** - Required untuk load balancers dan monitoring
- [x] **Caching Strategy** - Redis untuk frequently accessed data
- [x] **Database Indexes** - Optimized queries dengan proper indexing
- [x] **Rate Limiting** - Multiple tiers untuk different endpoints
- [x] **Metrics & Monitoring** - Application metrics untuk observability

### Additional Features ✅
- [x] **Sample Data** - 50+ realistic job postings
- [x] **Testing** - API testing script
- [x] **Documentation** - Comprehensive README dan API docs
- [x] **Code Quality** - Clean code dengan proper structure
- [x] **Performance** - Optimized queries dan frontend performance

## ✨ Fitur Utama

### Frontend (SvelteKit)
- 🎨 **UI/UX Modern & Clean** - Tampilan minimalis dengan desain yang elegan
- 📱 **Responsive Design** - Optimal di desktop, tablet, dan mobile
- 🔍 **Real-time Search** - Pencarian instan dengan filter yang canggih
- 📄 **Pagination** - Navigasi halaman yang smooth dengan 12 item per halaman
- 🎯 **Advanced Filtering** - Filter berdasarkan lokasi dan rentang gaji
- 📝 **Application System** - Sistem lamaran dengan upload CV
- ⚡ **Performance Optimized** - Loading cepat dan smooth interactions

### Backend (Go/Gin)
- 🏗️ **RESTful API** - API yang terstruktur dan mudah digunakan
- 📊 **PostgreSQL Database** - Database yang robust dan scalable
- 📚 **Swagger Documentation** - Dokumentasi API yang lengkap dan interaktif
- 🔒 **File Upload** - Upload CV dengan validasi dan security
- 📈 **Pagination Support** - Pagination yang efisien untuk data besar
- 🌐 **CORS Support** - Cross-origin requests untuk frontend

### Security & Stability
- 🛡️ **Input Validation** - Comprehensive validation untuk semua input
- 🚦 **Rate Limiting** - Multiple tiers: General (100/min), Search (200/min), Applications (5/min)
- 📝 **Structured Logging** - Request tracking dengan unique ID
- 🔍 **Error Tracking** - Detailed error responses dengan context
- 🗄️ **Database Indexing** - Optimized queries dengan proper indexes
- 🧹 **Input Sanitization** - XSS protection dan injection prevention

### Scaling & Performance
- 🔗 **Connection Pooling** - Database connection pool dengan configurable limits
- 🏥 **Health Checks** - Comprehensive health checks untuk load balancers
- 💾 **Redis Caching** - Frequently accessed data caching
- 📊 **Metrics & Monitoring** - Application metrics untuk observability
- ⚡ **Performance Optimization** - Optimized database queries dan caching strategy

## 🛠️ Tech Stack

### Frontend
- **Framework**: SvelteKit
- **Styling**: Tailwind CSS
- **Icons**: Lucide Svelte
- **Build Tool**: Vite
- **Language**: TypeScript

### Backend
- **Framework**: Gin (Go)
- **Database**: PostgreSQL
- **Cache**: Redis
- **Documentation**: Swagger/OpenAPI
- **File Upload**: Multipart form data
- **Language**: Go

### Infrastructure & Scaling
- **Connection Pooling**: Database connection pool dengan configurable limits
- **Caching**: Redis untuk frequently accessed data
- **Health Checks**: Comprehensive health checks untuk load balancers
- **Rate Limiting**: Multiple tiers untuk different endpoints
- **Monitoring**: Application metrics dan database statistics

## 📁 Struktur Project

```
job_portal-Roomah/
├── frontend/                 # SvelteKit Frontend
│   ├── src/
│   │   ├── lib/
│   │   │   ├── components/   # Reusable components
│   │   │   ├── stores/       # Svelte stores
│   │   │   └── api.ts        # API client
│   │   ├── routes/           # Pages
│   │   └── app.css          # Global styles
│   ├── package.json
│   └── README.md
├── backend/                  # Go Backend
│   ├── handlers/            # HTTP handlers
│   ├── models/              # Database models
│   ├── database/            # Database connection & indexes
│   ├── middleware/          # Security & validation middleware
│   ├── cache/              # Redis caching layer
│   ├── docs/               # Swagger docs
│   ├── uploads/            # File uploads
│   ├── main.go             # Entry point
│   └── README.md
└── README.md               # This file
```

## 🚀 Quick Start

### Prerequisites
- **Node.js** 18+ (untuk frontend)
- **Go** 1.21+ (untuk backend)
- **PostgreSQL** (untuk database)
- **Redis** (untuk caching, optional)

### 1. Clone Repository
```bash
git clone https://github.com/IdeaTech-Devs/job_portal.git
cd job_portal
```

### 2. Setup Backend

```bash
cd backend

# Install dependencies
go mod tidy

# Setup environment variables
cp env.example .env
# Edit .env sesuai konfigurasi database dan Redis

# Setup database (buat database PostgreSQL)
# Update konfigurasi di .env

# Seed data (opsional)
go run main.go -seed

# Run backend server
go run main.go
```

Backend akan berjalan di `http://localhost:8082`

### 3. Setup Frontend

```bash
cd frontend

# Install dependencies
npm install

# Setup environment variables
cp env.example .env
# Edit .env untuk API URL

# Run development server
npm run dev
```

Frontend akan berjalan di `http://localhost:5173`

## 📚 API Documentation

### Swagger UI
Setelah backend berjalan, akses dokumentasi API di:
- **Swagger UI**: http://localhost:8082/swagger/index.html
- **API JSON**: http://localhost:8082/swagger/doc.json

### Health Check Endpoints
- **Health Check**: `GET /health` - Comprehensive health check
- **Liveness Check**: `GET /health/live` - Kubernetes liveness probe
- **Readiness Check**: `GET /health/ready` - Kubernetes readiness probe
- **Metrics**: `GET /metrics` - Application metrics

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

## 🔒 Security Features

### Input Validation
- **Comprehensive Validation** - Semua input divalidasi dengan regex patterns
- **File Upload Security** - Validasi file type dan size untuk CV upload
- **SQL Injection Prevention** - Parameterized queries
- **XSS Protection** - Input sanitization untuk mencegah serangan

### Rate Limiting
- **General API**: 100 requests per minute
- **Search Endpoints**: 200 requests per minute
- **Application Submission**: 5 requests per minute
- **Job Creation**: 10 requests per minute

### Error Handling
- **Structured Error Responses** - Consistent error format
- **Request ID Tracking** - Unique ID untuk setiap request
- **Detailed Logging** - Comprehensive logging untuk debugging
- **Graceful Degradation** - System tetap berjalan meski cache down

## ⚡ Performance & Scaling

### Database Optimization
- **Connection Pooling** - Configurable connection limits
- **Database Indexes** - Optimized queries dengan proper indexing
- **Query Optimization** - Efficient pagination dan filtering
- **Connection Statistics** - Monitoring database performance

### Caching Strategy
- **Redis Caching** - Frequently accessed data caching
- **Cache Keys** - Structured cache keys untuk different data types
- **Cache Expiration** - Configurable TTL untuk different data
- **Cache Invalidation** - Automatic cache invalidation pada data changes

### Health Checks & Monitoring
- **Comprehensive Health Checks** - Database, cache, dan system health
- **Application Metrics** - Memory usage, goroutines, uptime
- **Database Statistics** - Connection pool stats dan query performance
- **Load Balancer Ready** - Health endpoints untuk load balancers

## 🎨 UI/UX Features

### Design Principles
- **Clean & Minimalist** - Interface yang bersih dan tidak berlebihan
- **Responsive** - Optimal di semua device
- **Accessible** - Mengikuti standar accessibility
- **Fast** - Loading cepat dan smooth interactions

### Key Components
- **Hero Section** - Landing page yang menarik dengan search bar
- **Job Cards** - Card design yang modern untuk job listings
- **Filter Panel** - Filter yang mudah digunakan
- **Pagination** - Navigasi halaman yang intuitif
- **Application Modal** - Form lamaran yang user-friendly

## 🗄️ Database Schema

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

-- Indexes for performance
CREATE INDEX idx_jobs_location ON jobs(location);
CREATE INDEX idx_jobs_salary_range ON jobs(salary_min, salary_max);
CREATE INDEX idx_jobs_created_at ON jobs(created_at DESC);
CREATE INDEX idx_jobs_location_salary ON jobs(location, salary_min, salary_max);
CREATE INDEX idx_jobs_company ON jobs(company);
CREATE INDEX idx_jobs_position ON jobs(position);
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

-- Indexes for performance
CREATE INDEX idx_applications_job_id ON applications(job_id);
CREATE INDEX idx_applications_applied_at ON applications(applied_at DESC);
CREATE INDEX idx_applications_email ON applications(email);
CREATE INDEX idx_applications_job_applied ON applications(job_id, applied_at DESC);
```

## 🧪 Testing

### Backend Testing
```bash
cd backend
./test_api.sh
```

### Health Check Testing
```bash
# Test health endpoints
curl http://localhost:8082/health
curl http://localhost:8082/health/live
curl http://localhost:8082/health/ready
curl http://localhost:8082/metrics
```

### Frontend Testing
```bash
cd frontend
npm run test
```

## 📖 Development

### Backend Development
```bash
cd backend

# Generate Swagger docs
~/go/bin/swag init

# Run with hot reload
go run main.go

# Test API
./test_api.sh

# Check health endpoints
curl http://localhost:8082/health
```

### Frontend Development
```bash
cd frontend

# Install dependencies
npm install

# Run development server
npm run dev

# Build for production
npm run build
```

## 🔧 Configuration

### Environment Variables

#### Backend (.env)
```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=job_portal

# Database Connection Pool Configuration
DB_MAX_OPEN_CONNS=25
DB_MAX_IDLE_CONNS=5
DB_MAX_LIFETIME_MINUTES=5
DB_MAX_IDLE_TIME_MINUTES=5

# Server Configuration
PORT=8082
GIN_MODE=debug

# Redis Cache Configuration
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=

# For Railway deployment
DATABASE_URL=postgresql://username:password@host:port/database
```

#### Frontend (.env)
```env
VITE_API_URL=http://localhost:8082/api
```

## 📊 Sample Data

Backend sudah dilengkapi dengan 50+ sample jobs yang mencakup:
- Frontend & Web Development
- Backend Development
- Mobile Development
- DevOps & Infrastructure
- Data & Analytics
- Design & UX
- Product & Project Management
- Quality Assurance
- System & Network
- Emerging Technologies

## 🚀 Deployment

### Backend Deployment
```bash
# Build binary
go build -o job-portal-backend main.go

# Run in production
./job-portal-backend
```

### Frontend Deployment
```bash
# Build for production
npm run build
```

### Docker Deployment (Optional)
```bash
# Build and run with Docker
docker build -t job-portal-backend .
docker run -p 8082:8082 job-portal-backend
```

## 📈 Performance Metrics

### Database Performance
- **Connection Pool**: Configurable limits (default: 25 max, 5 idle)
- **Query Optimization**: Indexed queries untuk fast retrieval
- **Connection Statistics**: Monitoring pool performance

### Caching Performance
- **Redis Cache**: Frequently accessed data caching
- **Cache Hit Rate**: Optimized untuk reduce database load
- **Cache Expiration**: Configurable TTL untuk different data types

### API Performance
- **Rate Limiting**: Multiple tiers untuk prevent abuse
- **Response Time**: Optimized dengan caching dan indexing
- **Error Handling**: Graceful degradation untuk stability

## 🎯 Technical Assessment Summary

### Frontend Assessment
- ✅ **Modern Framework Usage** - SvelteKit dengan TypeScript
- ✅ **Component Architecture** - Modular dan reusable components
- ✅ **State Management** - Proper state handling dengan Svelte stores
- ✅ **API Integration** - Clean API client dengan error handling
- ✅ **Responsive Design** - Mobile-first approach
- ✅ **Performance** - Optimized loading dan interactions

### Backend Assessment
- ✅ **Modern Framework Usage** - Gin dengan Go
- ✅ **Database Design** - Proper schema dan relationships
- ✅ **API Design** - RESTful dengan proper HTTP methods
- ✅ **Error Handling** - Comprehensive error responses
- ✅ **Documentation** - Swagger/OpenAPI documentation
- ✅ **File Upload** - Secure file handling dengan validation

### Security & Stability Assessment
- ✅ **Input Validation** - Comprehensive validation untuk semua endpoints
- ✅ **Rate Limiting** - Multiple tiers untuk prevent abuse
- ✅ **Error Handling** - Structured logging dan error tracking
- ✅ **Database Indexing** - Optimized queries dengan proper indexes
- ✅ **XSS Protection** - Input sanitization untuk security
- ✅ **Request Tracking** - Unique request ID untuk debugging

### Scaling & Performance Assessment
- ✅ **Connection Pooling** - Database connection pool dengan configurable limits
- ✅ **Health Checks** - Comprehensive health checks untuk load balancers
- ✅ **Caching Strategy** - Redis untuk frequently accessed data
- ✅ **Database Indexes** - Optimized queries dengan proper indexing
- ✅ **Rate Limiting** - Multiple tiers untuk different endpoints
- ✅ **Metrics & Monitoring** - Application metrics untuk observability

### Code Quality
- ✅ **Clean Code** - Readable dan maintainable
- ✅ **Proper Structure** - Organized folder structure
- ✅ **Documentation** - Comprehensive README dan comments
- ✅ **Testing** - API testing script included
- ✅ **Version Control** - Proper git structure

---

**🎯 Test Technical Fullstack Developer - Roomah**  
*Project ini dibuat untuk menampilkan kemampuan dalam pengembangan fullstack modern dengan best practices, clean architecture, dan optimasi untuk skalabilitas.*
