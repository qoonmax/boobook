package service_provider

import (
	"boobook/internal/http/handler"
	"boobook/internal/repository"
	"boobook/internal/repository/postgres"
	"boobook/internal/service"
)

var (
	postHandlerContainer    handler.PostHandler
	postServiceContainer    service.PostService
	postRepositoryContainer repository.PostRepository
)

func (sp *ServiceProvider) GetPostHandler() handler.PostHandler {
	if postHandlerContainer == nil {
		postHandlerContainer = handler.NewPostHandler(sp.Logger, sp.GetPostService(), sp.ClientRedis)
	}
	return postHandlerContainer
}

func (sp *ServiceProvider) GetPostService() service.PostService {
	if postServiceContainer == nil {
		postServiceContainer = service.NewPostService(sp.GetPostRepository())
	}
	return postServiceContainer
}

func (sp *ServiceProvider) GetPostRepository() repository.PostRepository {
	if postRepositoryContainer == nil {
		postRepositoryContainer = postgres.NewPostRepository(sp.DBReadConnection, sp.DBWriteConnection)
	}
	return postRepositoryContainer
}
