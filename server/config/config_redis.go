package config

import "fmt"

type RedisConfig struct {
	Host     string `yaml:"host" envconfig:"REDIS_HOST" default:""`
	Port     string `yaml:"port" envconfig:"REDIS_PORT" default:""`
	Password string `yaml:"password" envconfig:"REDIS_PASSWORD" default:""`
}

func (cfg *RedisConfig) GetAddress() string {
	return fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
}
