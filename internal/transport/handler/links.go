package handler

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type (
	RequestLink struct {
		Link string `json:"link"`
	}
	RequestAlias struct {
		Alias string `json:"alias"`
	}
)

func LinkSearch(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	link, err := PostgresLinks.SearchLink(ctx.Param("id"))
	if err != nil {
		ctx.Status(400)
		json.NewEncoder(ctx.Writer).Encode(ResponseError(err.Error(), ctx.ClientIP()))
		return
	}

	ctx.Status(200)
	json.NewEncoder(ctx.Writer).Encode(ResponseOk(link, ctx.ClientIP()))
}

func LinkRegister(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	if ctx.GetString("id") == "" || ctx.GetString("login") == "" {
		ctx.Status(401)
		json.NewEncoder(ctx.Writer).Encode(ResponseError("Unauthorization. Set Bearer token from /auth/register or /auth/login", ctx.ClientIP()))
		return
	}

	var request RequestLink
	if err := ctx.BindJSON(&request); err != nil {
		ctx.Status(400)
		json.NewEncoder(ctx.Writer).Encode(ResponseError(`Usage Body: {"link":"..."}`, ctx.ClientIP()))
		return
	}

	token, err := PostgresLinks.Register(
		ctx.GetString("id"),
		ctx.GetString("login"),
		request.Link, GenerateKey(),
	)
	if err != nil {
		ctx.Status(400)
		json.NewEncoder(ctx.Writer).Encode(ResponseError(err.Error(), ctx.ClientIP()))
		return
	}

	ctx.Status(200)
	json.NewEncoder(ctx.Writer).Encode(ResponseLinks(map[string]string{request.Link: token}, ctx.ClientIP()))
}

func LinkMy(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	if ctx.GetString("id") == "" || ctx.GetString("login") == "" {
		ctx.Status(401)
		json.NewEncoder(ctx.Writer).Encode(ResponseError("Unauthorization. Set Bearer token from /auth/register or /auth/login", ctx.ClientIP()))
		return
	}

	links, err := PostgresLinks.SearchLinks(
		ctx.GetString("id"),
		ctx.GetString("login"),
	)
	if err != nil {
		ctx.Status(400)
		json.NewEncoder(ctx.Writer).Encode(ResponseError(err.Error(), ctx.ClientIP()))
		return
	}

	ctx.Status(200)
	json.NewEncoder(ctx.Writer).Encode(ResponseLinks(links, ctx.ClientIP()))
}
