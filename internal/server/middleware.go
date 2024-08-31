package server

import (
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func (s *FiberServer) RegisterMiddlewares() {
	s.App.Use(logger.New(logger.Config{
		Format: "${time} | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${queryParams} | ${error}\n",
	}))
	s.App.Use(healthcheck.New())
	s.App.Use(recover.New())
}
