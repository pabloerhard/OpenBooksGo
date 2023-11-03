package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pabloerhard/openBooksGo/internal/configs"
	"github.com/pabloerhard/openBooksGo/internal/models"
	"github.com/pabloerhard/openBooksGo/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

func InsertUser(user models.User) error {

	inserted, err := configs.UsersCollection.InsertOne(context.Background(), user)
	utils.HasError(err)

	if err != nil {
		return err
	}
	fmt.Println(inserted)
	return nil

}
func FindUser(userId primitive.ObjectID) (models.User, error) {
	filter := bson.M{"_id": userId}
	var user models.User
	err := configs.UsersCollection.FindOne(context.Background(), filter).Decode(&user)

	if err != nil {
		log.Fatal(err)
		return user, err
	}

	return user, nil
}
func LoginEmail(email string, password string) (string, error) {

	filter := bson.M{
		"email": email,
	}

	result := configs.UsersCollection.FindOne(context.Background(), filter)

	if errors.Is(result.Err(), mongo.ErrNoDocuments) {
		return "", nil
	} else if result.Err() != nil {
		return "", result.Err()
	}

	var user models.User

	if err := result.Decode(&user); err != nil {
		return "", err
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(password)); err != nil {
			return "", err
		}
	}

	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func InsertBookToWantToRead(googleId string, userId string) (models.User, error) {
	var newUser models.User
	id, _ := primitive.ObjectIDFromHex(userId)
	user, err := FindUser(id)

	if err != nil {
		log.Fatal(err)
	}

	newBook, err := FindBookAndCreate(googleId)

	if err != nil {
		return user, err
	}

	user.WantToReadBooks = append(user.WantToReadBooks, newBook.GoogleId)

	newUser, err = UpdateUser(user)

	if err != nil {
		log.Fatal(err)
	}

	return newUser, nil

}
func InsertBookToRead(googleId string, userId string) (models.User, error) {
	var newUser models.User
	id, _ := primitive.ObjectIDFromHex(userId)
	user, err := FindUser(id)

	if err != nil {
		log.Fatal(err)
	}

	newBook, err := FindBookAndCreate(googleId)

	if err != nil {
		return user, err
	}

	user.ReadBooks = append(user.ReadBooks, newBook.GoogleId)

	newUser, err = UpdateUser(user)

	if err != nil {
		log.Fatal(err)
	}

	return newUser, nil
}
func InsertBookToReading(googleId string, userId string) (models.User, error) {
	var newUser models.User
	id, _ := primitive.ObjectIDFromHex(userId)
	user, err := FindUser(id)

	if err != nil {
		log.Fatal(err)
	}

	newBook, err := FindBookAndCreate(googleId)

	if err != nil {
		return user, err
	}

	user.ReadingBooks = append(user.ReadingBooks, newBook.GoogleId)

	newUser, err = UpdateUser(user)

	if err != nil {
		log.Fatal(err)
	}

	return newUser, nil
}
func UpdateUser(user models.User) (models.User, error) {
	var newUser models.User
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": user}

	err := configs.UsersCollection.FindOneAndUpdate(context.Background(), filter, update).Decode(&newUser)

	if err != nil {
		return user, err
	}

	return newUser, nil
}
