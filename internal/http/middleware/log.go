package middleware

import (
	"boobook/internal/http/contextkey"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
	"time"
)

func Log(logger *slog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()

		ctx.Set(string(contextkey.CtxKeyRequestID), uuid.New().String())

		logger = logger.With(
			slog.String(string(contextkey.CtxKeyRequestID), ctx.GetString(string(contextkey.CtxKeyRequestID))),
			slog.String("method", ctx.Request.Method),
			slog.String("path", ctx.Request.URL.Path),
		)

		logger.Info("request started")

		defer func() {
			logger = logger.With(
				slog.Int("status", ctx.Writer.Status()),
				slog.String("latency", fmt.Sprintf("%v ms", time.Since(startTime).Milliseconds())),
			)
			logger.Info("request finished")
		}()

		ctx.Next()
	}
}
