package handler

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"regexp"

	"github.com/gin-gonic/gin"
)

type (
	RequestAuth struct {
		Login    string `json:"login" validate:"required,alphanum"`
		Password string `json:"password" validate:"required"`
	}
)

func AuthRegister(ctx *gin.Context) {
	var request RequestAuth
	ctx.Header("Content-Type", "application/json")
	if err := ctx.BindJSON(&request); err != nil {
		ctx.Status(400)
		json.NewEncoder(ctx.Writer).Encode(ResponseError(`Register form is invalid. Usage: {"login":"...","password":"..."}`, ctx.ClientIP()))
		return
	}

	pattern := regexp.MustCompile("^[a-zA-Z0-9_.-]+$")
	if !pattern.MatchString(request.Login) || !pattern.MatchString(request.Password) {
		ctx.Status(400)
		json.NewEncoder(ctx.Writer).Encode(ResponseError(`Invalid Characters. Valide: "a-zA-Z0-9_.-"`, ctx.ClientIP()))
		return
	}

	hash := sha256.Sum256([]byte(request.Password))
	password := base64.URLEncoding.EncodeToString(hash[:])
	token, err := PostgresUsers.Register(request.Login, password, ctx.ClientIP())
	if err != nil {
		ctx.Status(400)
		json.NewEncoder(ctx.Writer).Encode(ResponseError(err.Error(), ctx.ClientIP()))
		return
	}

	ctx.Status(200)
	json.NewEncoder(ctx.Writer).Encode(ResponseOk(token, ctx.ClientIP()))
}

func AuthLogin(ctx *gin.Context) {
	var request RequestAuth
	ctx.Header("Content-Type", "application/json")
	if err := ctx.BindJSON(&request); err != nil {
		ctx.Status(400)
		json.NewEncoder(ctx.Writer).Encode(ResponseError(`Login form is invalid. Usage: {"login":"...","password":"..."}`, ctx.ClientIP()))
		return
	}

	pattern := regexp.MustCompile("^[a-zA-Z0-9_.-]+$")
	if !pattern.MatchString(request.Login) || !pattern.MatchString(request.Password) {
		ctx.Status(400)
		json.NewEncoder(ctx.Writer).Encode(ResponseError("Invalid Characters. Valide: \"a-zA-Z0-9_.-\"", ctx.ClientIP()))
		return
	}

	hash := sha256.Sum256([]byte(request.Password))
	password := base64.URLEncoding.EncodeToString(hash[:])
	token, err := PostgresUsers.Login(request.Login, password)
	if err != nil {
		ctx.Status(400)
		json.NewEncoder(ctx.Writer).Encode(ResponseError(err.Error(), ctx.ClientIP()))
		return
	}

	ctx.Status(200)
	json.NewEncoder(ctx.Writer).Encode(ResponseOk(token, ctx.ClientIP()))
}
