# ⚡ Vercel Deployment Guide

## Overview
Vercel adalah platform hosting yang sangat cocok untuk SvelteKit frontend dengan auto-deploy dan edge functions.

## Setup Steps

### 1. Prepare Frontend Repository
```bash
# Pastikan struktur project sudah benar
frontend/
├── src/
├── static/
├── package.json
├── svelte.config.js
├── vite.config.ts
└── tailwind.config.js
```

### 2. Create Vercel Account
1. Buka [vercel.com](https://vercel.com)
2. Sign up dengan GitHub
3. Create new project

### 3. Connect GitHub Repository
1. Klik "Import Git Repository"
2. Pilih repository `job_portal-Roomah`
3. Vercel akan auto-detect SvelteKit project

### 4. Configure Build Settings
```bash
# Framework Preset: SvelteKit
# Build Command: npm run build
# Output Directory: build
# Install Command: npm install
```

### 5. Set Environment Variables
```env
# API URL untuk backend
VITE_API_URL=https://your-backend-url.railway.app/api

# Production settings
NODE_ENV=production
```

### 6. Configure Project Settings

#### Build Configuration
```json
{
  "buildCommand": "npm run build",
  "outputDirectory": "build",
  "installCommand": "npm install",
  "framework": "sveltekit"
}
```

#### Vercel Configuration (vercel.json)
```json
{
  "buildCommand": "npm run build",
  "outputDirectory": "build",
  "installCommand": "npm install",
  "framework": "sveltekit",
  "functions": {
    "app/api/**/*.js": {
      "runtime": "nodejs18.x"
    }
  }
}
```

## Environment Variables Setup

### Development (.env.local)
```env
VITE_API_URL=http://localhost:8082/api
```

### Production (Vercel Dashboard)
```env
VITE_API_URL=https://your-backend-url.railway.app/api
NODE_ENV=production
```

## API Configuration

### Update API Client
```typescript
// src/lib/api.ts
const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8082/api';

export const api = {
  async getJobs(filters: JobFilter = {}, page: number = 1, limit: number = 12): Promise<PaginatedResponse> {
    const params = new URLSearchParams();
    
    if (filters.location) params.append('location', filters.location);
    if (filters.salaryMin) params.append('salary_min', filters.salaryMin.toString());
    if (filters.salaryMax) params.append('salary_max', filters.salaryMax.toString());
    params.append('page', page.toString());
    params.append('limit', limit.toString());

    const response = await fetch(`${API_BASE_URL}/jobs?${params}`);
    if (!response.ok) {
      throw new Error('Failed to fetch jobs');
    }
    return response.json();
  },
  // ... other methods
};
```

## Build Optimization

### SvelteKit Configuration
```javascript
// svelte.config.js
import adapter from '@sveltejs/adapter-auto';

/** @type {import('@sveltejs/kit').Config} */
const config = {
  kit: {
    adapter: adapter(),
    prerender: {
      handleHttpError: 'warn'
    }
  }
};

export default config;
```

### Vite Configuration
```typescript
// vite.config.ts
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
  plugins: [sveltekit()],
  build: {
    target: 'esnext',
    rollupOptions: {
      output: {
        manualChunks: {
          vendor: ['svelte', 'svelte/internal']
        }
      }
    }
  }
});
```

## Performance Optimization

### 1. Image Optimization
```html
<!-- Use Vercel's image optimization -->
<img src="/api/image?url=..." alt="..." />
```

### 2. Static Assets
```bash
# Place static assets in static/ directory
static/
├── favicon.ico
├── robots.txt
└── images/
```

### 3. Caching Strategy
```javascript
// +layout.server.js
export const prerender = true;
export const ssr = true;
```

## CORS Configuration

### Backend CORS Update
```go
// Update backend CORS to allow Vercel domain
func CORSMiddleware() gin.HandlerFunc {
    return cors.New(cors.Config{
        AllowOrigins:     []string{
            "https://your-app.vercel.app",
            "http://localhost:5173", // development
        },
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    })
}
```

## Custom Domain Setup

### 1. Add Custom Domain
1. Go to Vercel dashboard
2. Click on your project
3. Go to "Settings" → "Domains"
4. Add your custom domain

### 2. Configure DNS
```bash
# Add CNAME record
your-domain.com CNAME your-app.vercel.app
```

## Monitoring & Analytics

### Vercel Analytics
```javascript
// Enable Vercel Analytics
import { inject } from '@vercel/analytics';

inject();
```

### Performance Monitoring
1. Go to Vercel dashboard
2. Click on your project
3. Go to "Analytics" tab
4. Monitor Core Web Vitals

## Troubleshooting

### Common Issues

#### 1. Build Failures
```bash
# Check dependencies
npm install

# Clear cache
npm run build --force

# Check for TypeScript errors
npm run check
```

#### 2. API Connection Issues
```bash
# Verify environment variables
echo $VITE_API_URL

# Test API endpoint
curl https://your-backend-url.railway.app/api/jobs
```

#### 3. CORS Issues
```bash
# Check browser console for CORS errors
# Verify backend CORS configuration
```

### Debug Commands
```bash
# Local development
npm run dev

# Build testing
npm run build

# Preview build
npm run preview
```

## Production Checklist

### Before Deploy
- [ ] Environment variables set
- [ ] API URL updated
- [ ] CORS configured
- [ ] Build passes locally
- [ ] Tests passing

### After Deploy
- [ ] Check live site functionality
- [ ] Test API connections
- [ ] Verify responsive design
- [ ] Check performance scores
- [ ] Monitor error logs

## Performance Tips

### 1. Code Splitting
```javascript
// Use dynamic imports
const Component = await import('./Component.svelte');
```

### 2. Image Optimization
```html
<!-- Use next-gen formats -->
<img src="image.webp" alt="..." />
```

### 3. Bundle Analysis
```bash
# Analyze bundle size
npm run build
# Check build/ directory for bundle analysis
```

## Security Considerations

### 1. Environment Variables
- Never expose sensitive data in client-side code
- Use `VITE_` prefix for public variables
- Keep API keys server-side

### 2. Content Security Policy
```html
<!-- Add CSP headers -->
<meta http-equiv="Content-Security-Policy" content="default-src 'self'">
```

### 3. HTTPS
- Vercel provides automatic HTTPS
- Redirect HTTP to HTTPS

## Cost Optimization

### Vercel Pricing
- Free tier: 100GB bandwidth/month
- Job Portal frontend: ~$0/month (free tier sufficient)
- Custom domains: Free

### Tips to Reduce Costs
1. Optimize images
2. Use efficient bundling
3. Implement caching
4. Monitor bandwidth usage

## Next Steps

1. **Deploy Backend** to Railway
2. **Update API URL** in Vercel environment variables
3. **Test Full Application**
4. **Set up Custom Domain** (optional)
5. **Configure Analytics** and monitoring

## Integration with Backend

### Update API Configuration
```typescript
// After backend deployment, update VITE_API_URL
VITE_API_URL=https://job-portal-backend.railway.app/api
```

### Test Integration
```bash
# Test API connection
curl https://job-portal-backend.railway.app/api/jobs

# Test frontend with backend
# Open your Vercel app and test all features
```

---

**⚡ Vercel is perfect for SvelteKit frontends!** 