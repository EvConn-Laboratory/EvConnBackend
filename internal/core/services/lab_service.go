package services

import (
	"context"
	"evconn/internal/core/domain/models"
	"evconn/internal/core/ports"
)

type labService struct {
	labRepo ports.LabRepository
}

func NewLabService(labRepo ports.LabRepository) ports.LabService {
	return &labService{labRepo: labRepo}
}

func (s *labService) GetLabs(ctx context.Context) ([]*models.Lab, error) {
	return s.labRepo.FindAll(ctx)
}

func (s *labService) GetLabByID(ctx context.Context, id uint) (*models.Lab, error) {
	return s.labRepo.FindByID(ctx, id)
}

func (s *labService) CreateLab(ctx context.Context, lab *models.Lab) error {
	return s.labRepo.Create(ctx, lab)
}

func (s *labService) UpdateLab(ctx context.Context, lab *models.Lab) error {
	return s.labRepo.Update(ctx, lab)
}

func (s *labService) DeleteLab(ctx context.Context, id uint) error {
	return s.labRepo.Delete(ctx, id)
}
