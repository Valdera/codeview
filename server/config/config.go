package config

import (
	"io/ioutil"
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
	MaxBodyBytes          int64  `yaml:"maxBodyBytes" envconfig:"MAX_BODY_BYTES" default:""`
	ApiKey                string `yaml:"apiKey" envconfig:"X_API_KEY" default:""`
	JWTPublicKeyFilePath  string `yaml:"jwtPublicKeyFile" envconfig:"JWT_PUBLIC_KEY_FILE_PATH" default:""`
	JWTPrivateKeyFilePath string `yaml:"jwtPrivateKeyFile" envconfig:"JWT_PRIVATE_KEY_FILE_PATH" default:""`

	JWTPublicKey  []byte
	JWTPrivateKey []byte

	// It is recommended to use an authentication key with 32 or 64 bytes. The encryption key,
	// if set, must be either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256 modes.
	SessionKey string `yaml:"sessionKey" envconfig:"SESSION_KEY" default:""`

	// Persistence Configuration
	Postgres  PostgresConfig `yaml:"postgres"`
	GCStorage GCloudConfig   `yaml:"gcstorage"`
	Redis     RedisConfig    `yaml:"redis"`

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

func (cfg *AppConfig) LoadCerts() error {
	privateKey, err := ioutil.ReadFile(cfg.JWTPrivateKeyFilePath)
	if err != nil {
		return err
	}

	cfg.JWTPrivateKey = privateKey

	publicKey, err := ioutil.ReadFile(cfg.JWTPublicKeyFilePath)
	if err != nil {
		return err
	}

	cfg.JWTPublicKey = publicKey

	return nil
}
