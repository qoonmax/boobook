package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"socialNetwork/internal/app/service_provider"
	"socialNetwork/internal/config"
	"socialNetwork/internal/repository/postgres"
	"time"
)

func main() {
	cfg := config.MustLoad()
	//log := setupLogger(cfg.Env)
	//log.Info("starting server")

	// Create database connection
	dbConnection, err := postgres.NewConnection(cfg.DBString)
	if err != nil {
		log.Fatal("failed to connect to the database", err)
	}

	defer func(dbConn *sql.DB) {
		if err := postgres.CloseConnection(dbConn); err != nil {
			log.Fatal("failed to close the database connection", err)
		}
	}(dbConnection)

	// Setup service provider
	serviceProvider := service_provider.NewServiceProvider(dbConnection)

	// Setup server
	httpServer := &http.Server{
		Addr:           ":" + cfg.HTTPServerConfig.Port,
		Handler:        setupHandler(serviceProvider),
		MaxHeaderBytes: 1 << 2,
		ReadTimeout:    cfg.HTTPServerConfig.Timeout * time.Second,
		WriteTimeout:   cfg.HTTPServerConfig.Timeout * time.Second,
	}

	// Start server
	if err = httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

//
//const (
//	localEnv = "local"
//	devEnv   = "dev"
//	prodEnv  = "prod"
//)
//
//func setupLogger(env string) *slog.Logger {
//	var log *slog.Logger
//
//	switch env {
//	case localEnv:
//		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
//	case devEnv:
//		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
//	case prodEnv:
//		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
//	}
//
//	return log
//}

func setupHandler(serviceProvider *service_provider.ServiceProvider) *gin.Engine {
	h := gin.Default()

	api := h.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("/:id", serviceProvider.GetUserHandler().Get)
			users.POST("/", serviceProvider.GetUserHandler().Create)
		}
	}

	return h
}
