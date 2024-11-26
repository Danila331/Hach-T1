package main

import (
	"net/url"

	"github.com/Danila331/HACH-T1/app/servers"
	"github.com/Danila331/HACH-T1/app/store"
)

func main() {
	password := url.QueryEscape("g!AVY93W<$}d&x")
	conn, err := store.ConnectToPostgresSql("45.10.43.153", "5432", "gen_user", password, "default_db")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	_, _ = conn.Exec("CREATE TABLE IF NOT EXISTS files (id SERIAL PRIMARY KEY, name TEXT, path TEXT)")
	err = servers.StartServer()
	if err != nil {
		panic(err)
	}
}
