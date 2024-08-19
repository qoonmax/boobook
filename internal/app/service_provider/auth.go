package service_provider

import (
	"boobook/internal/http/handler"
	"boobook/internal/service"
)

var (
	authHandlerContainer handler.AuthHandler
	authServiceContainer service.AuthService
)

func (sp *ServiceProvider) GetAuthHandler() handler.AuthHandler {
	if authHandlerContainer == nil {
		authHandlerContainer = handler.NewAuthHandler(sp.Logger, sp.GetAuthService())
	}
	return authHandlerContainer
}

func (sp *ServiceProvider) GetAuthService() service.AuthService {
	if authServiceContainer == nil {
		authServiceContainer = service.NewAuthService(sp.GetUserRepository())
	}
	return authServiceContainer
}
