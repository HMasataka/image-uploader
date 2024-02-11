package usecase

import (
	"context"
	"mime/multipart"

	"github.com/HMasataka/image-uploader/domain/factory"
	"github.com/HMasataka/image-uploader/domain/model/inventory"
	"github.com/HMasataka/image-uploader/domain/repository"
)

type ImageUseCase interface {
	Post(ctx context.Context, fileSrc multipart.File) (inventory.Image, error)
}

type imageUseCase struct {
	bucketRepository repository.BucketRepository
}

func NewImageUseCase(bucketRepository repository.BucketRepository) ImageUseCase {
	return &imageUseCase{
		bucketRepository: bucketRepository,
	}
}

func (u *imageUseCase) Post(ctx context.Context, fileSrc multipart.File) (inventory.Image, error) {
	imageID, err := factory.NewImageID()
	if err != nil {
		return inventory.Image{}, err
	}

	url, err := u.bucketRepository.Write(ctx, imageID, fileSrc)
	if err != nil {
		return inventory.Image{}, err
	}

	return inventory.Image{
		ID:  imageID,
		URL: url,
	}, nil
}
