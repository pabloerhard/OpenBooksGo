package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/pabloerhard/openBooksGo/internal/configs"
	"github.com/pabloerhard/openBooksGo/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func InsertOneBook(book models.Book) error {

	_, err := configs.BooksCollection.InsertOne(context.Background(), book)

	if err != nil {
		return err
	}
	return nil
}
func DeleteOneBook(bookId string) {
	id, _ := primitive.ObjectIDFromHex(bookId)
	filter := bson.M{"_id": id}

	deleted, err := configs.BooksCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(deleted)
}
func ValidateBook(googleId string) (bool, error) {
	filter := bson.M{"googleid": googleId}

	result := configs.BooksCollection.FindOne(context.Background(), filter)

	if errors.Is(result.Err(), mongo.ErrNoDocuments) {
		return false, nil
	} else if result.Err() != nil {
		return false, result.Err()
	}

	return true, nil
}
func FindBookAndCreate(googleId string) (models.Book, error) {

	filter := bson.M{"googleid": googleId}

	var existing models.Book

	book := models.Book{
		ID:       primitive.NewObjectID(),
		GoogleId: googleId,
	}

	err := configs.BooksCollection.FindOne(context.Background(), filter).Decode(&existing)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			result := InsertOneBook(book)

			if result != nil {
				return book, err
			}
		} else {
			return book, err
		}
	}

	return book, nil

}
