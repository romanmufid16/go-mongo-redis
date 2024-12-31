package app

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/romanmufid16/go-mongo-redis/config"
	"github.com/romanmufid16/go-mongo-redis/model"
	"github.com/romanmufid16/go-mongo-redis/routes"
	"github.com/romanmufid16/go-mongo-redis/utils"
)

func ErrorMiddleware(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	errorResponse := model.BuildErrorResponse(err.Error())
	return ctx.Status(code).JSON(errorResponse)
}

func Server() *fiber.App {
	utils.LoadEnv()
	config.ConnectMongoDB()
	config.ConnectRedis()

	app := fiber.New(fiber.Config{
		ErrorHandler: ErrorMiddleware,
	})
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} ${method} ${path} ${latency}\n",
	}))

	routes.ProductRoutes(app)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	return app
}
