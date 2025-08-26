package services

import (
	"database/sql"
	"log/slog"
	"os"

	"github.com/gox7/shorturl/internal/repository"
	"github.com/gox7/shorturl/model"
)

type (
	Database struct {
		connect *sql.DB
		logger  *slog.Logger
	}
)

func NewPostgres(config *model.Config, logger *slog.Logger, database *Database) {
	connect, err := repository.NewConnect(
		config.POSTGRES_HOST,
		config.POSTGRES_PORT,
		config.POSTGRES_USER,
		config.POSTGRES_PASSWORD,
		config.POSTGRES_NAME,
	)

	if err != nil {
		logger.Error("postgres.connect", slog.String("error", err.Error()))
		os.Exit(1)
	}

	*database = Database{
		connect: connect,
		logger:  logger,
	}
}
