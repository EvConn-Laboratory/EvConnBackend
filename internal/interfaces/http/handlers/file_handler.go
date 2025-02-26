package handlers

import (
	"evconn/internal/core/domain/models"
	"evconn/internal/core/ports"
	"evconn/internal/pkg/response"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	fileService ports.FileService
}

func NewFileHandler(fileService ports.FileService) *FileHandler {
	return &FileHandler{
		fileService: fileService,
	}
}

// Helper function to determine file type from extension
func getFileType(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))

	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp":
		return "image"
	case ".pdf":
		return "pdf"
	case ".doc", ".docx":
		return "document"
	case ".xls", ".xlsx":
		return "spreadsheet"
	case ".ppt", ".pptx":
		return "presentation"
	case ".zip", ".rar", ".tar", ".gz":
		return "archive"
	case ".mp3", ".wav", ".ogg", ".flac":
		return "audio"
	case ".mp4", ".avi", ".mov", ".wmv":
		return "video"
	case ".txt", ".md", ".json", ".xml", ".csv":
		return "text"
	default:
		return "other"
	}
}

// Upload handles file upload for various entity types
func (h *FileHandler) Upload(c *gin.Context) {
	// Get entity type and ID from path
	entityType := c.Param("entityType")
	entityId, err := strconv.ParseUint(c.Param("entityId"), 10, 32)
	if err != nil {
		response.Error(c, 400, "Invalid entity ID")
		return
	}

	// Get authenticated user
	user, exists := c.Get("user")
	if !exists {
		response.Error(c, 401, "Unauthorized")
		return
	}
	currentUser := user.(*models.User)

	// Get file from request
	file, err := c.FormFile("file")
	if err != nil {
		response.Error(c, 400, fmt.Sprintf("File upload error: %v", err))
		return
	}

	// Open file
	fileReader, err := file.Open()
	if err != nil {
		response.Error(c, 500, fmt.Sprintf("Failed to open file: %v", err))
		return
	}
	defer fileReader.Close()

	// Prepare file model
	fileModel := &models.File{
		Name:       file.Filename,
		Type:       getFileType(file.Filename),
		Size:       file.Size,
		Path:       fmt.Sprintf("%s/%d/%s", entityType, entityId, file.Filename),
		EntityType: entityType,
		EntityID:   uint(entityId),
		UserID:     currentUser.ID,
		UploadedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	// Upload file
	err = h.fileService.UploadFile(c.Request.Context(), fileModel, fileReader)
	if err != nil {
		response.Error(c, 500, fmt.Sprintf("Failed to upload file: %v", err))
		return
	}

	response.Success(c, fileModel)
}

// GetByID retrieves file information and generates a download URL
func (h *FileHandler) GetByID(c *gin.Context) {
	// Get file ID from path
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, 400, "Invalid file ID")
		return
	}

	// Get file information
	file, err := h.fileService.GetFileByID(c.Request.Context(), uint(id))
	if err != nil {
		response.Error(c, 404, "File not found")
		return
	}

	// Generate download URL
	url, err := h.fileService.GetFileURL(c.Request.Context(), file.ID)
	if err != nil {
		response.Error(c, 500, "Failed to generate download URL")
		return
	}

	// Add URL to response
	fileResponse := map[string]interface{}{
		"id":         file.ID,
		"name":       file.Name,
		"type":       file.Type,
		"size":       file.Size,
		"entityType": file.EntityType,
		"entityID":   file.EntityID,
		"userID":     file.UserID,
		"uploadedAt": file.UploadedAt,
		"url":        url,
	}

	response.Success(c, fileResponse)
}

// GetByEntity retrieves files associated with an entity
func (h *FileHandler) GetByEntity(c *gin.Context) {
	// Get entity type and ID from path
	entityType := c.Param("entityType")
	entityId, err := strconv.ParseUint(c.Param("entityId"), 10, 32)
	if err != nil {
		response.Error(c, 400, "Invalid entity ID")
		return
	}

	// Get files
	files, err := h.fileService.GetFilesByEntity(c.Request.Context(), entityType, uint(entityId))
	if err != nil {
		response.Error(c, 500, "Failed to retrieve files")
		return
	}

	response.Success(c, files)
}

// Delete removes a file
func (h *FileHandler) Delete(c *gin.Context) {
	// Get file ID from path
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, 400, "Invalid file ID")
		return
	}

	// Get file information to verify ownership
	file, err := h.fileService.GetFileByID(c.Request.Context(), uint(id))
	if err != nil {
		response.Error(c, 404, "File not found")
		return
	}

	// Check if user is allowed to delete this file
	user, _ := c.Get("user")
	currentUser := user.(*models.User)
	if file.UserID != currentUser.ID && currentUser.Role != "admin" {
		response.Error(c, 403, "Not authorized to delete this file")
		return
	}

	// Delete file
	err = h.fileService.DeleteFile(c.Request.Context(), uint(id))
	if err != nil {
		response.Error(c, 500, "Failed to delete file")
		return
	}

	response.Success(c, gin.H{"message": "File deleted successfully"})
}
