package handlers

import (
	"evconn/internal/core/domain/models"
	"evconn/internal/core/ports"
	"evconn/internal/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EnrollmentHandler struct {
	enrollmentService ports.EnrollmentService
}

func NewEnrollmentHandler(enrollmentService ports.EnrollmentService) *EnrollmentHandler {
	return &EnrollmentHandler{
		enrollmentService: enrollmentService,
	}
}

func (h *EnrollmentHandler) Enroll(c *gin.Context) {
	// Get authenticated user
	user, exists := c.Get("user")
	if !exists {
		response.Error(c, 401, "Unauthorized")
		return
	}

	currentUser := user.(*models.User)

	// Get course ID from path
	courseID, err := strconv.ParseUint(c.Param("courseId"), 10, 32)
	if err != nil {
		response.Error(c, 400, "Invalid course ID")
		return
	}

	// Enroll in course
	enrollment, err := h.enrollmentService.EnrollInCourse(c.Request.Context(), currentUser.ID, uint(courseID))
	if err != nil {
		if err.Error() == "already enrolled in this course" {
			response.Error(c, 400, err.Error())
		} else {
			response.Error(c, 500, "Failed to enroll: "+err.Error())
		}
		return
	}

	response.Success(c, enrollment)
}

func (h *EnrollmentHandler) GetMy(c *gin.Context) {
	// Get authenticated user
	user, exists := c.Get("user")
	if !exists {
		response.Error(c, 401, "Unauthorized")
		return
	}

	currentUser := user.(*models.User)

	// Get enrollments
	enrollments, err := h.enrollmentService.GetEnrollmentsByStudent(c.Request.Context(), currentUser.ID)
	if err != nil {
		response.Error(c, 500, "Failed to get enrollments: "+err.Error())
		return
	}

	response.Success(c, enrollments)
}

func (h *EnrollmentHandler) GetByCourse(c *gin.Context) {
	// Get course ID from path
	courseID, err := strconv.ParseUint(c.Param("courseId"), 10, 32)
	if err != nil {
		response.Error(c, 400, "Invalid course ID")
		return
	}

	// Get enrollments
	enrollments, err := h.enrollmentService.GetEnrollmentsByCourse(c.Request.Context(), uint(courseID))
	if err != nil {
		response.Error(c, 500, "Failed to get enrollments: "+err.Error())
		return
	}

	response.Success(c, enrollments)
}

func (h *EnrollmentHandler) UpdateStatus(c *gin.Context) {
	// Get enrollment ID from path
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, 400, "Invalid enrollment ID")
		return
	}

	// Bind request body
	var req struct {
		Status string `json:"status" binding:"required"`
		Notes  string `json:"notes"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request: "+err.Error())
		return
	}

	// Update status
	err = h.enrollmentService.UpdateEnrollmentStatus(c.Request.Context(), uint(id), req.Status, req.Notes)
	if err != nil {
		response.Error(c, 500, "Failed to update status: "+err.Error())
		return
	}

	response.Success(c, gin.H{"message": "Enrollment status updated"})
}

func (h *EnrollmentHandler) Cancel(c *gin.Context) {
	// Get authenticated user
	user, exists := c.Get("user")
	if !exists {
		response.Error(c, 401, "Unauthorized")
		return
	}

	currentUser := user.(*models.User)

	// Get enrollment ID from path
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, 400, "Invalid enrollment ID")
		return
	}

	// Verify enrollment belongs to current user
	enrollment, err := h.enrollmentService.GetEnrollment(c.Request.Context(), uint(id))
	if err != nil {
		response.Error(c, 404, "Enrollment not found")
		return
	}

	if enrollment.StudentID != currentUser.ID && currentUser.Role != "admin" {
		response.Error(c, 403, "Not authorized to cancel this enrollment")
		return
	}

	// Cancel enrollment
	err = h.enrollmentService.CancelEnrollment(c.Request.Context(), uint(id))
	if err != nil {
		response.Error(c, 500, "Failed to cancel enrollment: "+err.Error())
		return
	}

	response.Success(c, gin.H{"message": "Enrollment cancelled"})
}
