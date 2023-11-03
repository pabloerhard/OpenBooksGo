package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pabloerhard/openBooksGo/internal/models"
	"github.com/pabloerhard/openBooksGo/internal/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func InsertUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		log.Fatal(err)
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.EncryptedPassword), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	user.EncryptedPassword = string(hash)

	errInsert := services.InsertUser(user)

	if errInsert != nil {
		c.IndentedJSON(http.StatusConflict, errInsert)
	} else {
		c.IndentedJSON(http.StatusCreated, user)
	}
}
func LoginUser(c *gin.Context) {
	var login struct {
		Email    string
		Password string
	}

	if err := c.BindJSON(&login); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
	}

	result, err := services.LoginEmail(login.Email, login.Password)

	if result == "" && err == nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "User Not Found",
		})
	} else if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", result, 3600*24*30, "", "", false, true)

	c.IndentedJSON(http.StatusOK, gin.H{
		"accessToken": result,
	})
}
func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.IndentedJSON(http.StatusOK, gin.H{
		"accessToken": user,
	})
}
func InsertWantToReadBook(c *gin.Context) {
	googleId := c.Param("googleId")
	userId := c.Param("userId")

	if googleId == "" || userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Both googleId and userId must be provided",
		})
		return
	}

	_, err := services.InsertBookToWantToRead(googleId, userId)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Error updating user",
		})
	}

	c.IndentedJSON(200, gin.H{
		"updated user": "user update successfully",
	})
}
func InsertReadBook(c *gin.Context) {
	googleId := c.Param("googleId")
	userId := c.Param("userId")

	if googleId == "" || userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Both googleId and userId must be provided",
		})
		return
	}

	_, err := services.InsertBookToRead(googleId, userId)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Error updating user",
			"error":   err,
		})
	}

	c.IndentedJSON(200, gin.H{
		"updated user": "user update successfully",
	})
}
func InsertReadingBook(c *gin.Context) {
	googleId := c.Param("googleId")
	userId := c.Param("googleId")

	if googleId == "" || userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Both googleId and userId must be provided",
		})
		return
	}

	_, err := services.InsertBookToReading(googleId, userId)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Error updating user",
			"error":   err,
		})
	}

	c.IndentedJSON(200, gin.H{
		"updated user": "user update successfully",
	})

}
func FindUser(c *gin.Context) {
	userId := c.Param("userId")

	if userId == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request: UserId field is empty",
		})
		return
	}

	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid UserId",
		})
		return
	}

	user, err := services.FindUser(objectId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": "User Not Found",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"user": user,
	})
}
