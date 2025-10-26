package config

import (
	"errors"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
}

type Server struct {
	Host         string        `yaml:"host"          env:"SERVER_HOST"`
	Port         string        `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	IdleTimeout  time.Duration `yaml:"idle_timeout"`
}

type Database struct {
	Host               string        `yaml:"host"                    env:"DB_HOST"`
	Port               string        `yaml:"port"`
	User               string        `yaml:"user"                    env:"POSTGRES_USER"`
	Password           string        `yaml:"password"                env:"POSTGRES_PASSWORD"`
	DatabaseName       string        `yaml:"name"`
	SSLMode            string        `yaml:"ssl_mode"`
	MaxIdleConnections int           `yaml:"max_idle_connections"`
	MaxOpenConnections int           `yaml:"max_open_connections"`
	ConnMaxLifetime    time.Duration `yaml:"connection_max_lifetime"`
	ConnTimeout        time.Duration `yaml:"connection_timeout"`
}

func MustLoadConfig(logger *zap.Logger) *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		logger.Fatal("CONFIG_PATH environment variable not set",
			zap.String("layer", logctx.LogMainFuncLayer),
			zap.String("function", logctx.LogMustLoadConfig),
			zap.Error(errors.New("empty path")),
		)
	}
	if _, errExistence := os.Stat(configPath); os.IsNotExist(errExistence) {
		logger.Fatal("CONFIG_PATH does not exist",
			zap.String("layer", logctx.LogMainFuncLayer),
			zap.String("function", logctx.LogMustLoadConfig),
			zap.Error(errors.New(configPath)),
		)
	}
	var cfg Config
	if errRead := cleanenv.ReadConfig(configPath, &cfg); errRead != nil {
		logger.Fatal("Unable to read config",
			zap.String("layer", logctx.LogMainFuncLayer),
			zap.String("function", logctx.LogMustLoadConfig),
			zap.Error(errRead))
	}
	logger.Info("Successfully loaded config",
		zap.String("layer", logctx.LogMainFuncLayer),
		zap.String("function", logctx.LogMustLoadConfig),
	)
	return &cfg
}
