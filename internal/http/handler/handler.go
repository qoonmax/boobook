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
}
