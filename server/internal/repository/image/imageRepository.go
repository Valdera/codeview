package repository

import (
	"codeview/domain"
	"codeview/exception"
	"codeview/internal/entity"
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"

	"cloud.google.com/go/storage"
)

type imageRepository struct {
	client     *storage.Client
	bucketName string
}

func New(gcClient *storage.Client, bucketName string) domain.ImageRepository {
	return &imageRepository{
		client:     gcClient,
		bucketName: bucketName,
	}
}

func (r *imageRepository) AddImage(ctx context.Context, objName string, imageFile multipart.File) (*entity.Image, error) {
	bckt := r.client.Bucket(r.bucketName)

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
		r.bucketName,
		objName,
	)

	return &entity.Image{
		URL: imageURL,
	}, nil

}
