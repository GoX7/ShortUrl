package handler

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func StatusServer(ctx *gin.Context) {
	ctx.Status(200)
	ctx.Header("Content-Type", "application/json")
	json.NewEncoder(ctx.Writer).
		Encode(ResponseOk(
			"server is work",
			ctx.ClientIP()),
		)
}

func StatusPostgres(ctx *gin.Context) {
	ctx.Status(200)
	ctx.Header("Content-Type", "application/json")
	json.NewEncoder(ctx.Writer).
		Encode(ResponseOk(
			"postgres is work",
			ctx.ClientIP()),
		)
}
