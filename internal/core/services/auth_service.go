package services

import (
	"context"
	"evconn/internal/core/domain/models"
	"evconn/internal/core/ports"
	"evconn/internal/pkg/errors"

	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	*BaseService
	userRepo ports.UserRepository
	jwtAuth  ports.JWTAuth
}

// Ensure authService implements AuthService interface
var _ ports.AuthService = (*authService)(nil)

func NewAuthService(userRepo ports.UserRepository, jwtAuth ports.JWTAuth) ports.AuthService {
	return &authService{
		BaseService: NewBaseService(),
		userRepo:    userRepo,
		jwtAuth:     jwtAuth,
	}
}

func (s *authService) Login(ctx context.Context, nim, password string) (string, error) {
	user, err := s.userRepo.FindByNIM(ctx, nim)
	if err != nil {
		return "", errors.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.ErrInvalidCredentials
	}

	return s.jwtAuth.GenerateToken(user.ID, user.Role)
}

func (s *authService) LoginWithUser(ctx context.Context, nim, password string) (string, *models.User, error) {
	user, err := s.userRepo.FindByNIM(ctx, nim)
	if err != nil {
		return "", nil, errors.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", nil, errors.ErrInvalidCredentials
	}

	token, err := s.jwtAuth.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", nil, errors.ErrInternal
	}

	// Create a copy of user without the password
	userWithoutPassword := *user
	userWithoutPassword.Password = ""

	return token, &userWithoutPassword, nil
}

func (s *authService) Register(ctx context.Context, user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.ErrInternal
	}

	user.Password = string(hashedPassword)
	return s.userRepo.Create(ctx, user)
}

func (s *authService) ValidateToken(token string) (*models.User, error) {
	return s.jwtAuth.ValidateToken(token)
}

func (s *authService) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.ErrUserNotFound
	}

	// Don't send password hash
	user.Password = ""
	return user, nil
}
