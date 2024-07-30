package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"time"
)

type Config struct {
	Env              string `env:"ENV" env-default:"dev"`
	DBString         string `env:"DATABASE_DBSTRING" env-required:"true"`
	HTTPServerConfig HTTPServerConfig
}

type HTTPServerConfig struct {
	Host             string        `env:"HOST" env-default:"localhost"`
	Port             string        `env:"PORT" env-default:"8080"`
	Timeout          time.Duration `env:"TIMEOUT" env-default:"4s"`
	KeepAliveTimeout time.Duration `env:"KEEP_ALIVE_TIMEOUT" env-default:"60s"`
}

func MustLoad() *Config {
	// Загружаем переменные окружения из файла .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg Config

	// Читаем конфигурацию из файла
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	return &cfg
}
