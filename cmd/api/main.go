package main

import (
	"boobook/internal/app/service_provider"
	"boobook/internal/config"
	"boobook/internal/http/middleware"
	"boobook/internal/repository/postgres"
	"boobook/internal/slogger"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
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
		Handler:        setupRouter(serviceProvider),
		MaxHeaderBytes: 1 << 2,
		ReadTimeout:    cfg.HTTPServerConfig.Timeout * time.Second,
		WriteTimeout:   cfg.HTTPServerConfig.Timeout * time.Second,
	}

	// Start server
	if err = httpServer.ListenAndServe(); err != nil {
		panic(fmt.Errorf("failed to start the server: %w", err))
	}
}

func setupRouter(serviceProvider *service_provider.ServiceProvider) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.Use(middleware.Log(serviceProvider.Logger))
		auth := api.Group("/auth")
		{
			auth.POST("/login", serviceProvider.GetAuthHandler().Login)
			auth.POST("/register", serviceProvider.GetAuthHandler().Register)
		}
		users := api.Group("/users")
		{
			users.Use(middleware.Auth()).GET("/:id", serviceProvider.GetUserHandler().Get)
		}
	}

	return router
}
