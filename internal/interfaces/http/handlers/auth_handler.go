package handlers

import (
	"context"
	"evconn/internal/core/domain/models"
	"evconn/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService ports.AuthService
}

type AuthService interface {
	Login(ctx context.Context, nim, password string) (string, error)
	LoginWithUser(ctx context.Context, nim, password string) (string, *models.User, error)
	// ...other methods
}

func NewAuthHandler(authService ports.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// GetAuthService returns the auth service for middleware use
func (h *AuthHandler) GetAuthService() ports.AuthService {
	return h.authService
}

func (h *AuthHandler) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.authService.Register(c.Request.Context(), &user); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "user registered successfully"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		NIM      string `json:"nim"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Get token and user data
	token, user, err := h.authService.LoginWithUser(c.Request.Context(), req.NIM, req.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	// Return both token and user
	c.JSON(200, gin.H{
		"token": token,
		"user":  user,
	})
}

func (h *AuthHandler) GetMe(c *gin.Context) {
	// Get user from context (set by auth middleware)
	user, exists := c.Get("user")
	if !exists {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	currentUser := user.(*models.User)

	// Get full user details
	userDetails, err := h.authService.GetUserByID(c.Request.Context(), currentUser.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, userDetails)
}
