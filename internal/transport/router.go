package transport

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gox7/shorturl/internal/crypto"
	"github.com/gox7/shorturl/internal/services"
	"github.com/gox7/shorturl/internal/transport/handler"
	"github.com/gox7/shorturl/internal/transport/middleware"
	"github.com/gox7/shorturl/model"
)

var (
	PostgresUsers *services.DatabaseUsers
	PostgresLinks *services.DatabaseLinks
	Crypto        *crypto.Engine
	Logger        *slog.Logger
)

func Register(engine *gin.Engine) {
	handler.SetResource(PostgresUsers, PostgresLinks)
	engine.Use(middleware.NewAuth(Crypto))
	engine.Use(middleware.NewLogger(Logger))

	engine.GET("/status/server", handler.StatusServer)
	engine.GET("/status/postgres", handler.StatusPostgres)

	engine.POST("/auth/register", handler.AuthRegister)
	engine.POST("/auth/login", handler.AuthLogin)

	engine.GET("/l", handler.LinkMy)
	engine.GET("/l/:id", handler.LinkSearch)
	engine.POST("/l", handler.LinkRegister)
}

func SetResource(engine *crypto.Engine, logger *slog.Logger, users *services.DatabaseUsers, links *services.DatabaseLinks) {
	PostgresUsers = users
	PostgresLinks = links
	Logger = logger
	Crypto = engine
}

func Run(engine *gin.Engine, config *model.Config) {
	server := http.Server{
		Handler:      engine,
		Addr:         config.SERVER_HOST + ":" + config.SERVER_PORT,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
	}

	fmt.Println("[+] server.listen...")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("[-] server.listen: " + err.Error())
	}
}
