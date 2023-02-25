package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	// Server Configuration
	RestServer struct {
		Port string `yaml:"port" envconfig:"REST_PORT" default:":8080"`
		Host string `yaml:"host" envconfig:"REST_HOST" default:"localhost"`
	} `yaml:"rest"`

	// Application Configuration
	MaxBodyBytes int64 `yaml:"maxBodyBytes" envconfig:"MAX_BODY_BYTES" default:""`

	// Persistence Configuration
	Postgres  PostgresConfig `yaml:"postgres"`
	GCStorage struct {
		BucketName string `yaml:"bucketName" envconfig:"GC_BUCKET_NAME" default:""`
	} `yaml:"gcstorage"`

	// Migration Configuration
	Migration struct {
		FilesPath string `envconfig:"MIGRATION_FILES_PATH" default:""`
	} `yaml:"migration"`
}

func Init() AppConfig {
	return AppConfig{}
}

func (cfg *AppConfig) LoadFromEnv() error {
	_ = godotenv.Load(".env")

	envconfig.MustProcess("", cfg)

	return nil
}

func (cfg *AppConfig) LoadFromYaml() error {
	f, err := os.Open("config.yml")
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		return err
	}

	return nil
}
