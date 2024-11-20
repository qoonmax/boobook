package handler

import (
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type UserHandler interface {
	Get(ctx *gin.Context)
	Search(ctx *gin.Context)
}

type PostHandler interface {
	GetList(ctx *gin.Context)
}
