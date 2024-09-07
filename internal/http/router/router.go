package router

import (
	"boobook/internal/app/service_provider"
	"boobook/internal/http/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(serviceProvider *service_provider.ServiceProvider) *gin.Engine {
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
			users.Use(middleware.Auth()).GET("/search", serviceProvider.GetUserHandler().Search)
		}
	}

	return router
}
