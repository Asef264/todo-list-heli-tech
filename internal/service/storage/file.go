package storage_service

import (
	"context"
	ports "todo-list/internal/ports/storage"
)

type StorageService interface {
	Upload(ctx context.Context, file []byte, filename string, isMock bool) error
	Download(ctx context.Context, filename string, isMock bool) ([]byte, error)
}

type storageService struct {
	storageRepository ports.Storage
}

func NewStorageService(s ports.Storage) StorageService {
	return &storageService{
		storageRepository: s,
	}
}

func (u *storageService) Upload(ctx context.Context, file []byte, filename string, isMock bool) error {
	return u.storageRepository.Upload(ctx, file, filename, isMock)
}

func (u *storageService) Download(ctx context.Context, filename string, isMock bool) ([]byte, error) {
	return u.storageRepository.Download(ctx, filename, isMock)
}
