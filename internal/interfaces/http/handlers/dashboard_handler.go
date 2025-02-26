package handlers

import (
	"evconn/internal/core/ports"
	"evconn/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	userService       ports.UserService
	courseService     ports.CourseService
	submissionService ports.SubmissionService
}

func NewDashboardHandler(
	userService ports.UserService,
	courseService ports.CourseService,
	submissionService ports.SubmissionService,
) *DashboardHandler {
	return &DashboardHandler{
		userService:       userService,
		courseService:     courseService,
		submissionService: submissionService,
	}
}

func (h *DashboardHandler) GetStatistics(c *gin.Context) {
	ctx := c.Request.Context()

	// Get user statistics
	users, err := h.userService.GetAllUsers(ctx)
	if err != nil {
		response.Error(c, 500, "Failed to get users: "+err.Error())
		return
	}

	// Count users by role
	var adminCount, mentorCount, studentCount int
	for _, user := range users {
		switch user.Role {
		case "admin":
			adminCount++
		case "mentor":
			mentorCount++
		case "student":
			studentCount++
		}
	}

	// Get courses
	courses, err := h.courseService.GetAllCourses(ctx)
	if err != nil {
		response.Error(c, 500, "Failed to get courses: "+err.Error())
		return
	}

	// Create statistics response
	stats := map[string]interface{}{
		"users": map[string]int{
			"total":   len(users),
			"admin":   adminCount,
			"mentor":  mentorCount,
			"student": studentCount,
		},
		"courses": len(courses),
	}

	response.Success(c, stats)
}
