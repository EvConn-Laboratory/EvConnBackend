package handlers

import (
	"evconn/internal/core/domain/models"
	"evconn/internal/core/ports"
	"evconn/internal/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AssignmentHandler struct {
	assignmentService ports.AssignmentService
}

func NewAssignmentHandler(assignmentService ports.AssignmentService) *AssignmentHandler {
	return &AssignmentHandler{
		assignmentService: assignmentService,
	}
}

func (h *AssignmentHandler) GetPending(c *gin.Context) {
	// Get user from context (set by auth middleware)
	user, exists := c.Get("user")
	if !exists {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	// Update this to use the correct method name
	assignments, err := h.assignmentService.GetPendingAssignments(c.Request.Context(), user.(*models.User).ID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, assignments)
}

func (h *AssignmentHandler) GetByModule(c *gin.Context) {
	moduleID, err := strconv.ParseUint(c.Param("moduleId"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid module id"})
		return
	}

	// Update this to use the correct method name
	assignments, err := h.assignmentService.GetAssignmentsByModule(c.Request.Context(), uint(moduleID))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, assignments)
}

func (h *AssignmentHandler) Create(c *gin.Context) {
	var assignment models.Assignment
	if err := c.ShouldBindJSON(&assignment); err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	// Check if user is authenticated and has proper role
	user, exists := c.Get("user")
	if !exists {
		response.Error(c, 401, "Unauthorized")
		return
	}

	currentUser := user.(*models.User)
	if currentUser.Role != "admin" && currentUser.Role != "mentor" {
		response.Error(c, 403, "Only admins and mentors can create assignments")
		return
	}

	err := h.assignmentService.CreateAssignment(c.Request.Context(), &assignment)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, assignment)
}

func (h *AssignmentHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, 400, "Invalid ID")
		return
	}

	var assignment models.Assignment
	if err := c.ShouldBindJSON(&assignment); err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	assignment.ID = uint(id)
	err = h.assignmentService.UpdateAssignment(c.Request.Context(), &assignment)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, assignment)
}

func (h *AssignmentHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, 400, "Invalid ID")
		return
	}

	err = h.assignmentService.DeleteAssignment(c.Request.Context(), uint(id))
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "Assignment deleted successfully"})
}
