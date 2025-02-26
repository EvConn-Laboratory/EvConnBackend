package repositories

import (
	"context"
	"evconn/internal/core/domain/models"
	"evconn/internal/core/ports"

	"gorm.io/gorm"
)

type userRepository struct {
	*BaseRepository[models.User]
}

func NewUserRepository(db *gorm.DB) ports.UserRepository {
	return &userRepository{
		BaseRepository: NewBaseRepository[models.User](db),
	}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	return r.DB(ctx).Create(user).Error
}

func (r *userRepository) FindByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	err := r.DB(ctx).First(&user, id).Error
	return &user, err
}

func (r *userRepository) FindByNIM(ctx context.Context, nim string) (*models.User, error) {
	var user models.User
	err := r.DB(ctx).Where("nim = ?", nim).First(&user).Error
	return &user, err
}

func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	return r.DB(ctx).Save(user).Error
}

func (r *userRepository) Delete(ctx context.Context, id uint) error {
	return r.DB(ctx).Delete(&models.User{}, id).Error
}

func (r *userRepository) CreateBatch(ctx context.Context, users []*models.User) error {
	// Use transaction to ensure all users are created or none
	return r.DB(ctx).Transaction(func(tx *gorm.DB) error {
		for _, user := range users {
			if err := tx.Create(user).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *userRepository) FindByLab(ctx context.Context, lab string) ([]*models.User, error) {
	var users []*models.User
	err := r.DB(ctx).Where("lab = ?", lab).Find(&users).Error
	return users, err
}

func (r *userRepository) FindAll(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	err := r.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.DB(ctx).Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepository) FindAllPaginated(ctx context.Context, offset, limit int) ([]*models.User, int64, error) {
	var users []*models.User
	var total int64

	err := r.db.WithContext(ctx).Model(&models.User{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
