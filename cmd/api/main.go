package main

import (
	"boobook/internal/app/service_provider"
	"boobook/internal/config"
	"boobook/internal/http/router"
	"boobook/internal/repository/postgres"
	"boobook/internal/slogger"
	"context"
	"database/sql"
	"fmt"
	"github.com/redis/go-redis/v9"
	"net/http"
	"time"
)

func main() {
	logger := slogger.NewLogger()

	cfg := config.MustLoad()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: cfg.RedisPassword,
		DB:       0,
	})
	var ctx = context.Background()
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Errorf("error connecting to Redis:: %w", err))
	}
	defer func(redisClient *redis.Client) {
		if err = redisClient.Close(); err != nil {
			logger.Error("failed to close the redis connection", slogger.Err(err))
			return
		}
	}(redisClient)

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
		redisClient,
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
