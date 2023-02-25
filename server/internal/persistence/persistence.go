package persistence

import (
	"codeview/config"
	gcloudstorage "codeview/internal/persistence/gcloud/storage"
	postgres "codeview/internal/persistence/postgres"
	"log"
	"sync"

	"cloud.google.com/go/storage"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

type Persistence struct {
	mu        sync.Mutex
	GCStorage *storage.Client
	Postgres  *gorm.DB
}

func (p *Persistence) Close() error {
	postgres, _ := p.Postgres.DB()
	if err := postgres.Close(); err != nil {
		return err
	}

	gcStorage := p.GCStorage
	if err := gcStorage.Close(); err != nil {
		return err
	}

	return nil
}

func Init(config config.AppConfig) (*Persistence, error) {
	persistence := &Persistence{}
	g := new(errgroup.Group)

	g.Go(func() error {
		log.Printf("Connecting to GCloud Storage\n")
		storage, err := gcloudstorage.Init(gcloudstorage.Configuration{})
		if err != nil {
			log.Fatalf("error creating cloud storage client: %v", err)
			return err
		}

		persistence.mu.Lock()
		persistence.GCStorage = storage
		persistence.mu.Unlock()

		log.Printf("Connected to GCloud Storage\n")

		return nil
	})

	g.Go(func() error {
		log.Printf("Connecting to Postgres Database\n")
		postgres, err := postgres.Init(postgres.Configuration{
			Host:         config.Postgres.Host,
			Port:         config.Postgres.Port,
			User:         config.Postgres.User,
			Password:     config.Postgres.Password,
			DatabaseName: config.Postgres.DatabaseName,
		})
		if err != nil {
			log.Fatalf("error creating postgres client: %v", err)
			return err
		}

		persistence.mu.Lock()
		persistence.Postgres = postgres
		persistence.mu.Unlock()
		log.Printf("Connected to Postgres Database\n")

		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	log.Println("Successfully connected to all persitence instances")
	return persistence, nil
}
