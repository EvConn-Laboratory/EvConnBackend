package repositories

import (
	"context"
	"evconn/internal/core/domain/models"
	"evconn/internal/core/ports"

	"gorm.io/gorm"
)

type fileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) ports.FileRepository {
	return &fileRepository{db: db}
}

func (r *fileRepository) Create(ctx context.Context, file *models.File) error {
	return r.db.WithContext(ctx).Create(file).Error
}

func (r *fileRepository) FindByID(ctx context.Context, id uint) (*models.File, error) {
	var file models.File
	err := r.db.WithContext(ctx).First(&file, id).Error
	return &file, err
}

func (r *fileRepository) FindByEntityTypeAndID(ctx context.Context, entityType string, entityID uint) ([]*models.File, error) {
	var files []*models.File
	err := r.db.WithContext(ctx).
		Where("entity_type = ? AND entity_id = ?", entityType, entityID).
		Find(&files).Error
	return files, err
}

func (r *fileRepository) FindByUserID(ctx context.Context, userID uint) ([]*models.File, error) {
	var files []*models.File
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Find(&files).Error
	return files, err
}

func (r *fileRepository) Update(ctx context.Context, file *models.File) error {
	return r.db.WithContext(ctx).Save(file).Error
}

func (r *fileRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.File{}, id).Error
}
