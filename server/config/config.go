package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	REST_PORT         string `envconfig:"REST_PORT" default:":8080"`
	POSTGRES_USER     string `envconfig:"POSTGRES_USER" default:""`
	POSTGRES_PASSWORD string `envconfig:"POSTGRES_PASSWORD" default:""`
	POSTGRES_DBNAME   string `envconfig:"POSTGRES_DBNAME" default:""`
	POSTGRES_HOST     string `envconfig:"POSTGRES_HOST" default:""`
	POSTGRES_P0RT     string `envconfig:"POSTGRES_PORT" default:""`
	GC_BUCKET_NAME    string `envconfig:"GC_BUCKET_NAME" default:""`
	MAX_BODY_BYTES    int64  `envconfig:"MAX_BODY_BYTES" default:""`
	TIMEOUT_DURATION  int64  `envconfig:"TIMEOUT_DURATION" default:"5"`
}

// Get to get defined configuration
func Get() Config {
	_ = godotenv.Load(".env")
	cfg := Config{}
	envconfig.MustProcess("", &cfg)

	fmt.Println(cfg)
	return cfg
}
