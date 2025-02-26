package ports

import (
	"context"
	"evconn/internal/core/domain/models"
)

type JWTAuth interface {
	GenerateToken(userID uint, role string) (string, error)
	ValidateToken(token string) (*models.User, error)
}

type AuthService interface {
	Login(ctx context.Context, nim, password string) (string, error)
	LoginWithUser(ctx context.Context, nim, password string) (string, *models.User, error)
	Register(ctx context.Context, user *models.User) error
	ValidateToken(token string) (*models.User, error)
	GetUserByID(ctx context.Context, id uint) (*models.User, error)
}
