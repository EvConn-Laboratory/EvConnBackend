package handlers

import (
	"evconn/internal/core/domain/models"
	"evconn/internal/core/ports"
	"evconn/internal/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	courseService ports.CourseService
}

func NewCourseHandler(courseService ports.CourseService) *CourseHandler {
	return &CourseHandler{
		courseService: courseService,
	}
}

func (h *CourseHandler) Create(c *gin.Context) {
	var course models.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Get the current user from context
	user, exists := c.Get("user")
	if !exists {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	// Verify user is a mentor
	currentUser := user.(*models.User)
	if currentUser.Role != "mentor" {
		c.JSON(403, gin.H{"error": "only mentors can create courses"})
		return
	}

	// Set the mentor ID from the authenticated user
	course.MentorID = currentUser.ID

	if err := h.courseService.CreateCourse(c.Request.Context(), &course); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, course)
}

func (h *CourseHandler) GetByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	course, err := h.courseService.GetCourse(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "course not found"})
		return
	}

	c.JSON(200, course)
}

func (h *CourseHandler) GetAll(c *gin.Context) {
	courses, err := h.courseService.GetAllCourses(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, courses)
}

// Add to your handlers
func (h *CourseHandler) Search(c *gin.Context) {
	query := c.Query("q")
	courses, err := h.courseService.SearchCourses(c.Request.Context(), query)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}
	response.Success(c, courses)
}
