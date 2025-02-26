package storage

import (
	"context"
	"evconn/internal/core/ports"
	"fmt"
	"io"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioConfig struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	UseSSL    bool
}

type minioStorage struct {
	client *minio.Client
}

func NewMinioStorage(config *MinioConfig) (ports.StorageService, error) {
	// Initialize minio client object.
	client, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.SecretKey, ""),
		Secure: config.UseSSL,
	})

	if err != nil {
		return nil, err
	}

	return &minioStorage{
		client: client,
	}, nil
}

func (s *minioStorage) UploadFile(bucketName string, objectName string, reader io.Reader, size int64) error {
	ctx := context.Background()

	// Create bucket if not exists
	exists, err := s.client.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}

	if !exists {
		err = s.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
	}

	// Upload object
	_, err = s.client.PutObject(ctx, bucketName, objectName, reader, size, minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})

	return err
}

func (s *minioStorage) GetFileURL(bucketName string, objectName string, expiry int64) (string, error) {
	ctx := context.Background()

	// Set default expiry to 24 hours if not provided
	if expiry <= 0 {
		expiry = 24 * 60 * 60 // 24 hours
	}

	// Generate presigned URL
	reqParams := make(url.Values)
	presignedURL, err := s.client.PresignedGetObject(ctx, bucketName, objectName, time.Duration(expiry)*time.Second, reqParams)
	if err != nil {
		return "", err
	}

	return presignedURL.String(), nil
}

func (s *minioStorage) DeleteFile(bucketName string, objectName string) error {
	ctx := context.Background()

	err := s.client.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete object: %w", err)
	}

	return nil
}
