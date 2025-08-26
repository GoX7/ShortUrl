package services

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/gox7/shorturl/internal/repository"
)

type (
	DatabaseLinks struct {
		connect *sql.DB
		logger  *slog.Logger
	}
)

func NewLinks(database *Database, links *DatabaseLinks) {
	repository.CreateTableLinks(database.connect)
	*links = DatabaseLinks{
		connect: database.connect,
		logger:  database.logger,
	}
}

func (links *DatabaseLinks) Register(user_id string, user_login string, link string, alias string) (string, error) {
	_, err := repository.SearchLink(links.connect, alias)
	if err == nil {
		return "", fmt.Errorf("%s", "Invalid Alias, try again")
	}

	token, err := repository.CreateLink(
		links.connect,
		fmt.Sprintf("%s:%s", user_id, user_login),
		link, alias,
	)
	if err.Error() == "pq: duplicate key value violates unique constraint \"links_link_original_key\"" {
		return token, fmt.Errorf("%s", "Link is exists")
	}

	return alias, err
}

func (links *DatabaseLinks) SearchLinks(user_id string, user_login string) (map[string]string, error) {
	return repository.SearchLinks(
		links.connect,
		fmt.Sprintf("%s:%s", user_id, user_login),
	)
}

func (links *DatabaseLinks) SearchLink(alias string) (string, error) {
	return repository.SearchLink(links.connect, alias)
}
