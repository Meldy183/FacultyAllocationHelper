package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"gitlab.pg.innopolis.university/f.markin/fah/Fah-parsing_demo/internal/infrastructure/logger"
)

type Config struct {
	HTTPserver HTTPserver          `yaml:"server"`
	Logger     logger.LoggerConfig `yaml:"logging"`
}
type HTTPserver struct {
	Host         string        `yaml:"host"`
	Port         int           `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	IdleTimeout  time.Duration `yaml:"idle_timeout"`
}

func New() (*Config, error) {
	cfgPath := "./config/config.yaml"
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
