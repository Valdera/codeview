package config

import "fmt"

type PostgresConfig struct {
	User         string `yaml:"user" envconfig:"POSTGRES_USER" default:""`
	Password     string `yaml:"password" envconfig:"POSTGRES_PASSWORD" default:""`
	DatabaseName string `yaml:"dbname" envconfig:"POSTGRES_DBNAME" default:""`
	Host         string `yaml:"host" envconfig:"POSTGRES_HOST" default:""`
	Port         string `yaml:"port" envconfig:"POSTGRES_PORT" default:""`
}

func (cfg *PostgresConfig) GetDatabaseURL() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", cfg.Host, cfg.User, cfg.Password, cfg.DatabaseName, cfg.Port)
}
