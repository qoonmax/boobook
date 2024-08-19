package service_provider

import (
	"boobook/internal/http/handler"
	"boobook/internal/repository"
	"boobook/internal/repository/postgres"
	"boobook/internal/service"
)

var (
	userHandlerContainer    handler.UserHandler
	userServiceContainer    service.UserService
	userRepositoryContainer repository.UserRepository
)

func (sp *ServiceProvider) GetUserHandler() handler.UserHandler {
	if userHandlerContainer == nil {
		userHandlerContainer = handler.NewUserHandler(sp.Logger, sp.GetUserService())
	}
	return userHandlerContainer
}

func (sp *ServiceProvider) GetUserService() service.UserService {
	if userServiceContainer == nil {
		userServiceContainer = service.NewUserService(sp.GetUserRepository())
	}
	return userServiceContainer
}

func (sp *ServiceProvider) GetUserRepository() repository.UserRepository {
	if userRepositoryContainer == nil {
		userRepositoryContainer = postgres.NewUserRepository(sp.DBConnection)
	}
	return userRepositoryContainer
}
