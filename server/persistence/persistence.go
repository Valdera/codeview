package persistence

import (
	"codeview/config"
	"log"
	"sync"

	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

type Persistence struct {
	mu        sync.Mutex
	GCStorage *GCloudStorage
	Postgres  *gorm.DB
}

func (p *Persistence) Close() error {
	postgres, _ := p.Postgres.DB()
	if err := postgres.Close(); err != nil {
		return err
	}

	gcStorage := p.GCStorage.Client
	if err := gcStorage.Close(); err != nil {
		return err
	}

	return nil
}

func Init(config config.Config) (*Persistence, error) {
	persistence := &Persistence{}
	g := new(errgroup.Group)

	g.Go(func() error {
		log.Printf("Connecting to GCloud Storage\n")
		storage, err := InitGCloudStorage(GCloudStorageConfiguration{
			BucketName: config.GC_BUCKET_NAME,
		})
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
		postgres, err := InitPostgres(PostgresConfiguration{
			Host:         config.POSTGRES_HOST,
			Port:         config.POSTGRES_P0RT,
			User:         config.POSTGRES_USER,
			Password:     config.POSTGRES_PASSWORD,
			DatabaseName: config.POSTGRES_DBNAME,
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
