package service_provider

import (
	"database/sql"
	"github.com/redis/go-redis/v9"
	"log/slog"
)

type ServiceProvider struct {
	Logger            *slog.Logger
	DBReadConnection  *sql.DB
	DBWriteConnection *sql.DB
	ClientRedis       *redis.Client
}

func NewServiceProvider(logger *slog.Logger, DBReadConnection *sql.DB, DBWriteConnection *sql.DB, clientRedis *redis.Client) *ServiceProvider {
	return &ServiceProvider{
		Logger:            logger,
		DBReadConnection:  DBReadConnection,
		DBWriteConnection: DBWriteConnection,
		ClientRedis:       clientRedis,
	}
}
