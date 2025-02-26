package repositories

import (
	"context"
	"evconn/internal/core/domain/models"
	"evconn/internal/core/ports"

	"gorm.io/gorm"
)

type submissionRepository struct {
	*BaseRepository[models.Submission]
}

func NewSubmissionRepository(db *gorm.DB) ports.SubmissionRepository {
	return &submissionRepository{
		BaseRepository: NewBaseRepository[models.Submission](db),
	}
}

func (r *submissionRepository) Create(ctx context.Context, submission *models.Submission) error {
	return r.DB(ctx).Create(submission).Error
}

func (r *submissionRepository) FindByID(ctx context.Context, id uint) (*models.Submission, error) {
	var submission models.Submission
	err := r.DB(ctx).Preload("Assignment").Preload("User").First(&submission, id).Error
	return &submission, err
}

func (r *submissionRepository) FindByAssignmentID(ctx context.Context, assignmentID uint) ([]models.Submission, error) {
	var submissions []models.Submission
	err := r.DB(ctx).Where("assignment_id = ?", assignmentID).Find(&submissions).Error
	return submissions, err
}

func (r *submissionRepository) FindByUserID(ctx context.Context, userID uint) ([]models.Submission, error) {
	var submissions []models.Submission
	err := r.DB(ctx).Where("user_id = ?", userID).Find(&submissions).Error
	return submissions, err
}

func (r *submissionRepository) Update(ctx context.Context, submission *models.Submission) error {
	return r.DB(ctx).Save(submission).Error
}

func (r *submissionRepository) Delete(ctx context.Context, id uint) error {
	return r.DB(ctx).Delete(&models.Submission{}, id).Error
}
