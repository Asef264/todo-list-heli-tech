package ports

import "context"

type Storage interface {
	Upload(ctx context.Context, file []byte, filename string, isMock bool) error
	Download(ctx context.Context, filename string, isMock bool) ([]byte, error)
}
