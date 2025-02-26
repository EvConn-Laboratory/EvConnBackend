package repositories

import (
	"context"
	"evconn/internal/core/domain/models"

	"gorm.io/gorm"
)

type feedbackRepository struct {
	db *gorm.DB
}

func NewFeedbackRepository(db *gorm.DB) *feedbackRepository {
	return &feedbackRepository{db: db}
}

func (r *feedbackRepository) Create(ctx context.Context, feedback *models.Feedback) error {
	return r.db.WithContext(ctx).Create(feedback).Error
}

func (r *feedbackRepository) FindByID(ctx context.Context, id uint) (*models.Feedback, error) {
	var feedback models.Feedback
	err := r.db.WithContext(ctx).First(&feedback, id).Error
	if err != nil {
		return nil, err
	}
	return &feedback, nil
}

// Update the return type to match the interface expectation
func (r *feedbackRepository) FindByModuleID(ctx context.Context, moduleID uint) ([]models.Feedback, error) {
	var feedbacks []models.Feedback
	err := r.db.WithContext(ctx).Where("module_id = ?", moduleID).Find(&feedbacks).Error
	if err != nil {
		return nil, err
	}
	return feedbacks, nil
}

func (r *feedbackRepository) FindByUserID(ctx context.Context, userID uint) ([]models.Feedback, error) {
	var feedbacks []models.Feedback
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&feedbacks).Error
	if err != nil {
		return nil, err
	}
	return feedbacks, nil
}

func (r *feedbackRepository) FindByModuleAndType(ctx context.Context, moduleID uint, feedbackType string) ([]models.Feedback, error) {
	var feedbacks []models.Feedback
	err := r.db.WithContext(ctx).Where("module_id = ? AND type = ?", moduleID, feedbackType).Find(&feedbacks).Error
	if err != nil {
		return nil, err
	}
	return feedbacks, nil
}

func (r *feedbackRepository) Update(ctx context.Context, feedback *models.Feedback) error {
	return r.db.WithContext(ctx).Save(feedback).Error
}

func (r *feedbackRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Feedback{}, id).Error
}
