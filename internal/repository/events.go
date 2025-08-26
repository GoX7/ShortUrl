package repository

import (
	"database/sql"
)

func CreateTableUser(repository *sql.DB) (sql.Result, error) {
	return repository.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		login VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		client VARCHAR(255) NOT NULL
	)`)
}

func CreateTableLinks(repository *sql.DB) (sql.Result, error) {
	return repository.Exec(`CREATE TABLE IF NOT EXISTS links (
		id SERIAL PRIMARY KEY,
		user_token VARCHAR(255) NOT NULL,
		link_original VARCHAR(255) NOT NULL UNIQUE,
		link_alias VARCHAR(255) NOT NULL UNIQUE
	)`)
}

func CreateUser(repository *sql.DB, login string, password string, client string) (int64, error) {
	var id int64
	err := repository.QueryRow("INSERT INTO users (login, password, client) VALUES ($1, $2, $3) RETURNING id",
		login,
		password,
		client,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func CreateLink(repository *sql.DB, user_token string, link_original string, link_alias string) (string, error) {
	var token string
	err := repository.QueryRow("INSERT INTO links (user_token, link_original, link_alias) VALUES ($1, $2, $3) RETURNING link_alias",
		user_token,
		link_original,
		link_alias,
	).Scan(&token)
	if err != nil {
		return token, err
	}

	return link_alias, nil
}

func LoginUser(repository *sql.DB, login string, password string) (int64, error) {
	var id int64
	err := repository.QueryRow("SELECT id FROM users WHERE login=$1 AND password=$2", login, password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func SearchLink(repository *sql.DB, alias string) (string, error) {
	var link string
	err := repository.QueryRow("SELECT link_original FROM links WHERE link_alias=$1", alias).Scan(&link)
	if err != nil {
		return "", err
	}
	return link, nil
}

func SearchLinks(repository *sql.DB, token string) (map[string]string, error) {
	links := make(map[string]string)
	query, err := repository.Query("SELECT link_original, link_alias FROM links WHERE user_token=$1", token)
	if err != nil {
		return nil, err
	}

	for query.Next() {
		var link, alias string
		query.Scan(&link, &alias)
		links[link] = alias
	}

	return links, nil
}
