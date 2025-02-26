package services

import (
	"context"
	"evconn/internal/core/domain/models"
	"evconn/internal/core/ports"
	"fmt"
	"io"
	"time"
)

type submissionService struct {
	*BaseService
	submissionRepo ports.SubmissionRepository
	storageService ports.StorageService // Add this field
}

func NewSubmissionService(submissionRepo ports.SubmissionRepository, storageService ports.StorageService) ports.SubmissionService {
	return &submissionService{
		BaseService:    NewBaseService(),
		submissionRepo: submissionRepo,
		storageService: storageService, // Initialize the field
	}
}

func (s *submissionService) CreateSubmission(ctx context.Context, submission *models.Submission) error {
	// Format time as string in MySQL datetime format
	submission.SubmittedAt = time.Now().Format("2006-01-02 15:04:05")
	submission.Status = "pending"
	return s.submissionRepo.Create(ctx, submission)
}

func (s *submissionService) GetSubmission(ctx context.Context, id uint) (*models.Submission, error) {
	return s.submissionRepo.FindByID(ctx, id)
}

func (s *submissionService) GetSubmissionsByAssignment(ctx context.Context, assignmentID uint) ([]models.Submission, error) {
	return s.submissionRepo.FindByAssignmentID(ctx, assignmentID)
}

func (s *submissionService) GetSubmissionsByUser(ctx context.Context, userID uint) ([]models.Submission, error) {
	return s.submissionRepo.FindByUserID(ctx, userID)
}

func (s *submissionService) GradeSubmission(ctx context.Context, id uint, score float64) error {
	submission, err := s.submissionRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	submission.Score = &score
	submission.Status = "graded"
	return s.submissionRepo.Update(ctx, submission)
}

// Add method to handle file uploads
func (s *submissionService) StoreSubmissionFile(ctx context.Context, submissionID uint, fileName string, fileReader io.Reader, fileSize int64) error {
	// Use MinioStorage to store the file
	objectName := fmt.Sprintf("submissions/%d/%s", submissionID, fileName)
	return s.storageService.UploadFile("submissions", objectName, fileReader, fileSize)
}
