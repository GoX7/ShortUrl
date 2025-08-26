package handler

import (
	"math/rand"
	"time"

	"github.com/gox7/shorturl/internal/services"
)

var (
	PostgresUsers *services.DatabaseUsers
	PostgresLinks *services.DatabaseLinks
)

type (
	Response struct {
		Status  string            `json:"status"`
		Message string            `json:"message,omitempty"`
		Error   string            `json:"error,omitempty"`
		Links   map[string]string `json:"links,omitempty"`
		Meta    Meta              `json:"meta"`
	}
	Meta struct {
		Client    string `json:"client"`
		Timestamp int64  `json:"timestamp"`
	}
)

func GenerateKey() string {
	approves := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := make([]byte, 6)
	for iter := range bytes {
		bytes[iter] = approves[rand.Intn(len(approves))]
	}
	return string(bytes)
}

func ResponseOk(message string, client string) *Response {
	return &Response{
		Status:  "ok",
		Message: message,
		Meta: Meta{
			Client:    client,
			Timestamp: time.Now().Unix(),
		},
	}
}

func ResponseLinks(links map[string]string, client string) *Response {
	return &Response{
		Status: "ok",
		Links:  links,
		Meta: Meta{
			Client:    client,
			Timestamp: time.Now().Unix(),
		},
	}
}

func ResponseError(err string, client string) *Response {
	return &Response{
		Status: "error",
		Error:  err,
		Meta: Meta{
			Client:    client,
			Timestamp: time.Now().Unix(),
		},
	}
}

func SetResource(users *services.DatabaseUsers, links *services.DatabaseLinks) {
	PostgresUsers = users
	PostgresLinks = links
}
