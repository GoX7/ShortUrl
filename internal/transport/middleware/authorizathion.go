package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gox7/shorturl/internal/crypto"
)

func NewAuth(engine *crypto.Engine) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		if header == "" {
			ctx.Next()
			return
		}

		split := strings.Split(header, " ")
		if split[0] != "Bearer" {
			ctx.Next()
			return
		}

		cipherbyte, err := engine.Open(split[1])
		if err != nil {
			ctx.Next()
			return
		}

		cipherSplit := strings.Split(string(cipherbyte), "|")
		ctx.Set("id", cipherSplit[0])
		ctx.Set("login", cipherSplit[1])
		ctx.Set("password", cipherSplit[2])

		ctx.Next()
	}
}
