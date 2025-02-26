package services

import (
	"context"
	"evconn/internal/core/domain/models"
	"evconn/internal/core/ports"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	*BaseService
	userRepo ports.UserRepository
}

func NewUserService(userRepo ports.UserRepository) ports.UserService {
	return &userService{
		BaseService: NewBaseService(),
		userRepo:    userRepo,
	}
}

func (s *userService) Create(ctx context.Context, user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.userRepo.Create(ctx, user)
}

func (s *userService) GetByID(ctx context.Context, id uint) (*models.User, error) {
	return s.userRepo.FindByID(ctx, id)
}

func (s *userService) GetByNIM(ctx context.Context, nim string) (*models.User, error) {
	return s.userRepo.FindByNIM(ctx, nim)
}

func (s *userService) GetByLab(ctx context.Context, lab string) ([]*models.User, error) {
	return s.userRepo.FindByLab(ctx, lab)
}

func (s *userService) ImportUsers(ctx context.Context, users []*models.User) ([]string, error) {
	var errors []string

	for _, user := range users {
		// Hash password before storing
		hashedPassword, err := s.hashPassword(user.Password)
		if err != nil {
			errors = append(errors, "Failed to hash password for user "+user.NIM+": "+err.Error())
			continue
		}

		user.Password = hashedPassword

		// Create user
		err = s.userRepo.Create(ctx, user)
		if err != nil {
			errors = append(errors, "Failed to create user "+user.NIM+": "+err.Error())
			continue
		}
	}

	return errors, nil
}

func (s *userService) hashPassword(password string) (string, error) {
	// Implementation depends on your password hashing strategy
	// Typically use bcrypt
	return password, nil // Replace with actual implementation
}

func (s *userService) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	return s.userRepo.FindAll(ctx)
}

func (s *userService) GetAllUsersPaginated(ctx context.Context, page, pageSize int) ([]*models.User, int64, error) {
	offset := (page - 1) * pageSize
	return s.userRepo.FindAllPaginated(ctx, offset, pageSize)
}
