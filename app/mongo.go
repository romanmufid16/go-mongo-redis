package app

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var MongoClient *mongo.Client

func ConnectMongoDB() {
	clientOptions := options.Client().ApplyURI(os.Getenv("DATABASE_URI"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("MongoDB not reachable: %v", err)
	}

	log.Println("Connected to MongoDB!")
	MongoClient = client
}

func GetMongoCollection(database, collection string) *mongo.Collection {
	return MongoClient.Database(database).Collection(collection)
}
