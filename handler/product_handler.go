package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/romanmufid16/go-mongo-redis/model"
	"github.com/romanmufid16/go-mongo-redis/service"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{
		productService: service.NewProductService(),
	}
}

func (h *ProductHandler) CreateProduct(ctx *fiber.Ctx) error {
	var data *model.Product
	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.BuildErrorResponse("Invalid input data"))
	}

	result, err := h.productService.CreateProduct(data, ctx)
	if err != nil {
		return err
	}

	response := model.BuildResponse("Product data created", result)
	return ctx.Status(fiber.StatusCreated).JSON(response)
}
