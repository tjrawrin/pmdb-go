package server

import (
	"pmdb-go/internal/database"

	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	*fiber.App
	DB *database.Database
}

func NewFiberServer() *FiberServer {
	return &FiberServer{
		App: fiber.New(fiber.Config{
			AppName: "pmdb-go",
		}),
		DB: database.New(),
	}
}
