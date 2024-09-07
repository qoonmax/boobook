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

	dbConnection, err := postgres.NewConnection(cfg.DBString)
	if err != nil {
		panic(fmt.Errorf("failed to connect to the database: %w", err))
	}

	defer func(dbConn *sql.DB) {
		if err = postgres.CloseConnection(dbConn); err != nil {
			logger.Error("failed to close the database connection", slogger.Err(err))
			return
		}
	}(dbConnection)

	serviceProvider := service_provider.NewServiceProvider(
		logger,
		dbConnection,
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
