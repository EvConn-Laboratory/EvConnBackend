package http

import (
	"evconn/internal/interfaces/http/handlers"
	"evconn/internal/interfaces/http/middleware"
)

func (s *Server) SetupRoutes(
	authHandler *handlers.AuthHandler,
	courseHandler *handlers.CourseHandler,
	moduleHandler *handlers.ModuleHandler,
	assignmentHandler *handlers.AssignmentHandler,
	submissionHandler *handlers.SubmissionHandler,
	feedbackHandler *handlers.FeedbackHandler,
	labHandler *handlers.LabHandler,
	dashboardHandler *handlers.DashboardHandler,
	userHandler *handlers.UserHandler,
	enrollmentHandler *handlers.EnrollmentHandler, // Add this parameter
	fileHandler *handlers.FileHandler, // Add this parameter
) {
	// Public routes
	auth := s.router.Group("/api/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	// Protected routes
	api := s.router.Group("/api")
	api.Use(middleware.AuthMiddleware(authHandler.GetAuthService()))
	{
		// Add the /me endpoint in protected routes
		auth := api.Group("/auth")
		{
			auth.GET("/me", authHandler.GetMe)
		}

		// Course routes
		courses := api.Group("/courses")
		{
			courses.POST("/", courseHandler.Create)
			courses.GET("/", courseHandler.GetAll)
			courses.GET("/:id", courseHandler.GetByID)
		}

		// Module routes
		modules := api.Group("/modules")
		{
			modules.POST("", moduleHandler.Create)
			modules.GET("/:moduleId", moduleHandler.GetByID)
			modules.POST("/:moduleId/mentors/:mentorId", moduleHandler.AssignMentor)
			modules.POST("/:moduleId/feedback", feedbackHandler.Create)
			modules.GET("/:moduleId/feedback", feedbackHandler.GetByModule)
		}

		// Assignment routes
		assignments := api.Group("/assignments")
		{
			assignments.GET("/pending", assignmentHandler.GetPending)
			assignments.GET("/module/:moduleId", assignmentHandler.GetByModule)
			assignments.POST("", middleware.MentorOrAdmin(), assignmentHandler.Create)
			assignments.PUT("/:id", middleware.MentorOrAdmin(), assignmentHandler.Update)
			assignments.DELETE("/:id", middleware.MentorOrAdmin(), assignmentHandler.Delete)
		}

		// Submission routes
		submissions := api.Group("/submissions")
		{
			submissions.POST("/", submissionHandler.Create)
			submissions.POST("/:id/grade", submissionHandler.Grade)
			submissions.GET("/user", submissionHandler.GetByUser)
		}

		// Lab routes
		lab := api.Group("/labs")
		{
			lab.GET("", labHandler.GetAll)
			lab.GET("/:id", labHandler.GetByID)
			lab.POST("", middleware.AdminOnly(), labHandler.Create)
			lab.PUT("/:id", middleware.AdminOnly(), labHandler.Update)
			lab.DELETE("/:id", middleware.AdminOnly(), labHandler.Delete)
		}

		// Enrollment routes
		enrollments := api.Group("/enrollments")
		{
			enrollments.POST("/courses/:courseId", enrollmentHandler.Enroll)
			enrollments.GET("/my", enrollmentHandler.GetMy)
			enrollments.GET("/courses/:courseId", middleware.MentorOrAdmin(), enrollmentHandler.GetByCourse)
			enrollments.PUT("/:id/status", middleware.MentorOrAdmin(), enrollmentHandler.UpdateStatus)
			enrollments.DELETE("/:id", enrollmentHandler.Cancel)
		}

		// File routes
		files := api.Group("/files")
		{
			files.POST("/upload/:entityType/:entityId", fileHandler.Upload)
			files.GET("/:id", fileHandler.GetByID)
			files.GET("/entity/:entityType/:entityId", fileHandler.GetByEntity)
			files.DELETE("/:id", fileHandler.Delete)
		}

		// Admin dashboard routes
		admin := api.Group("/admin")
		admin.Use(middleware.AdminOnly())
		{
			admin.GET("/statistics", dashboardHandler.GetStatistics)
			admin.GET("/users", userHandler.GetAll)
			admin.POST("/users/import", userHandler.Import)
		}
	}
}
