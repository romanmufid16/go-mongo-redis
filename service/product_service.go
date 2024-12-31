package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/romanmufid16/go-mongo-redis/config"
	"github.com/romanmufid16/go-mongo-redis/model"
	"github.com/romanmufid16/go-mongo-redis/validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

//var productCollection = app.GetMongoCollection("products")

type ProductService interface {
	CreateProduct(data *model.Product, ctx *fiber.Ctx) (*model.Product, error)
}

type productService struct {
	db                *mongo.Client
	productCollection *mongo.Collection
}

func NewProductService() ProductService {
	return &productService{
		db:                config.MongoClient,
		productCollection: config.GetMongoCollection("products"),
	}
}

func (service *productService) CreateProduct(data *model.Product, ctx *fiber.Ctx) (*model.Product, error) {
	if err := validation.ValidationHandler(data, validation.ProductValidation(data)); err != nil {
		errorResponse := model.BuildErrorResponse(err.Error())
		errorMessage := strings.Join(errorResponse.Errors.([]string), ", ")
		return nil, fiber.NewError(fiber.StatusBadRequest, errorMessage)
	}

	result, err := service.productCollection.InsertOne(ctx.Context(), data)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	data.ID = result.InsertedID.(primitive.ObjectID)
	return data, nil
}
