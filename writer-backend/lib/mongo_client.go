package lib

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var DB *mongo.Client

func ConnectDB(mongoURI string) {
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable is not set")
	}

	client, err := mongo.Connect(options.Client().ApplyURI(mongoURI))

	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	DB = client
	log.Println("Connected to MongoDB successfully")
}

func DisconnectDB() {
	if DB != nil {
		if err := DB.Disconnect(context.Background()); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
		log.Println("Disconnected from MongoDB successfully")
		return
	}
}

func GetDB() *mongo.Client {
	if DB == nil {
		log.Fatal("MongoDB client is not initialized")
	}
	return DB
}
