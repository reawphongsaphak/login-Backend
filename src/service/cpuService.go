package service

import (
	"context"
	"time"

	"main/src/database"
	"main/src/model"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllCPU () ([]models.CPU, error) {
	client := database.ConnectDB()

	// Get collection
	col := client.Database("mydatabase").Collection("CPUs")

	// Context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find all documents
	cursor, err := col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var cpus []models.CPU
	for cursor.Next(ctx) {
		var cpu models.CPU
		if err := cursor.Decode(&cpu); err != nil {
			return nil, err
		}
		cpus = append(cpus, cpu)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return cpus, nil
}