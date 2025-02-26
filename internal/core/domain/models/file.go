package models

type FileType string

const (
	FileTypeDocument FileType = "document"
	FileTypeImage    FileType = "image"
	FileTypeVideo    FileType = "video"
	FileTypeOther    FileType = "other"
)

type File struct {
	Base
	Name       string `json:"name" gorm:"not null"`
	Type       string `json:"type" gorm:"not null"` // image, pdf, document, etc.
	Size       int64  `json:"size" gorm:"not null"`
	Path       string `json:"path" gorm:"not null"`        // Storage path
	EntityType string `json:"entity_type" gorm:"not null"` // submission, assignment, etc.
	EntityID   uint   `json:"entity_id" gorm:"not null"`
	UserID     uint   `json:"user_id" gorm:"not null"`
	UploadedAt string `json:"uploaded_at" gorm:"not null"`
}
