package persistence

import (
	"codeview/config"
	gcStorageCli "codeview/internal/persistence/client/gcloud/storage"
	postgresCli "codeview/internal/persistence/client/postgres"
	redisCli "codeview/internal/persistence/client/redis"
	sessionStoreCli "codeview/internal/persistence/client/session"
	"log"
	"sync"

	"cloud.google.com/go/storage"
	"github.com/gin-contrib/sessions/redis"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

type Persistence struct {
	mu           sync.Mutex
	GCStorage    *storage.Client
	Postgres     *gorm.DB
	Redis        redisCli.RedisDriver
	SessionStore redis.Store
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

func Init(cfg config.AppConfig) (*Persistence, error) {
	persistence := &Persistence{}
	g := new(errgroup.Group)

	g.Go(func() error {
		log.Printf("Connecting to GCloud Storage\n")
		client, err := gcStorageCli.Init(gcStorageCli.Configuration{})
		if err != nil {
			log.Fatalf("error creating cloud storage client: %v", err)
			return err
		}

		persistence.mu.Lock()
		persistence.GCStorage = client
		persistence.mu.Unlock()

		log.Printf("Connected to GCloud Storage\n")

		return nil
	})

	g.Go(func() error {
		log.Printf("Connecting to Postgres Database\n")
		client, err := postgresCli.Init(postgresCli.Configuration{
			Host:         cfg.Postgres.Host,
			Port:         cfg.Postgres.Port,
			User:         cfg.Postgres.User,
			Password:     cfg.Postgres.Password,
			DatabaseName: cfg.Postgres.DatabaseName,
		})
		if err != nil {
			log.Fatalf("error creating postgres client: %v", err)
			return err
		}

		persistence.mu.Lock()
		persistence.Postgres = client
		persistence.mu.Unlock()
		log.Printf("Connected to Postgres Database\n")

		return nil
	})

	g.Go(func() error {
		log.Printf("Connecting to Redis Database\n")
		client, err := redisCli.Init(redisCli.Configuration{
			Host:     cfg.Redis.Host,
			Port:     cfg.Redis.Port,
			Password: cfg.Redis.Password,
		})

		if err != nil {
			log.Fatalf("error creating redis client: %v", err)
			return err
		}

		persistence.mu.Lock()
		persistence.Redis = client
		persistence.mu.Unlock()
		log.Printf("Connected to Redis Database\n")

		return nil
	})

	g.Go(func() error {
		log.Printf("Connecting to Session Store\n")
		client, err := sessionStoreCli.Init(sessionStoreCli.Configuration{
			Host:       cfg.Redis.Host,
			Port:       cfg.Redis.Port,
			Password:   cfg.Redis.Password,
			SessionKey: cfg.SessionKey,
		})

		if err != nil {
			log.Fatalf("error creating session store client: %v", err)
			return err
		}

		persistence.mu.Lock()
		persistence.SessionStore = client
		persistence.mu.Unlock()

		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	log.Println("Successfully connected to all persitence instances")
	return persistence, nil
}
