package main

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/gox7/shorturl/internal/crypto"
	"github.com/gox7/shorturl/internal/services"
	"github.com/gox7/shorturl/internal/transport"
	"github.com/gox7/shorturl/model"
)

func main() {
	config := new(model.Config)
	swLogger := new(slog.Logger)
	pqLogger := new(slog.Logger)

	services.NewConfig(config)
	services.NewLogger("postgres", &pqLogger)
	services.NewLogger("server", &swLogger)

	aesEngine := new(crypto.Engine)
	crypto.NewEngine(config.SERVER_PASSWORD, aesEngine)

	postgres := new(services.Database)
	postgresUsers := new(services.DatabaseUsers)
	postgresLinks := new(services.DatabaseLinks)
	services.NewPostgres(config, pqLogger, postgres)
	services.NewUsers(postgres, aesEngine, postgresUsers)
	services.NewLinks(postgres, postgresLinks)

	engine := gin.Default()
	transport.SetResource(aesEngine, swLogger, postgresUsers, postgresLinks)
	transport.Register(engine)
	transport.Run(engine, config)
}
