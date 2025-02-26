package services

import (
	"context"
	"evconn/internal/core/domain/models"
	"evconn/internal/core/ports"
)

type feedbackService struct {
	feedbackRepo ports.FeedbackRepository
}

func NewFeedbackService(feedbackRepo ports.FeedbackRepository) ports.FeedbackService {
	return &feedbackService{feedbackRepo: feedbackRepo}
}

func (s *feedbackService) CreateFeedback(ctx context.Context, feedback *models.Feedback) error {
	return s.feedbackRepo.Create(ctx, feedback)
}

func (s *feedbackService) GetFeedbackByID(ctx context.Context, id uint) (*models.Feedback, error) {
	return s.feedbackRepo.FindByID(ctx, id)
}

func (s *feedbackService) GetFeedbackByModule(ctx context.Context, moduleID uint) ([]models.Feedback, error) {
	return s.feedbackRepo.FindByModuleID(ctx, moduleID)
}

func (s *feedbackService) GetFeedbackByUser(ctx context.Context, userID uint) ([]models.Feedback, error) {
	return s.feedbackRepo.FindByUserID(ctx, userID)
}

func (s *feedbackService) GetFeedbackByModuleAndType(ctx context.Context, moduleID uint, feedbackType string) ([]models.Feedback, error) {
	return s.feedbackRepo.FindByModuleAndType(ctx, moduleID, feedbackType)
}

func (s *feedbackService) UpdateFeedback(ctx context.Context, feedback *models.Feedback) error {
	return s.feedbackRepo.Update(ctx, feedback)
}

func (s *feedbackService) DeleteFeedback(ctx context.Context, id uint) error {
	return s.feedbackRepo.Delete(ctx, id)
}
