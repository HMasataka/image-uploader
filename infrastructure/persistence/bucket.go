package persistence

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"

	"cloud.google.com/go/storage"
	"github.com/HMasataka/config"
	"github.com/HMasataka/image-uploader/domain/repository"
	"github.com/rs/zerolog/log"
)

type bucketRepository struct {
	bucketName    string
	storageClient *storage.Client
}

func NewBucketRepository(cfg *config.Config, storageClient *storage.Client) repository.BucketRepository {
	return &bucketRepository{
		bucketName:    cfg.GcsBucketName,
		storageClient: storageClient,
	}
}

func (b *bucketRepository) Write(ctx context.Context, key string, fileSrc multipart.File) (string, error) {
	writer := b.storageClient.Bucket(b.bucketName).Object(key).NewWriter(ctx)
	defer func() {
		log.Debug().Err(writer.Close()).Send()
	}()

	buf := make([]byte, 1024)

	for {
		n, err := fileSrc.Read(buf)
		if err != nil && !errors.Is(err, io.EOF) {
			return "", err
		}

		if n == 0 {
			break
		}

		if _, err := writer.Write(buf[:n]); err != nil {
			return "", err
		}
	}

	url := fmt.Sprintf("https://storage.googleapis.com/%v/%v", b.bucketName, key)

	return url, nil
}
