package handler

import (
	"boobook/internal/http/request"
	"boobook/internal/repository"
	"boobook/internal/service"
	"boobook/internal/slogger"
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin/binding"
	_ "github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
)

type authHandler struct {
	logger      *slog.Logger
	authService service.AuthService
}

func NewAuthHandler(logger *slog.Logger, authService service.AuthService) AuthHandler {
	return &authHandler{
		logger:      logger,
		authService: authService,
	}
}

func (h *authHandler) Register(ctx *gin.Context) {
	var RegisterRequest request.RegisterRequest
	if err := ctx.BindJSON(&RegisterRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload: " + err.Error()})
		return
	}

	if err := h.authService.Register(&RegisterRequest); err != nil {
		if errors.Is(err, repository.ErrUserAlreadyExists) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
		} else {
			h.logger.ErrorContext(ctx, "failed to register", slogger.Err(err))
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(200, gin.H{"status": "ok"})
}

func (h *authHandler) Login(ctx *gin.Context) {
	var LoginRequest request.LoginRequest
	if err := ctx.BindJSON(&LoginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload: " + err.Error()})
		return
	}

	encryptedToken, err := h.authService.Login(&LoginRequest)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		} else if errors.Is(err, service.ErrInvalidPassword) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		} else {
			h.logger.ErrorContext(ctx, "failed to login", slogger.Err(err))
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": encryptedToken})
}
