package ports

import (
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/minio/minio-go/v7"
)

type Storage interface {
	Upload(ctx context.Context, file []byte, filename string, isMock bool) error
	Download(ctx context.Context, filename string, isMock bool) ([]byte, error)
}

type s3Storage struct {
	client     *s3.S3
	mockClient map[string][]byte
}

func NewS3Storage(client *s3.S3, mockClient map[string][]byte) Storage {
	return &s3Storage{
		client:     client,
		mockClient: make(map[string][]byte),
	}
}

var bucket string = "helitech-storage"

func (s *s3Storage) Upload(ctx context.Context, file []byte, filename string, isMock bool) error {
	if isMock {
		s.mockClient[filename] = file
		return nil
	}
	_, err := s.client.PutObject(&s3.PutObjectInput{
		Body:   strings.NewReader(string(file)),
		Bucket: &bucket,
		Key:    &filename,
	},
	)
	return err
}

func (s *s3Storage) Download(ctx context.Context, filename string, isMock bool) ([]byte, error) {
	if isMock {
		res := s.mockClient[filename]
		return res, nil
	}
	obj, err := s.client.GetObject(
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(filename),
		})
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(obj.Body)
	if err != nil {
		return nil, err
	}

	base64Content := base64.StdEncoding.EncodeToString(data)

	return []byte(base64Content), nil
}

type minioStorage struct {
	client     *minio.Client
	mockClient map[string][]byte
}

func NewMinioStorage(client *minio.Client, mockClient map[string][]byte) Storage {
	return &minioStorage{
		client:     client,
		mockClient: make(map[string][]byte),
	}
}

func (s *minioStorage) Upload(ctx context.Context, file []byte, fileName string, isMock bool) error {
	if isMock {
		s.mockClient[fileName] = file
		return nil
	}
	_, err := s.client.PutObject(ctx, "helitech", fileName, bytes.NewReader(file), int64(len(file)), minio.PutObjectOptions{})
	return err

}

func (s *minioStorage) Download(ctx context.Context, filename string, isMock bool) ([]byte, error) {
	if isMock {
		res := s.mockClient[filename]
		return res, nil
	}
	obj, err := s.client.GetObject(ctx, "helitech", filename, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer obj.Close()

	data, err := io.ReadAll(obj)
	if err != nil {
		return nil, err
	}

	base64Content := base64.StdEncoding.EncodeToString(data)

	return []byte(base64Content), nil
}
