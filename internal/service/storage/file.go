package storage_service

import (
	"context"
	"todo-list/internal/ports"
)

type StorageService interface {
	Upload(ctx context.Context, file []byte, filename string) error
	Download(ctx context.Context, filename string) ([]byte, error)
}

type storageService struct {
	storageRepository ports.Storage
}

func NewStorageService(s ports.Storage) StorageService {
	return &storageService{
		storageRepository: s,
	}
}

func (u *storageService) Upload(ctx context.Context, file []byte, filename string) error {
	return u.storageRepository.Upload(ctx, file, filename)
}

func (u *storageService) Download(ctx context.Context, filename string) ([]byte, error) {
	return u.storageRepository.Download(ctx, filename)
}
