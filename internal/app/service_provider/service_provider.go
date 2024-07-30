package service_provider

import (
	"database/sql"
)

type ServiceProvider struct {
	DBConnection *sql.DB
}

func NewServiceProvider(DBConnection *sql.DB) *ServiceProvider {
	return &ServiceProvider{
		DBConnection: DBConnection,
	}
}
