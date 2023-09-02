package main

import (
	"L0/config"
	"L0/internal/handlers"
	"L0/internal/repository"
	"L0/internal/service"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	handlers   *handlers.Handlers
	repository *repository.Repository
	routers    *fiber.App
	service    *service.Service
}

func main() {
	app := &App{}
	cfg := config.Config_load()
	app.repository = repository.New(cfg)
	app.service = service.New(app.repository)
	app.handlers = handlers.New(app.repository)
	//database.ConnectDB(cfg)
	app.routers = fiber.New()
	go func() {
		err := app.repository.CacheRecovery()
		if err != nil {
			return
		}
		for 1 == 1 {
			err = app.service.ConsumeMessage()
		}
	}()
	app.routers.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("And the API is UP!")
		return err
	})
	app.routers.Get("/:Id", app.handlers.Get)
	err := app.routers.Listen(":3000")
	if err != nil {
		return
	}
}
