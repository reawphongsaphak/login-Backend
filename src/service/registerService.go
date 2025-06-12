package service

import (
	"context"
	"fmt"
	"time"

	"main/src/database"
	"main/src/model"
	"main/src/utils"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func RegisterNewUser(user models.User) (*mongo.InsertOneResult, error){
	client := database.ConnectDB()
	
	coll := client.Database("test01").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	existingUser := coll.FindOne(ctx, bson.M{"email": user.Email})
	if existingUser.Err() == nil {
		return nil, fmt.Errorf("this email already register")
	} else if existingUser.Err() != mongo.ErrNoDocuments {
		return nil, existingUser.Err()
	}

	HashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	HashString := string(HashPassword)

	result, err := coll.InsertOne(
		ctx,
		bson.M{
			"email" : user.Email,
			"username": user.UserName,
			"password": HashString,
	})

	return result, err
}