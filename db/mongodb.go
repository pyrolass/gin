// db.go
package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func init() {
	// Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func ConnectDB() *mongo.Client {
	// Replace the uri string with your MongoDB Atlas connection string
	uri := os.Getenv("MONGO_URL")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB Atlas")
	DB = client

	ensureIndexes(DB)

	return client
}

func ensureIndexes(client *mongo.Client) {
	// Select your database and collection.
	collection := client.Database("books").Collection("users")

	// Define the index model for the `email` field.
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1}, // index in ascending order
		Options: options.Index().SetUnique(true),
	}

	// Create the index.
	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Unique index on `email` field ensured.")
}

func GetDBCollection(collectionName string) *mongo.Collection {
	// Replace "your_database" with your database name

	return DB.Database("books").Collection(collectionName)

}
