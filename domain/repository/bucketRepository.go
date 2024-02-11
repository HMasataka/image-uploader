package repository

import (
	"context"
	"mime/multipart"
)

type BucketRepository interface {
	Write(ctx context.Context, key string, fileSrc multipart.File) (string, error)
}
