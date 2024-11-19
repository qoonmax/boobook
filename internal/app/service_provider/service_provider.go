package service_provider

import (
	"database/sql"
	"log/slog"
)

type ServiceProvider struct {
	Logger            *slog.Logger
	DBReadConnection  *sql.DB
	DBWriteConnection *sql.DB
}

func NewServiceProvider(logger *slog.Logger, DBReadConnection *sql.DB, DBWriteConnection *sql.DB) *ServiceProvider {
	return &ServiceProvider{
		Logger:            logger,
		DBReadConnection:  DBReadConnection,
		DBWriteConnection: DBWriteConnection,
	}
}
