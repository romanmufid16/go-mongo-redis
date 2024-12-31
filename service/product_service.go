package service

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/romanmufid16/go-mongo-redis/config"
	"github.com/romanmufid16/go-mongo-redis/model"
	"github.com/romanmufid16/go-mongo-redis/validation"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

//var productCollection = app.GetMongoCollection("products")

type ProductService interface {
	CreateProduct(data *model.Product, ctx *fiber.Ctx) (*model.Product, error)
	GetAllProducts(ctx *fiber.Ctx) ([]*model.Product, error)
	GetProductById(id string, ctx *fiber.Ctx) (*model.Product, error)
}

type productService struct {
	productCollection *mongo.Collection
	cache             *redis.Client
}

func NewProductService() ProductService {
	return &productService{
		productCollection: config.GetMongoCollection("products"),
		cache:             config.RedisClient,
	}
}

func (service *productService) CreateProduct(data *model.Product, ctx *fiber.Ctx) (*model.Product, error) {
	if err := validation.ValidationHandler(data, validation.ProductValidation(data)); err != nil {
		errorMessage := validation.HandleValidationError(err)
		return nil, fiber.NewError(fiber.StatusBadRequest, errorMessage)
	}

	result, err := service.productCollection.InsertOne(ctx.Context(), data)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	data.ID = result.InsertedID.(primitive.ObjectID)
	return data, nil
}

func (service *productService) GetAllProducts(ctx *fiber.Ctx) ([]*model.Product, error) {
	cachedProducts, err := service.cache.Get(ctx.Context(), "products").Result()
	if err == nil && cachedProducts != "" {
		var products []*model.Product
		err := json.Unmarshal([]byte(cachedProducts), &products)
		if err != nil {
			// Jika gagal meng-decode cache, log error dan lanjutkan untuk mengambil data dari database
			log.Println("Error decoding cache:", err)
		} else {
			// Jika berhasil meng-decode cache, log dan kembalikan data
			log.Println("Data retrieved from cache.")
			return products, nil
		}
	}

	log.Println("Data not found in cache, retrieving from database...")
	cursor, err := service.productCollection.Find(ctx.Context(), bson.D{})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer cursor.Close(ctx.Context())

	var products []*model.Product

	for cursor.Next(ctx.Context()) {
		var product model.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		products = append(products, &product)
	}

	if err := cursor.Err(); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	productJSON, err := json.Marshal(products)
	if err == nil {
		service.cache.Set(ctx.Context(), "products", string(productJSON), 1*time.Minute)
	} else {
		log.Println("Error caching products", err)
	}

	log.Println("Data retrieved from database.")
	return products, nil
}

func (service *productService) GetProductById(id string, ctx *fiber.Ctx) (*model.Product, error) {
	var product model.Product

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid product ID")
	}

	err = service.productCollection.FindOne(ctx.Context(), bson.M{"_id": objectID}).Decode(&product)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fiber.NewError(fiber.StatusNotFound, "Product not found")
		}
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &product, nil
}
