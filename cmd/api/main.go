package main

import (
	"boobook/internal/app/service_provider"
	"boobook/internal/config"
	"boobook/internal/http/router"
	"boobook/internal/repository/postgres"
	"boobook/internal/slogger"
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

func main() {
	logger := slogger.NewLogger()

	cfg := config.MustLoad()

	dbReadConnection, err := postgres.NewReadConnection(cfg.ReadDBString)
	if err != nil {
		panic(fmt.Errorf("failed to connect to the database: %w", err))
	}
	dbWriteConnection, err := postgres.NewWriteConnection(cfg.WriteDBString)
	if err != nil {
		panic(fmt.Errorf("failed to connect to the database: %w", err))
	}

	defer func(dbReadConnection *sql.DB, dbWriteConnection *sql.DB) {
		if err = postgres.CloseConnection(dbReadConnection); err != nil {
			logger.Error("failed to close the database connection", slogger.Err(err))
			return
		}
		if err = postgres.CloseConnection(dbWriteConnection); err != nil {
			logger.Error("failed to close the database connection", slogger.Err(err))
			return
		}
	}(dbReadConnection, dbWriteConnection)

	serviceProvider := service_provider.NewServiceProvider(
		logger,
		dbReadConnection,
		dbWriteConnection,
	)

	// Setup server
	httpServer := &http.Server{
		Addr:           ":" + cfg.HTTPServerConfig.Port,
		Handler:        router.SetupRouter(serviceProvider),
		MaxHeaderBytes: 1 << 2,
		ReadTimeout:    cfg.HTTPServerConfig.Timeout * time.Second,
		WriteTimeout:   cfg.HTTPServerConfig.Timeout * time.Second,
	}

	// Start server
	if err = httpServer.ListenAndServe(); err != nil {
		panic(fmt.Errorf("failed to start the server: %w", err))
	}
}
