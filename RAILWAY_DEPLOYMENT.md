# 🚂 Railway Deployment Guide

## Overview
Railway adalah platform hosting yang sangat cocok untuk Go backend dengan PostgreSQL database.

## Setup Steps

### 1. Prepare Repository
```bash
# Pastikan struktur project sudah benar
backend/
├── main.go
├── go.mod
├── go.sum
├── handlers/
├── models/
├── database/
├── middleware/
└── uploads/
```

### 2. Create Railway Account
1. Buka [railway.app](https://railway.app)
2. Sign up dengan GitHub
3. Create new project

### 3. Connect GitHub Repository
1. Klik "Deploy from GitHub repo"
2. Pilih repository `job_portal-Roomah`
3. Railway akan auto-detect Go project

### 4. Configure Environment Variables
```env
DB_HOST=your-postgresql-host
DB_PORT=5432
DB_USER=your-db-user
DB_PASSWORD=your-db-password
DB_NAME=your-db-name
PORT=8082
```

### 5. Add PostgreSQL Database
1. Klik "New" → "Database" → "PostgreSQL"
2. Railway akan generate connection string
3. Update environment variables dengan connection string

### 6. Configure Build Settings
```bash
# Build Command (optional, Railway auto-detect)
go build -o main .

# Start Command
./main
```

### 7. Deploy
1. Railway akan auto-deploy setiap push ke main branch
2. Monitor logs di Railway dashboard
3. Get your app URL (e.g., https://job-portal-backend.railway.app)

## Environment Variables Setup

### Database Configuration
```env
# Railway PostgreSQL connection
DATABASE_URL=postgresql://username:password@host:port/database

# Or separate variables
DB_HOST=containers-us-west-1.railway.app
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your-password
DB_NAME=railway
```

### App Configuration
```env
PORT=8082
GIN_MODE=release
```

## Database Migration

### Option 1: Manual Setup
```sql
-- Connect to Railway PostgreSQL
-- Run these SQL commands

CREATE TABLE jobs (
    id SERIAL PRIMARY KEY,
    position VARCHAR(255) NOT NULL,
    company VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL,
    salary_min INTEGER NOT NULL,
    salary_max INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE applications (
    id SERIAL PRIMARY KEY,
    job_id INTEGER REFERENCES jobs(id),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    cv_filename VARCHAR(255) NOT NULL,
    applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Option 2: Auto Seed
```bash
# Add seed flag to Railway start command
./main -seed
```

## File Upload Configuration

### Update Upload Path
```go
// In handlers/applications.go
uploadPath := fmt.Sprintf("/tmp/%s", filename) // Railway uses /tmp
```

### Or Use Cloud Storage
```go
// Consider using AWS S3 or similar for production
// For now, use /tmp directory
```

## CORS Configuration

### Update CORS Middleware
```go
// In middleware/cors.go
func CORSMiddleware() gin.HandlerFunc {
    return cors.New(cors.Config{
        AllowOrigins:     []string{"https://your-frontend-domain.vercel.app"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    })
}
```

## Monitoring & Logs

### View Logs
1. Go to Railway dashboard
2. Click on your service
3. Go to "Logs" tab
4. Monitor real-time logs

### Health Check
```bash
# Test your deployed API
curl https://your-app.railway.app/api/jobs
```

## Troubleshooting

### Common Issues

#### 1. Build Failures
```bash
# Check if all dependencies are in go.mod
go mod tidy
go mod download
```

#### 2. Database Connection Issues
```bash
# Verify DATABASE_URL format
# Check if database is accessible
```

#### 3. Port Issues
```bash
# Railway sets PORT environment variable
# Make sure your app uses os.Getenv("PORT")
```

### Debug Commands
```bash
# Local testing
go run main.go

# Build testing
go build -o main .
./main
```

## Production Considerations

### 1. Environment Variables
- Never commit sensitive data
- Use Railway's environment variable system
- Rotate database passwords regularly

### 2. Database Backups
- Railway provides automatic backups
- Consider additional backup strategies

### 3. Monitoring
- Set up alerts for downtime
- Monitor resource usage
- Track API performance

### 4. Security
- Enable HTTPS (automatic with Railway)
- Validate all inputs
- Implement rate limiting
- Use secure headers

## Cost Optimization

### Railway Pricing
- Free tier: $5 credit/month
- Job Portal backend: ~$2-3/month
- PostgreSQL: ~$1-2/month

### Tips to Reduce Costs
1. Optimize database queries
2. Use efficient pagination
3. Implement caching
4. Monitor resource usage

## Next Steps

1. **Deploy Frontend** to Vercel
2. **Update Frontend API URL** to Railway URL
3. **Test Full Application**
4. **Set up Custom Domain** (optional)
5. **Configure CI/CD** for automatic deployments

---

**🎯 Railway is perfect for Go backends with PostgreSQL!** 