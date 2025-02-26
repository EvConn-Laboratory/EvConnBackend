package repositories

import (
	"context"
	"evconn/internal/core/domain/models"
	"evconn/internal/core/ports"

	"gorm.io/gorm"
)

type moduleMentorRepository struct {
	*BaseRepository[models.ModuleMentor]
}

func NewModuleMentorRepository(db *gorm.DB) ports.ModuleMentorRepository {
	return &moduleMentorRepository{
		BaseRepository: NewBaseRepository[models.ModuleMentor](db),
	}
}

func (r *moduleMentorRepository) Create(ctx context.Context, moduleMentor *models.ModuleMentor) error {
	return r.DB(ctx).Create(moduleMentor).Error
}

func (r *moduleMentorRepository) FindByID(ctx context.Context, id uint) (*models.ModuleMentor, error) {
	var moduleMentor models.ModuleMentor
	err := r.DB(ctx).Preload("Module").Preload("Mentor").First(&moduleMentor, id).Error
	return &moduleMentor, err
}

func (r *moduleMentorRepository) FindByModuleID(ctx context.Context, moduleID uint) ([]models.ModuleMentor, error) {
	var moduleMentors []models.ModuleMentor
	err := r.DB(ctx).Where("module_id = ?", moduleID).Preload("Mentor").Find(&moduleMentors).Error
	return moduleMentors, err
}

func (r *moduleMentorRepository) FindByMentorID(ctx context.Context, mentorID uint) ([]models.ModuleMentor, error) {
	var moduleMentors []models.ModuleMentor
	err := r.DB(ctx).Where("mentor_id = ?", mentorID).Preload("Module").Find(&moduleMentors).Error
	return moduleMentors, err
}

func (r *moduleMentorRepository) Delete(ctx context.Context, id uint) error {
	return r.DB(ctx).Delete(&models.ModuleMentor{}, id).Error
}
