package handlers

import (
	"evconn/internal/core/domain/models"
	"evconn/internal/core/ports"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FeedbackHandler struct {
	feedbackService ports.FeedbackService
}

func NewFeedbackHandler(feedbackService ports.FeedbackService) *FeedbackHandler {
	return &FeedbackHandler{
		feedbackService: feedbackService,
	}
}

func (h *FeedbackHandler) Create(c *gin.Context) {
	var feedback models.Feedback
	if err := c.ShouldBindJSON(&feedback); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Get current user from context
	user, _ := c.Get("user")
	currentUser := user.(*models.User)
	feedback.UserID = currentUser.ID

	if err := h.feedbackService.CreateFeedback(c.Request.Context(), &feedback); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, feedback)
}

func (h *FeedbackHandler) GetByModule(c *gin.Context) {
	moduleId, err := strconv.ParseUint(c.Param("moduleId"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid module ID"})
		return
	}

	feedbackType := c.Query("type")
	var feedbacks []models.Feedback

	if feedbackType != "" {
		feedbacks, err = h.feedbackService.GetFeedbackByModuleAndType(c.Request.Context(), uint(moduleId), feedbackType)
	} else {
		feedbacks, err = h.feedbackService.GetFeedbackByModule(c.Request.Context(), uint(moduleId))
	}

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, feedbacks)
}
