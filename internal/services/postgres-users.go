package services

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/gox7/shorturl/internal/crypto"
	"github.com/gox7/shorturl/internal/repository"
)

type (
	DatabaseUsers struct {
		connect *sql.DB
		logger  *slog.Logger
		engine  *crypto.Engine
	}
)

func NewUsers(database *Database, engine *crypto.Engine, users *DatabaseUsers) {
	repository.CreateTableUser(database.connect)
	*users = DatabaseUsers{
		connect: database.connect,
		logger:  database.logger,
		engine:  engine,
	}
}

func (users *DatabaseUsers) Register(login string, password string, client string) (string, error) {
	id, err := repository.CreateUser(
		users.connect,
		login, password,
		client,
	)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_login_key\"" {
			return "error", fmt.Errorf("%s", "Duplicate Account")
		}
		return "error", err
	}

	token := fmt.Sprintf("%d|%s|%s", id, login, password)
	return users.engine.Seal([]byte(token)), nil
}

func (users *DatabaseUsers) Login(login string, password string) (string, error) {
	id, err := repository.LoginUser(
		users.connect,
		login, password,
	)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return "error", fmt.Errorf("%s", "Account not found")
		}
		return "error", err
	}

	token := fmt.Sprintf("%d|%s|%s", id, login, password)
	return users.engine.Seal([]byte(token)), nil
}
