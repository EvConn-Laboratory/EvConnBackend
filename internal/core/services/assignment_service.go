package services

import (
	"context"
	"evconn/internal/core/domain/models"
	"evconn/internal/core/ports"
	"time"
)

type assignmentService struct {
	*BaseService
	assignmentRepo ports.AssignmentRepository
	submissionRepo ports.SubmissionRepository
}

func NewAssignmentService(
	assignmentRepo ports.AssignmentRepository,
	submissionRepo ports.SubmissionRepository,
) ports.AssignmentService {
	return &assignmentService{
		BaseService:    NewBaseService(),
		assignmentRepo: assignmentRepo,
		submissionRepo: submissionRepo,
	}
}

func (s *assignmentService) GetPendingAssignments(ctx context.Context, userID uint) ([]*models.Assignment, error) {
	// Get all submissions by user
	submissions, err := s.submissionRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// Create a map of completed assignment IDs
	completedAssignments := make(map[uint]bool)
	for _, sub := range submissions {
		completedAssignments[sub.AssignmentID] = true
	}

	// Get all assignments
	assignments, err := s.assignmentRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// Filter out completed assignments
	var pendingAssignments []*models.Assignment
	for i := range assignments {
		if !completedAssignments[assignments[i].ID] {
			pendingAssignments = append(pendingAssignments, &assignments[i])
		}
	}

	return pendingAssignments, nil
}

func (s *assignmentService) GetAssignmentsByModule(ctx context.Context, moduleID uint) ([]models.Assignment, error) {
	return s.assignmentRepo.FindByModuleID(ctx, moduleID)
}

func (s *assignmentService) CreateAssignment(ctx context.Context, assignment *models.Assignment) error {
	// Set default deadline if not provided
	if assignment.Deadline == "" {
		assignment.Deadline = time.Now().AddDate(0, 0, 7).Format("2006-01-02 15:04:05") // Default 1 week
	}
	return s.assignmentRepo.Create(ctx, assignment)
}

func (s *assignmentService) UpdateAssignment(ctx context.Context, assignment *models.Assignment) error {
	return s.assignmentRepo.Update(ctx, assignment)
}

func (s *assignmentService) DeleteAssignment(ctx context.Context, id uint) error {
	return s.assignmentRepo.Delete(ctx, id)
}

func (s *assignmentService) GetAssignment(ctx context.Context, id uint) (*models.Assignment, error) {
	return s.assignmentRepo.FindByID(ctx, id)
}
