package main

import (
	"github.com/Danila331/HACH-T1/app/servers"
	"github.com/Danila331/HACH-T1/app/store"
)

func main() {
	conn, err := store.ConnectDB()
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
