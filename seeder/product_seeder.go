package seeder

import (
	"context"
	"github.com/romanmufid16/go-mongo-redis/config"
	"github.com/romanmufid16/go-mongo-redis/model"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func SeedProducts() {
	collection := config.GetMongoCollection("products")

	// Menyiapkan data produk
	products := []model.Product{
		{Name: "Smartphone A", Price: 5000000, Category: "Electronics"},
		{Name: "Smartphone B", Price: 4500000, Category: "Electronics"},
		{Name: "Laptop X", Price: 12000000, Category: "Electronics"},
		{Name: "Laptop Y", Price: 10000000, Category: "Electronics"},
		{Name: "Headphone A", Price: 500000, Category: "Electronics"},
		{Name: "Headphone B", Price: 750000, Category: "Electronics"},
		{Name: "Smartwatch A", Price: 1500000, Category: "Electronics"},
		{Name: "Smartwatch B", Price: 1000000, Category: "Electronics"},
		{Name: "Tablet A", Price: 4500000, Category: "Electronics"},
		{Name: "Tablet B", Price: 4000000, Category: "Electronics"},
		{Name: "Monitor A", Price: 2000000, Category: "Electronics"},
		{Name: "Monitor B", Price: 2500000, Category: "Electronics"},
		{Name: "Keyboard A", Price: 350000, Category: "Electronics"},
		{Name: "Keyboard B", Price: 400000, Category: "Electronics"},
		{Name: "Mouse A", Price: 150000, Category: "Electronics"},
		{Name: "Mouse B", Price: 180000, Category: "Electronics"},
		{Name: "External Hard Drive A", Price: 800000, Category: "Electronics"},
		{Name: "External Hard Drive B", Price: 900000, Category: "Electronics"},
		{Name: "Router A", Price: 600000, Category: "Electronics"},
		{Name: "Router B", Price: 700000, Category: "Electronics"},
		{Name: "Flash Drive A", Price: 150000, Category: "Electronics"},
		{Name: "Flash Drive B", Price: 180000, Category: "Electronics"},
		{Name: "Digital Camera A", Price: 3000000, Category: "Electronics"},
		{Name: "Digital Camera B", Price: 3500000, Category: "Electronics"},
		{Name: "Speakers A", Price: 500000, Category: "Electronics"},
		{Name: "Speakers B", Price: 600000, Category: "Electronics"},
		{Name: "Electric Kettle A", Price: 250000, Category: "Home Appliances"},
		{Name: "Electric Kettle B", Price: 300000, Category: "Home Appliances"},
		{Name: "Microwave Oven A", Price: 1500000, Category: "Home Appliances"},
		{Name: "Microwave Oven B", Price: 1700000, Category: "Home Appliances"},
		{Name: "Blender A", Price: 800000, Category: "Home Appliances"},
		{Name: "Blender B", Price: 950000, Category: "Home Appliances"},
		{Name: "Washing Machine A", Price: 3500000, Category: "Home Appliances"},
		{Name: "Washing Machine B", Price: 4000000, Category: "Home Appliances"},
		{Name: "Refrigerator A", Price: 5000000, Category: "Home Appliances"},
		{Name: "Refrigerator B", Price: 6000000, Category: "Home Appliances"},
		{Name: "Air Conditioner A", Price: 3000000, Category: "Home Appliances"},
		{Name: "Air Conditioner B", Price: 3500000, Category: "Home Appliances"},
		{Name: "Vacuum Cleaner A", Price: 1000000, Category: "Home Appliances"},
		{Name: "Vacuum Cleaner B", Price: 1200000, Category: "Home Appliances"},
		{Name: "Iron A", Price: 500000, Category: "Home Appliances"},
		{Name: "Iron B", Price: 550000, Category: "Home Appliances"},
		{Name: "Hair Dryer A", Price: 300000, Category: "Home Appliances"},
		{Name: "Hair Dryer B", Price: 350000, Category: "Home Appliances"},
		{Name: "Electric Fan A", Price: 400000, Category: "Home Appliances"},
		{Name: "Electric Fan B", Price: 450000, Category: "Home Appliances"},
		{Name: "Water Heater A", Price: 1500000, Category: "Home Appliances"},
		{Name: "Water Heater B", Price: 1800000, Category: "Home Appliances"},
		{Name: "Bicycle A", Price: 2000000, Category: "Sports"},
		{Name: "Bicycle B", Price: 2500000, Category: "Sports"},
		{Name: "Treadmill A", Price: 4000000, Category: "Sports"},
		{Name: "Treadmill B", Price: 4500000, Category: "Sports"},
	}

	// Melakukan pengecekan dan insert data
	for _, product := range products {
		// Mengecek apakah produk sudah ada di koleksi berdasarkan 'Name'
		var existingProduct model.Product
		err := collection.FindOne(context.Background(), bson.D{{"name", product.Name}}).Decode(&existingProduct)

		if err != nil {
			// Jika produk belum ada (error karena tidak ditemukan), insert produk baru
			_, err := collection.InsertOne(context.Background(), product)
			if err != nil {
				log.Fatalf("Failed to insert product: %v", err)
			}
			log.Printf("Product %s inserted successfully!", product.Name)
		} else {
			log.Printf("Product %s already exists. Skipping...", product.Name)
		}
	}

	log.Println("Products seeding completed!")
}
