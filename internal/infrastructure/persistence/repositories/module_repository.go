package repositories

import (
	"context"
	"evconn/internal/core/domain/models"
	"evconn/internal/core/ports"

	"gorm.io/gorm"
)

type moduleRepository struct {
	*BaseRepository[models.Module]
}

func NewModuleRepository(db *gorm.DB) ports.ModuleRepository {
	return &moduleRepository{
		BaseRepository: NewBaseRepository[models.Module](db),
	}
}

func (r *moduleRepository) Create(ctx context.Context, module *models.Module) error {
	return r.DB(ctx).Create(module).Error
}

func (r *moduleRepository) FindByID(ctx context.Context, id uint) (*models.Module, error) {
	var module models.Module
	err := r.DB(ctx).Preload("Course").Preload("Assignments").First(&module, id).Error
	return &module, err
}

func (r *moduleRepository) FindByCourseID(ctx context.Context, courseID uint) ([]models.Module, error) {
	var modules []models.Module
	err := r.DB(ctx).Where("course_id = ?", courseID).Find(&modules).Error
	return modules, err
}

func (r *moduleRepository) Update(ctx context.Context, module *models.Module) error {
	return r.DB(ctx).Save(module).Error
}

func (r *moduleRepository) Delete(ctx context.Context, id uint) error {
	return r.DB(ctx).Delete(&models.Module{}, id).Error
}
