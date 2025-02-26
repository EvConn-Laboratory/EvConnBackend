package main

import (
	"context"
	"evconn/internal/core/services"
	"evconn/internal/infrastructure/auth"
	"evconn/internal/infrastructure/database"
	"evconn/internal/infrastructure/persistence/repositories"
	"evconn/internal/infrastructure/storage"
	"evconn/internal/interfaces/http"
	"evconn/internal/interfaces/http/handlers"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Initialize DB
	db, err := database.NewMySQLConnection(&database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	courseRepo := repositories.NewCourseRepository(db)
	moduleRepo := repositories.NewModuleRepository(db)
	moduleMentorRepo := repositories.NewModuleMentorRepository(db)
	assignmentRepo := repositories.NewAssignmentRepository(db)
	submissionRepo := repositories.NewSubmissionRepository(db)
	feedbackRepo := repositories.NewFeedbackRepository(db)
	labRepo := repositories.NewLabRepository(db)
	fileRepo := repositories.NewFileRepository(db)
	enrollmentRepo := repositories.NewEnrollmentRepository(db)

	// Initialize storage
	storageConfig := &storage.MinioConfig{
		Endpoint:  os.Getenv("MINIO_ENDPOINT"),
		AccessKey: os.Getenv("MINIO_ACCESS_KEY"),
		SecretKey: os.Getenv("MINIO_SECRET_KEY"),
		UseSSL:    os.Getenv("MINIO_USE_SSL") == "true",
	}
	storageService, err := storage.NewMinioStorage(storageConfig)
	if err != nil {
		log.Fatalf("Failed to initialize storage service: %v", err)
	}

	// Initialize JWT auth
	jwtAuth := auth.NewJWTAuth(os.Getenv("JWT_SECRET"), 24*time.Hour)

	// Initialize services
	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(userRepo, jwtAuth)
	courseService := services.NewCourseService(courseRepo)
	moduleService := services.NewModuleService(moduleRepo, moduleMentorRepo)
	assignmentService := services.NewAssignmentService(assignmentRepo, submissionRepo)
	submissionService := services.NewSubmissionService(submissionRepo, storageService)
	feedbackService := services.NewFeedbackService(feedbackRepo)
	labService := services.NewLabService(labRepo)
	fileService := services.NewFileService(fileRepo, storageService)
	enrollmentService := services.NewEnrollmentService(enrollmentRepo, courseRepo, userRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	courseHandler := handlers.NewCourseHandler(courseService)
	moduleHandler := handlers.NewModuleHandler(moduleService)
	assignmentHandler := handlers.NewAssignmentHandler(assignmentService)
	submissionHandler := handlers.NewSubmissionHandler(submissionService)
	feedbackHandler := handlers.NewFeedbackHandler(feedbackService)
	labHandler := handlers.NewLabHandler(labService)
	userHandler := handlers.NewUserHandler(userService)
	dashboardHandler := handlers.NewDashboardHandler(userService, courseService, submissionService)
	enrollmentHandler := handlers.NewEnrollmentHandler(enrollmentService)
	fileHandler := handlers.NewFileHandler(fileService)

	// Initialize and setup server
	server := http.NewServer(os.Getenv("PORT"))
	server.SetupMiddleware()
	server.SetupRoutes(
		authHandler,
		courseHandler,
		moduleHandler,
		assignmentHandler,
		submissionHandler,
		feedbackHandler,
		labHandler,
		dashboardHandler,
		userHandler,
		enrollmentHandler,
		fileHandler,
	)

	// Start server
	go func() {
		if err := server.Start(); err != nil {
			log.Printf("Server error: %v\n", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
}
