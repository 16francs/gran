package storage

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"github.com/google/uuid"
)

// Storage - Cloud Storageの構造体
type Storage struct {
	BucketName string
	Bucket     *storage.BucketHandle
}

const publicURL = "https://storage.googleapis.com/%s/%s"

// NewClient - Storageに接続
func NewClient(ctx context.Context, app *firebase.App, bucketName string) (*Storage, error) {
	client, err := app.Storage(ctx)
	if err != nil {
		return nil, err
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}

	return &Storage{bucketName, bucket}, nil
}

// Write - CLoud Storageにアップロード
func (s *Storage) Write(ctx context.Context, path string, file multipart.File) (string, error) {
	fileName := fmt.Sprintf("%s/%s", path, uuid.New().String())

	w := s.Bucket.Object(fileName).NewWriter(ctx)
	w.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
	w.CacheControl = "public, max-age=86400" // 1 day

	if _, err := io.Copy(w, file); err != nil {
		return "", err
	}

	if err := w.Close(); err != nil {
		return "", err
	}

	return fmt.Sprintf(publicURL, s.BucketName, fileName), nil
}
