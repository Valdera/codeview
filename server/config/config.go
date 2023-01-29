package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	REST_PORT        string `envconfig:"REST_PORT" default:":8080"`
	GC_BUCKET_NAME   string `envconfig:"GC_BUCKET_NAME" default:""`
	MAX_BODY_BYTES   int64  `envconfig:"MAX_BODY_BYTES" default:""`
	TIMEOUT_DURATION int64  `envconfig:"TIMEOUT_DURATION" default:"5"`
}

// Get to get defined configuration
func Get() Config {
	_ = godotenv.Load(".env")
	cfg := Config{}
	envconfig.MustProcess("", &cfg)

	fmt.Println(cfg)
	return cfg
}
