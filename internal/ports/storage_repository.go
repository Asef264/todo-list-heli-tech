package ports

import (
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"io/ioutil"
	"todo-list/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/minio/minio-go/v7"
)

type Storage interface {
	Upload(ctx context.Context, file []byte, filename string) error
	Download(ctx context.Context, filename string) ([]byte, error)
}

type s3Storage struct {
	Client *s3.Client
}

func NewS3Storage(client *s3.Client) Storage {
	return &s3Storage{
		Client: client,
	}
}

func (s *s3Storage) Upload(ctx context.Context, file []byte, filename string) error {
	_, err := s.Client.PutObject(ctx,
		&s3.PutObjectInput{
			Bucket: aws.String("helitech"),
			Key:    aws.String(filename),
			Body:   bytes.NewReader(file),
		},
	)
	return err
}

func (s *s3Storage) Download(ctx context.Context, filename string) ([]byte, error) {
	obj, err := s.Client.GetObject(ctx,
		&s3.GetObjectInput{
			Bucket: aws.String(config.AppConfig.S3Config.Bucket),
			Key:    aws.String(filename),
		})
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(obj.Body)
	if err != nil {
		return nil, err
	}
	base64Content := base64.StdEncoding.EncodeToString(data)
	return []byte(base64Content), nil
}

type minioStorage struct {
	client *minio.Client
}

func NewMinioStorage(client *minio.Client) Storage {
	return &minioStorage{
		client: client,
	}
}

func (s *minioStorage) Upload(ctx context.Context, file []byte, fileName string) error {
	_, err := s.client.PutObject(ctx, "helitech", fileName, bytes.NewReader(file), int64(len(file)), minio.PutObjectOptions{})
	return err

}

func (s *minioStorage) Download(ctx context.Context, filename string) ([]byte, error) {
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
