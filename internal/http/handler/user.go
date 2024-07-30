package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin/binding"
	_ "github.com/go-playground/validator/v10"
	"socialNetwork/internal/repository/model"
	"socialNetwork/internal/service"
)

type Handler struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) *Handler {
	return &Handler{
		UserService: userService,
	}
}

func (h *Handler) Create(ctx *gin.Context) {
	user := &model.User{}
	if err := ctx.BindJSON(user); err != nil {
		ctx.JSON(400, gin.H{"error": "invalid payload: " + err.Error()})
		return
	}
	fmt.Println(user)

	if err := h.UserService.Create(user); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "ok"})
}

func (h *Handler) Get(ctx *gin.Context) {
	panic("implement me")
}
