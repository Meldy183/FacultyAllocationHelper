package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
}

type Server struct {
	Host         string        `yaml:"host" env:"SERVER_HOST"`
	Port         string        `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	IdleTimeout  time.Duration `yaml:"idle_timeout"`
}

type Database struct {
	Host               string        `yaml:"host" env:"DB_HOST"`
	Port               string        `yaml:"port"`
	User               string        `yaml:"user" env:"POSTGRES_USER"`
	Password           string        `yaml:"password" env:"POSTGRES_PASSWORD"`
	DatabaseName       string        `yaml:"name"`
	SSLMode            string        `yaml:"ssl_mode"`
	MaxIdleConnections int           `yaml:"max_idle_connections"`
	MaxOpenConnections int           `yaml:"max_open_connections"`
	ConnMaxLifetime    time.Duration `yaml:"connection_max_lifetime"`
	ConnTimeout        time.Duration `yaml:"connection_timeout"`
}

func MustLoadConfig() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable not set")
	}
	if _, errExtstance := os.Stat(configPath); os.IsNotExist(errExtstance) {
		log.Fatalf("CONFIG_PATH does not exist: %s", configPath)
	}
	var cfg Config
	if errRead := cleanenv.ReadConfig(configPath, &cfg); errRead != nil {
		log.Fatal(errRead)
	}
	return &cfg
}
