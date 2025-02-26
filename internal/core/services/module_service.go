package services

import (
	"context"
	"evconn/internal/core/domain/models"
	"evconn/internal/core/ports"
)

type moduleService struct {
	*BaseService
	moduleRepo       ports.ModuleRepository
	moduleMentorRepo ports.ModuleMentorRepository
}

func NewModuleService(moduleRepo ports.ModuleRepository, moduleMentorRepo ports.ModuleMentorRepository) ports.ModuleService {
	return &moduleService{
		BaseService:      NewBaseService(),
		moduleRepo:       moduleRepo,
		moduleMentorRepo: moduleMentorRepo,
	}
}

func (s *moduleService) CreateModule(ctx context.Context, module *models.Module) error {
	return s.moduleRepo.Create(ctx, module)
}

func (s *moduleService) GetModule(ctx context.Context, id uint) (*models.Module, error) {
	return s.moduleRepo.FindByID(ctx, id)
}

func (s *moduleService) GetModulesByCourse(ctx context.Context, courseID uint) ([]models.Module, error) {
	return s.moduleRepo.FindByCourseID(ctx, courseID)
}

func (s *moduleService) UpdateModule(ctx context.Context, module *models.Module) error {
	return s.moduleRepo.Update(ctx, module)
}

func (s *moduleService) DeleteModule(ctx context.Context, id uint) error {
	return s.moduleRepo.Delete(ctx, id)
}

func (s *moduleService) AssignMentor(ctx context.Context, moduleID, mentorID uint) error {
	moduleMentor := &models.ModuleMentor{
		ModuleID: moduleID,
		MentorID: mentorID,
	}
	return s.moduleMentorRepo.Create(ctx, moduleMentor)
}
