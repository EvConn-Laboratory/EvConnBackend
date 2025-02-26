package ports

import "io"

// StorageService defines interface for file storage operations
type StorageService interface {
	UploadFile(bucketName string, objectName string, reader io.Reader, size int64) error
	GetFileURL(bucketName string, objectName string, expiry int64) (string, error)
	DeleteFile(bucketName string, objectName string) error
}
