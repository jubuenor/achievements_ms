package db

import (
	"context"
	"fmt"
	"log"
	"os"

    "github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Database
var AchievementsCollection *mongo.Collection
var UserAchievementsCollection *mongo.Collection

func DBConnection() *mongo.Client {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("No MONGO_URI in .env file")
	}
    database := os.Getenv("MONGODB_DB")
    if database == "" {
        log.Fatal("No MONGODB_DATABASE in .env file")
    }
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

    Database = client.Database(database)
    AchievementsCollection = Database.Collection("achievements")
    UserAchievementsCollection = Database.Collection("user_achievements")

    if err := client.Ping(context.TODO(), nil); err != nil {
        panic(err)
    }

    fmt.Println("Connected to MongoDB!")
    return client
}
