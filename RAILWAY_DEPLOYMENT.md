# Railway Deployment Guide

## Konfigurasi Railway

### 1. Root Directory
Railway akan menggunakan root directory sebagai working directory. File konfigurasi:
- `railway.toml` - Konfigurasi Railway
- `nixpacks.toml` - Konfigurasi build

### 2. Environment Variables
Set di Railway dashboard:
- `DATABASE_URL=postgresql://postgres:ddjqKDcKWNAUYMUySinspZzpUUtLZOwB@nozomi.proxy.rlwy.net:37974/railway`
- `PORT=8082`
- `GIN_MODE=release`

### 3. Healthcheck
Railway akan menggunakan endpoint `/health/live` untuk healthcheck (tidak mengecek database).

### 4. Build Process
1. Railway akan masuk ke folder `backend/`
2. Menjalankan `go mod download`
3. Build dengan `go build -o job-portal-backend main.go`
4. Start dengan `./job-portal-backend`

### 5. Troubleshooting
- Jika healthcheck gagal, cek log di Railway dashboard
- Pastikan DATABASE_URL sudah benar
- Gunakan endpoint `/health/live` untuk healthcheck yang lebih toleran
