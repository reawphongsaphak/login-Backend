package service

import (
	"context"
	"fmt"
	"time"

	"main/src/database"
	"main/src/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func AddNewUser(user models.User) (*mongo.InsertOneResult, error){
	client := database.ConnectDB()
	
	coll := client.Database("test01").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	existingUser := coll.FindOne(ctx, bson.M{"username": user.UserName})
	if existingUser.Err() == nil {
		return nil, fmt.Errorf("user with username '%s' already exists", user.UserName)
	} else if existingUser.Err() != mongo.ErrNoDocuments {
		return nil, existingUser.Err()
	}

	result, err := coll.InsertOne(
		ctx,
		bson.M{
			"username": user.UserName,
			"password": user.Password,
	})

	return result, err
}