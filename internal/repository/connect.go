package repository

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func NewConnect(host string, port string, user string, password string, dbname string) (*sql.DB, error) {
	psc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port,
		user, password,
		dbname, "disable")

	var connect *sql.DB
	var err error
	for i := 0; i < 10; i++ {
		connect, err = sql.Open("postgres", psc)
		if err == nil {
			fmt.Println("[+] postgres.connect: ok")
			break
		}
		time.Sleep(400 * time.Millisecond)
	}
	if err != nil {
		fmt.Println("[-] postgres.connect: " + err.Error())
		return nil, err
	}

	for i := 0; i < 10; i++ {
		err = connect.Ping()
		if err == nil {
			fmt.Println("[+] postgres.ping: ok")
			break
		}
		time.Sleep(400 * time.Millisecond)
	}
	if err != nil {
		fmt.Println("[-] postgres.ping: " + err.Error())
		return nil, err
	}

	return connect, nil
}
