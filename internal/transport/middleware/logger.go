package middleware

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func NewLogger(logger *slog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		info := logger.With(
			slog.String("method", ctx.Request.Method),
			slog.String("path", ctx.Request.URL.Path),
			slog.String("client", ctx.ClientIP()),
		)

		ctx.Next()
		info.Info(
			"request",
			slog.Int("status", ctx.Writer.Status()),
			slog.Int("size", ctx.Writer.Size()),
		)
	}
}
