package configs

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

const dbName = "OpenBooksGo"
const usersColName = "users"
const booksColName = "books"

var UsersCollection *mongo.Collection
var BooksCollection *mongo.Collection

func init() {
	err := godotenv.Load("C:\\Users\\pablo\\GolandProjects\\OpenBooks\\.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	mongoApi := os.Getenv("MONGO_API")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoApi).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	fmt.Println("MongoDB connection success")
	UsersCollection = client.Database(dbName).Collection(usersColName)
	_, err = UsersCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	BooksCollection = client.Database(dbName).Collection(booksColName)
	_, err = BooksCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "googleid", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Collection instance is ready")
}
