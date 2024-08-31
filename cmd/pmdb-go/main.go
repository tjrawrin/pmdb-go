package main

import (
	"pmdb-go/internal/server"
)

func main() {
	server := server.NewFiberServer()
	server.RegisterMiddlewares()
	server.RegisterHandlers()
	err := server.Listen(":3000")
	if err != nil {
		panic("Error starting server: " + err.Error())
	}
}
