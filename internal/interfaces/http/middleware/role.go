package middleware

import (
	"evconn/internal/core/domain/models"

	"github.com/gin-gonic/gin"
)

// MentorOrAdmin ensures the user is either a mentor or admin
func MentorOrAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
			return
		}

		currentUser, ok := user.(*models.User)
		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid user object"})
			return
		}

		if currentUser.Role != "mentor" && currentUser.Role != "admin" {
			c.AbortWithStatusJSON(403, gin.H{"error": "requires mentor or admin role"})
			return
		}

		c.Next()
	}
}
