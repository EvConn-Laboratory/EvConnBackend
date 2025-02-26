package middleware

import (
	"github.com/gin-gonic/gin"
)

// AdminOnly middleware ensures the user is an admin
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
			return
		}

		userObj, ok := user.(interface {
			GetRole() string
		})

		if !ok || userObj.GetRole() != "admin" {
			c.AbortWithStatusJSON(403, gin.H{"error": "admin access required"})
			return
		}

		c.Next()
	}
}
