package client

import (
	"fmt"

	"github.com/gin-contrib/sessions/redis"
)

type Configuration struct {
	Host       string
	Port       string
	Password   string
	SessionKey string
}

func (cfg *Configuration) GetAddress() string {
	return fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
}

func Init(cfg Configuration) (redis.Store, error) {
	store, err := redis.NewStore(3, "tcp", cfg.GetAddress(), cfg.Password, []byte(cfg.SessionKey))
	if err != nil {
		return nil, err
	}

	return store, nil
}
