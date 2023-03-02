package client

import (
	"context"

	"cloud.google.com/go/storage"
)

type Configuration struct {
}

func Init(cfg Configuration) (*storage.Client, error) {
	storage, err := storage.NewClient(context.Background())
	if err != nil {
		return nil, err
	}

	return storage, nil
}
