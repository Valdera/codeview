package persistence

import (
	"context"

	"cloud.google.com/go/storage"
)

type GCloudStorageConfiguration struct {
	BucketName string
}

type GCloudStorage struct {
	Client     *storage.Client
	BucketName string
}

func InitGCloudStorage(cfg GCloudStorageConfiguration) (*GCloudStorage, error) {
	storage, err := storage.NewClient(context.Background())
	if err != nil {
		return nil, err
	}

	return &GCloudStorage{
		Client:     storage,
		BucketName: cfg.BucketName,
	}, nil
}
