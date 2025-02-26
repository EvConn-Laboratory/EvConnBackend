package repositories

import (
	"context"
	"evconn/internal/core/domain/models"
	"evconn/internal/core/ports"

	"gorm.io/gorm"
)

type assignmentRepository struct {
	db *gorm.DB
}

func NewAssignmentRepository(db *gorm.DB) ports.AssignmentRepository {
	return &assignmentRepository{
		db: db,
	}
}

func (r *assignmentRepository) DB(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx)
}

func (r *assignmentRepository) Create(ctx context.Context, assignment *models.Assignment) error {
	return r.DB(ctx).Create(assignment).Error
}

func (r *assignmentRepository) FindByID(ctx context.Context, id uint) (*models.Assignment, error) {
	var assignment models.Assignment
	err := r.DB(ctx).First(&assignment, id).Error
	if err != nil {
		return nil, err
	}
	return &assignment, nil
}

func (r *assignmentRepository) FindByModuleID(ctx context.Context, moduleID uint) ([]models.Assignment, error) {
	var assignments []models.Assignment
	err := r.DB(ctx).Where("module_id = ?", moduleID).Find(&assignments).Error
	return assignments, err
}

func (r *assignmentRepository) Update(ctx context.Context, assignment *models.Assignment) error {
	return r.DB(ctx).Save(assignment).Error
}

func (r *assignmentRepository) Delete(ctx context.Context, id uint) error {
	return r.DB(ctx).Delete(&models.Assignment{}, id).Error
}

func (r *assignmentRepository) GetAll(ctx context.Context) ([]models.Assignment, error) {
	var assignments []models.Assignment
	err := r.DB(ctx).Find(&assignments).Error
	return assignments, err
}
