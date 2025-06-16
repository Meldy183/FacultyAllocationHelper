package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTPserver HTTPserver `yaml:"http_server"`
	Database   Database   `yaml:"database"`
	JWT        JWT        `yaml:"jwt"`
	Cookies    Cookies    `yaml:"cookies"`
	Sequrity   Sequrity   `yaml:"security"`
}
type HTTPserver struct {
	Host         string        `yaml:"host"`
	Port         int           `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	IdleTimeout  time.Duration `yaml:"idle_timeout"`
}
type Token struct {
	Secret    string        `yaml:"secret"`
	Expiry    time.Duration `yaml:"expiry"`
	Algorithm string        `yaml:"algorithm"`
}

type JWT struct {
	AccessToken  Token `yaml:"access_token"`
	RefreshToken Token `yaml:"refresh_token"`
}
type Cookies struct {
	AccessToken  TokenCookies `yaml:"access_token"`
	RefreshToken TokenCookies `yaml:"refresh_token"`
}
type TokenCookies struct {
	Name     string `yaml:"name"`
	HTTPOnly bool   `yaml:"http_only"`
	Secure   bool   `yaml:"secure"`
	SameSite string `yaml:"same_site"`
	Path     string `yaml:"path"`
	MaxAge   int    `yaml:"max_age"`
}
type Sequrity struct {
	Bcrypt    Bcrypt    `yaml:"bcrypt"`
	RateLimit RateLimit `yaml:"rate_limit"`
	CORS      CORS      `yaml:"cors"`
}
type Bcrypt struct {
	Cost int `yaml:"cost"`
}
type RateLimit struct {
	Enabled   bool `yaml:"enabled"`
	ReqPerMin int  `yaml:"requests_per_minute"`
	Burst     int  `yaml:"burst"`
}
type CORS struct {
	AllowedOrigins    []string `yaml:"allowed_origins"`
	AllowedHeaders    []string `yaml:"allowed_headers"`
	Allow_Credentials bool     `yaml:"allow_credentials"`
	MaxAge            int      `yaml:"max_age"`
}

type Database struct {
	Host         string        `yaml:"host"`
	Port         int           `yaml:"port"`
	Name         string        `yaml:"name"`
	User         string        `yaml:"user"`
	Password     string        `yaml:"password"`
	SSLMode      string        `yaml:"ssl_mode"`
	MaxOpenConns int           `yaml:"max_open_connections"`
	MaxIdleConns int           `yaml:"max_idle_connections"`
	ConnLifetime time.Duration `yaml:"connection_max_lifetime"`
	ConnTimeOut  time.Duration `yaml:"connection_timeout"`
}

func New() (*Config, error) {
	cfgPath := "config/config.yaml"
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
