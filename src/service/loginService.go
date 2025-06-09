package service

import (
	"context"
	"fmt"
	"time"
	"os"

	"main/src/database"
	"main/src/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
)

func loginUser(user models.UserLogin) (string, error) {
	client := database.ConnectDB()
	coll := client.Database("test01").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Try to find the user by email
	var foundUser models.User
	err := coll.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", fmt.Errorf("user not found, please register first")
		}
		return "", fmt.Errorf("database error: %v", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err != nil {
		return "", fmt.Errorf("incorrect password")
	}

	// Generate JWT token
	claims := jwt.MapClaims{
		"username": foundUser.UserName,     // include user ID or any unique field
		"email":   foundUser.Email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(os.Getenv("SECRET_KEY")) // Load secret from env variable
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return tokenString, nil

}