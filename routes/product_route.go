package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/romanmufid16/go-mongo-redis/handler"
)

func ProductRoutes(app *fiber.App) {
	productHandler := handler.NewProductHandler()
	products := app.Group("/products")

	products.Post("/", productHandler.CreateProduct)
}
