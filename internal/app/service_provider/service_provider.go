package service_provider

import (
	"database/sql"
	"log/slog"
)

type ServiceProvider struct {
	Logger       *slog.Logger
	DBConnection *sql.DB
}

func NewServiceProvider(logger *slog.Logger, DBConnection *sql.DB) *ServiceProvider {
	return &ServiceProvider{
		Logger:       logger,
		DBConnection: DBConnection,
	}
}
