package repositories

import (
	"context"
	"evconn/internal/core/domain/models"

	"gorm.io/gorm"
)

type labRepository struct {
	db *gorm.DB
}

func NewLabRepository(db *gorm.DB) *labRepository {
	return &labRepository{db: db}
}

func (r *labRepository) Create(ctx context.Context, lab *models.Lab) error {
	return r.db.WithContext(ctx).Create(lab).Error
}

func (r *labRepository) FindAll(ctx context.Context) ([]*models.Lab, error) {
	var labs []*models.Lab
	err := r.db.WithContext(ctx).Find(&labs).Error
	if err != nil {
		return nil, err
	}
	return labs, nil
}

func (r *labRepository) FindByID(ctx context.Context, id uint) (*models.Lab, error) {
	var lab models.Lab
	err := r.db.WithContext(ctx).First(&lab, id).Error
	if err != nil {
		return nil, err
	}
	return &lab, nil
}

func (r *labRepository) Update(ctx context.Context, lab *models.Lab) error {
	return r.db.WithContext(ctx).Save(lab).Error
}

func (r *labRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Lab{}, id).Error
}
