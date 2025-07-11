package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server HTTPServer `yaml:"server"`
	DB     Database   `yaml:"database"`
}
type HTTPServer struct {
	Host         string        `yaml:"host" env:"ACADEMIC_HOST"`
	Port         string        `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	IdleTimeout  time.Duration `yaml:"idle_timeout"`
}
type Database struct {
	Host            string        `yaml:"host" env:"DB_HOST"`
	Port            string        `yaml:"port"`
	DBName          string        `yaml:"name"`
	User            string        `yaml:"user"`
	Password        string        `yaml:"password" env:"DB_PASSWORD"`
	SSLMode         string        `yaml:"ssl_mode"`
	MaxOpenConns    int           `yaml:"max_open_connections"`
	MaxIdleConns    int           `yaml:"max_idle_connections"`
	ConnMaxLifetime time.Duration `yaml:"connection_max_lifetime"`
}

func MustLoadCfg() (*Config, error) {
	cfgPath := os.Getenv("ACADEMIC_CONFIG_PATH")
	if cfgPath == "" {
		log.Fatal("CONFIG_PATH environment variable is not set")
	}
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		log.Fatalf("Configuration file does not exist at path: %s", cfgPath)
	}
	var cfg Config
	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		log.Fatalf("Failed to read configuration file: %v", err)
	}

	return &cfg, nil
}
