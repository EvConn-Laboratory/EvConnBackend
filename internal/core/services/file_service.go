package services

import (
	"context"
	"evconn/internal/core/domain/models"
	"evconn/internal/core/ports"
	"io"
)

type fileService struct {
	*BaseService
	fileRepo       ports.FileRepository
	storageService ports.StorageService
}

func NewFileService(fileRepo ports.FileRepository, storageService ports.StorageService) ports.FileService {
	return &fileService{
		BaseService:    NewBaseService(),
		fileRepo:       fileRepo,
		storageService: storageService,
	}
}

func (s *fileService) UploadFile(ctx context.Context, file *models.File, reader io.Reader) error {
	// Upload to storage first
	err := s.storageService.UploadFile("files", file.Path, reader, file.Size)
	if err != nil {
		return err
	}

	// Save file metadata to database
	return s.fileRepo.Create(ctx, file)
}

func (s *fileService) GetFileByID(ctx context.Context, id uint) (*models.File, error) {
	return s.fileRepo.FindByID(ctx, id)
}

func (s *fileService) GetFilesByEntity(ctx context.Context, entityType string, entityID uint) ([]*models.File, error) {
	return s.fileRepo.FindByEntityTypeAndID(ctx, entityType, entityID)
}

func (s *fileService) DeleteFile(ctx context.Context, id uint) error {
	// Get file first to get the path
	file, err := s.fileRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	// Delete from storage first
	err = s.storageService.DeleteFile("files", file.Path)
	if err != nil {
		return err
	}

	// Delete metadata from database
	return s.fileRepo.Delete(ctx, id)
}

func (s *fileService) GetFileURL(ctx context.Context, id uint) (string, error) {
	file, err := s.fileRepo.FindByID(ctx, id)
	if err != nil {
		return "", err
	}

	// Get presigned URL with 1-hour expiry
	url, err := s.storageService.GetFileURL("files", file.Path, 3600)
	if err != nil {
		return "", err
	}

	return url, nil
}
