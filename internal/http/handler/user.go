package handler

import (
	"boobook/internal/repository"
	"boobook/internal/service"
	"boobook/internal/slogger"
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin/binding"
	_ "github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
	"strconv"
)

type userHandler struct {
	logger      *slog.Logger
	userService service.UserService
}

func NewUserHandler(logger *slog.Logger, userService service.UserService) UserHandler {
	return &userHandler{
		logger:      logger,
		userService: userService,
	}
}

func (h *userHandler) Get(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	user, err := h.userService.Get(uint(id))
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		} else {
			h.logger.ErrorContext(ctx, "failed to get user", slogger.Err(err))
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
