// @title Job Portal API
// @version 1.0
// @description API untuk Job Portal - Platform lowongan kerja terbaik di Indonesia
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8082
// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

package main

import (
	"flag"
	"job-portal-backend/database"
	"job-portal-backend/handlers"
	"job-portal-backend/middleware"
	"job-portal-backend/models"
	"log"
	"os"

	_ "job-portal-backend/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func seedData() {
	// Sample job data
	sampleJobs := []models.Job{
		// Frontend & Web Development
		{
			Position:  "Frontend Developer",
			Company:   "TechCorp Indonesia",
			Location:  "Jakarta",
			SalaryMin: 2000000,
			SalaryMax: 4000000,
		},
		{
			Position:  "React Developer",
			Company:   "WebTech Solutions",
			Location:  "Jakarta",
			SalaryMin: 3000000,
			SalaryMax: 5000000,
		},
		{
			Position:  "Vue.js Developer",
			Company:   "Digital Creative",
			Location:  "Bandung",
			SalaryMin: 2500000,
			SalaryMax: 4500000,
		},
		{
			Position:  "Angular Developer",
			Company:   "Enterprise Web",
			Location:  "Surabaya",
			SalaryMin: 3500000,
			SalaryMax: 5500000,
		},
		{
			Position:  "Svelte Developer",
			Company:   "Modern Web Studio",
			Location:  "Yogyakarta",
			SalaryMin: 2800000,
			SalaryMax: 4800000,
		},

		// Backend Development
		{
			Position:  "Backend Developer",
			Company:   "Digital Solutions",
			Location:  "Surabaya",
			SalaryMin: 3000000,
			SalaryMax: 5000000,
		},
		{
			Position:  "Python Developer",
			Company:   "CodeCraft Indonesia",
			Location:  "Surabaya",
			SalaryMin: 3500000,
			SalaryMax: 6000000,
		},
		{
			Position:  "Java Developer",
			Company:   "Enterprise Solutions",
			Location:  "Bandung",
			SalaryMin: 4000000,
			SalaryMax: 7000000,
		},
		{
			Position:  "Go Developer",
			Company:   "Microservices Inc",
			Location:  "Yogyakarta",
			SalaryMin: 4500000,
			SalaryMax: 7500000,
		},
		{
			Position:  "Node.js Developer",
			Company:   "JavaScript Pro",
			Location:  "Jakarta",
			SalaryMin: 3200000,
			SalaryMax: 5200000,
		},
		{
			Position:  "PHP Developer",
			Company:   "Web Solutions",
			Location:  "Bandung",
			SalaryMin: 2500000,
			SalaryMax: 4500000,
		},

		// Full Stack Development
		{
			Position:  "Full Stack Developer",
			Company:   "Innovation Labs",
			Location:  "Bandung",
			SalaryMin: 4000000,
			SalaryMax: 7000000,
		},
		{
			Position:  "MERN Stack Developer",
			Company:   "Modern Tech",
			Location:  "Jakarta",
			SalaryMin: 4500000,
			SalaryMax: 7500000,
		},
		{
			Position:  "Laravel Developer",
			Company:   "PHP Studio",
			Location:  "Surabaya",
			SalaryMin: 3000000,
			SalaryMax: 5500000,
		},

		// Mobile Development
		{
			Position:  "Mobile Developer",
			Company:   "AppWorks",
			Location:  "Bandung",
			SalaryMin: 3000000,
			SalaryMax: 5500000,
		},
		{
			Position:  "React Native Developer",
			Company:   "Cross Platform Apps",
			Location:  "Jakarta",
			SalaryMin: 3500000,
			SalaryMax: 6000000,
		},
		{
			Position:  "Flutter Developer",
			Company:   "Google Tech Indonesia",
			Location:  "Surabaya",
			SalaryMin: 3200000,
			SalaryMax: 5800000,
		},
		{
			Position:  "iOS Developer",
			Company:   "Apple Developer Studio",
			Location:  "Jakarta",
			SalaryMin: 4000000,
			SalaryMax: 7000000,
		},
		{
			Position:  "Android Developer",
			Company:   "Android Studio Pro",
			Location:  "Bandung",
			SalaryMin: 3800000,
			SalaryMax: 6500000,
		},

		// DevOps & Infrastructure
		{
			Position:  "DevOps Engineer",
			Company:   "CloudTech",
			Location:  "Yogyakarta",
			SalaryMin: 5000000,
			SalaryMax: 8000000,
		},
		{
			Position:  "Site Reliability Engineer",
			Company:   "Reliability First",
			Location:  "Jakarta",
			SalaryMin: 5500000,
			SalaryMax: 8500000,
		},
		{
			Position:  "Cloud Engineer",
			Company:   "AWS Solutions",
			Location:  "Surabaya",
			SalaryMin: 4500000,
			SalaryMax: 7500000,
		},
		{
			Position:  "Kubernetes Engineer",
			Company:   "Container Tech",
			Location:  "Bandung",
			SalaryMin: 4800000,
			SalaryMax: 7800000,
		},

		// Data & Analytics
		{
			Position:  "Data Scientist",
			Company:   "Analytics Pro",
			Location:  "Surabaya",
			SalaryMin: 4000000,
			SalaryMax: 6000000,
		},
		{
			Position:  "Data Engineer",
			Company:   "Big Data Solutions",
			Location:  "Jakarta",
			SalaryMin: 4500000,
			SalaryMax: 7000000,
		},
		{
			Position:  "Machine Learning Engineer",
			Company:   "AI Innovation",
			Location:  "Bandung",
			SalaryMin: 5000000,
			SalaryMax: 8000000,
		},
		{
			Position:  "Business Intelligence Analyst",
			Company:   "BI Solutions",
			Location:  "Yogyakarta",
			SalaryMin: 3500000,
			SalaryMax: 5500000,
		},

		// Design & UX
		{
			Position:  "UI/UX Designer",
			Company:   "Creative Studio",
			Location:  "Jakarta",
			SalaryMin: 2500000,
			SalaryMax: 4500000,
		},
		{
			Position:  "Product Designer",
			Company:   "Design Thinking",
			Location:  "Bandung",
			SalaryMin: 3000000,
			SalaryMax: 5000000,
		},
		{
			Position:  "Graphic Designer",
			Company:   "Visual Arts Studio",
			Location:  "Surabaya",
			SalaryMin: 2000000,
			SalaryMax: 3500000,
		},
		{
			Position:  "UX Researcher",
			Company:   "User Experience Lab",
			Location:  "Jakarta",
			SalaryMin: 2800000,
			SalaryMax: 4800000,
		},

		// Product & Project Management
		{
			Position:  "Product Manager",
			Company:   "StartupHub",
			Location:  "Bali",
			SalaryMin: 6000000,
			SalaryMax: 10000000,
		},
		{
			Position:  "Project Manager",
			Company:   "Project Solutions",
			Location:  "Jakarta",
			SalaryMin: 5000000,
			SalaryMax: 8000000,
		},
		{
			Position:  "Scrum Master",
			Company:   "Agile Solutions",
			Location:  "Bandung",
			SalaryMin: 4500000,
			SalaryMax: 7000000,
		},
		{
			Position:  "Product Owner",
			Company:   "Product First",
			Location:  "Surabaya",
			SalaryMin: 5500000,
			SalaryMax: 8500000,
		},

		// Quality Assurance
		{
			Position:  "QA Engineer",
			Company:   "Quality First",
			Location:  "Yogyakarta",
			SalaryMin: 2000000,
			SalaryMax: 3500000,
		},
		{
			Position:  "Test Automation Engineer",
			Company:   "Automation Pro",
			Location:  "Jakarta",
			SalaryMin: 3500000,
			SalaryMax: 5500000,
		},
		{
			Position:  "Performance Tester",
			Company:   "Performance Lab",
			Location:  "Bandung",
			SalaryMin: 3000000,
			SalaryMax: 5000000,
		},

		// System & Network
		{
			Position:  "System Administrator",
			Company:   "IT Solutions",
			Location:  "Bali",
			SalaryMin: 3500000,
			SalaryMax: 5500000,
		},
		{
			Position:  "Network Engineer",
			Company:   "Network Solutions",
			Location:  "Jakarta",
			SalaryMin: 4000000,
			SalaryMax: 6000000,
		},
		{
			Position:  "Security Engineer",
			Company:   "Cyber Security Pro",
			Location:  "Surabaya",
			SalaryMin: 4500000,
			SalaryMax: 7000000,
		},

		// Emerging Technologies
		{
			Position:  "Blockchain Developer",
			Company:   "Crypto Labs",
			Location:  "Bali",
			SalaryMin: 5000000,
			SalaryMax: 8000000,
		},
		{
			Position:  "IoT Developer",
			Company:   "Internet of Things",
			Location:  "Bandung",
			SalaryMin: 4000000,
			SalaryMax: 6500000,
		},
		{
			Position:  "AR/VR Developer",
			Company:   "Virtual Reality Studio",
			Location:  "Jakarta",
			SalaryMin: 4500000,
			SalaryMax: 7500000,
		},

		// Entry Level Positions
		{
			Position:  "Junior Developer",
			Company:   "Junior Tech",
			Location:  "Jakarta",
			SalaryMin: 1500000,
			SalaryMax: 2500000,
		},
		{
			Position:  "Graduate Developer",
			Company:   "Fresh Graduate Hub",
			Location:  "Bandung",
			SalaryMin: 1200000,
			SalaryMax: 2000000,
		},
		{
			Position:  "Intern Developer",
			Company:   "Internship Program",
			Location:  "Surabaya",
			SalaryMin: 800000,
			SalaryMax: 1500000,
		},

		// Senior & Lead Positions
		{
			Position:  "Senior Developer",
			Company:   "Senior Tech Solutions",
			Location:  "Jakarta",
			SalaryMin: 8000000,
			SalaryMax: 12000000,
		},
		{
			Position:  "Tech Lead",
			Company:   "Leadership Tech",
			Location:  "Bandung",
			SalaryMin: 10000000,
			SalaryMax: 15000000,
		},
		{
			Position:  "Engineering Manager",
			Company:   "Management Solutions",
			Location:  "Surabaya",
			SalaryMin: 12000000,
			SalaryMax: 18000000,
		},

		// Specialized Roles
		{
			Position:  "Game Developer",
			Company:   "Game Studio Indonesia",
			Location:  "Yogyakarta",
			SalaryMin: 3500000,
			SalaryMax: 6000000,
		},
		{
			Position:  "Embedded Systems Engineer",
			Company:   "Hardware Solutions",
			Location:  "Bandung",
			SalaryMin: 4000000,
			SalaryMax: 6500000,
		},
		{
			Position:  "Database Administrator",
			Company:   "Database Solutions",
			Location:  "Jakarta",
			SalaryMin: 4500000,
			SalaryMax: 7000000,
		},
		{
			Position:  "API Developer",
			Company:   "API Solutions",
			Location:  "Surabaya",
			SalaryMin: 3500000,
			SalaryMax: 5500000,
		},
	}

	// Insert sample jobs
	for _, job := range sampleJobs {
		err := models.CreateJob(&job)
		if err != nil {
			log.Printf("Error creating job %s: %v", job.Position, err)
		} else {
			log.Printf("Created job: %s at %s", job.Position, job.Company)
		}
	}

	log.Println("Sample data seeding completed!")
}

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Parse command line flags
	seedFlag := flag.Bool("seed", false, "Seed the database with sample data")
	flag.Parse()

	// Initialize database
	database.InitDB()

	// Seed data if flag is provided
	if *seedFlag {
		seedData()
		return
	}

	// Set Gin mode
	gin.SetMode(gin.ReleaseMode)

	// Create router
	r := gin.Default()

	// Add CORS middleware
	r.Use(middleware.CORSMiddleware())

	// API routes
	api := r.Group("/api")
	{
		// Jobs endpoints
		api.GET("/jobs", handlers.GetJobs)
		api.GET("/jobs/:id", handlers.GetJobByID)
		api.POST("/jobs", handlers.CreateJob)
		api.GET("/locations", handlers.GetLocations)

		// Applications endpoints
		api.POST("/applications", handlers.CreateApplication)
		api.GET("/applications", handlers.GetApplications)
		api.GET("/applications/:id", handlers.GetApplicationByID)
	}

	// Serve static files from uploads directory
	r.Static("/uploads", "./uploads")

	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Get port from .env file
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is required in .env file")
	}

	log.Printf("Server starting on port %s...", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
