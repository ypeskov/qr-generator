package config

import (
	"os"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Port     string `env:"SERVER_PORT" envDefault:"80"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
}

func New() *Config {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Working directory: %s", wd)
	cfg := &Config{}
	if err := godotenv.Load(); err != nil {
		log.Warn("No .env file found")
	}
	if err := env.Parse(cfg); err != nil {
		log.Fatal(err)
	}

	cfg.LogLevel = "info"

	return cfg
}
