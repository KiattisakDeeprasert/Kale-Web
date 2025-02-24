package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	
	uri := "mongodb+srv://nawinwin88:Nawinwin46@cluster0.xdqtr.mongodb.net/"

	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("❌ Failed to connect MongoDB:", err)
	}

	
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("❌ MongoDB Ping Failed:", err)
	}

	log.Println("✅ Connected to MongoDB!")

	
	DB = client.Database("auth_db")
}
