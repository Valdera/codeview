package repository

import (
	"codeview/config"
	"codeview/internal/entity"
	"codeview/internal/repository"
	"codeview/internal/util/exception"

	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"

	"cloud.google.com/go/storage"
)

type imageRepository struct {
	cfg     config.AppConfig
	storage *storage.Client
}

func NewImageRepository(cfg config.AppConfig, storage *storage.Client) repository.ImageRepository {
	return &imageRepository{
		cfg,
		storage,
	}
}

func (r *imageRepository) UploadImage(ctx context.Context, objName string, imageFile multipart.File) (*entity.Image, error) {
	bckt := r.storage.Bucket(r.cfg.GCStorage.BucketName)

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
		r.cfg.GCStorage.BucketName,
		objName,
	)

	return &entity.Image{
		URL: imageURL,
	}, nil

}
