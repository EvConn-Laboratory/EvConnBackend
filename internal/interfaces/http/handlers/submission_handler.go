package handlers

import (
	"evconn/internal/core/domain/models"
	"evconn/internal/core/ports"
	"evconn/internal/pkg/response"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type SubmissionHandler struct {
	submissionService ports.SubmissionService
}

func NewSubmissionHandler(submissionService ports.SubmissionService) *SubmissionHandler {
	return &SubmissionHandler{
		submissionService: submissionService,
	}
}

// Create handles submission creation with file upload
func (h *SubmissionHandler) Create(c *gin.Context) {
	// Get authenticated user
	user, exists := c.Get("user")
	if !exists {
		response.Error(c, 401, "Unauthorized access")
		return
	}
	currentUser := user.(*models.User)

	// Parse form data
	assignmentIDStr := c.PostForm("assignmentID")
	content := c.PostForm("content")

	if assignmentIDStr == "" {
		response.Error(c, 400, "Assignment ID is required")
		return
	}

	assignmentID := response.StringToUint(assignmentIDStr)

	// Parse file
	file, err := c.FormFile("file")
	if err != nil {
		response.Error(c, 400, fmt.Sprintf("File upload error: %v", err))
		return
	}

	// Create submission model
	submission := &models.Submission{
		StudentID:    currentUser.ID, // Use StudentID instead of UserID
		AssignmentID: assignmentID,
		Content:      content,
		Status:       "pending",
		SubmittedAt:  time.Now().Format("2006-01-02 15:04:05"),
	}

	// Open file
	fileReader, err := file.Open()
	if err != nil {
		response.Error(c, 500, fmt.Sprintf("Failed to open file: %v", err))
		return
	}
	defer fileReader.Close()

	// Create the submission without file handling in this method
	// The SubmissionService's CreateSubmission only accepts context and submission
	err = h.submissionService.CreateSubmission(c.Request.Context(), submission)
	if err != nil {
		response.Error(c, 500, fmt.Sprintf("Failed to create submission: %v", err))
		return
	}

	response.Success(c, submission)
}

func (h *SubmissionHandler) Grade(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var req struct {
		Score float64 `json:"score" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.submissionService.GradeSubmission(c.Request.Context(), uint(id), req.Score)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "submission graded successfully"})
}

// Fix the GetByUser method
func (h *SubmissionHandler) GetByUser(c *gin.Context) {
	// Get user from context (set by auth middleware)
	user, exists := c.Get("user")
	if !exists {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	userID := user.(*models.User).ID

	submissions, err := h.submissionService.GetSubmissionsByUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, submissions)
}
