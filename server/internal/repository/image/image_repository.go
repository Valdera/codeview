package repository

import (
	"codeview/domain"
	"codeview/exception"
	"codeview/internal/entity"
	"codeview/persistence"
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
)

type imageRepository struct {
	storage *persistence.GCloudStorage
}

func New(gcClient *persistence.GCloudStorage) domain.ImageRepository {
	return &imageRepository{
		storage: gcClient,
	}
}

func (r *imageRepository) UploadImage(ctx context.Context, objName string, imageFile multipart.File) (*entity.Image, error) {
	bckt := r.storage.Client.Bucket(r.storage.BucketName)

	object := bckt.Object(objName)
	wc := object.NewWriter(ctx)

	// set cache control so profile image will be served fresh by browsers
	// To do this with object handle, you'd first have to upload, then update
	wc.ObjectAttrs.CacheControl = "Cache-Control:no-cache, max-age=0"
	// multipart.File has a reader!
	if _, err := io.Copy(wc, imageFile); err != nil {
		log.Printf("Unable to write file to Google Cloud Storage: %v\n", err)
		return nil, exception.NewInternal()
	}

	if err := wc.Close(); err != nil {
		return nil, fmt.Errorf("Writer.Close: %v", err)
	}

	imageURL := fmt.Sprintf(
		"https://storage.googleapis.com/%s/%s",
		r.storage.BucketName,
		objName,
	)

	return &entity.Image{
		URL: imageURL,
	}, nil

}
