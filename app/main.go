package main

import "github.com/Danila331/HACH-T1/app/servers"

func main() {
	err := servers.StartServer()
	if err != nil {
		panic(err)
	}
}
