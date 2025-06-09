package database

import (
	"fmt"
	"log"
	"os"
	"sync"
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var clientInstance *mongo.Client
var clientOnce sync.Once

func ConnectDB() *mongo.Client {
	clientOnce.Do(func() {
		uri := os.Getenv("MONGODB_URI")
		if uri == "" {
			log.Fatal("You must set your 'MONGODB_URI' environment variable.")
		}

		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

		var errConn error
		clientInstance, errConn = mongo.Connect(opts)
		if errConn != nil {
			log.Fatalf("MongoDB connection error: %v", errConn)
		}

		if err := clientInstance.Ping(context.TODO(), readpref.Primary()); err != nil {
            log.Fatalf("MongoDB ping error: %v", err)
        }

		fmt.Println("Connected to MongoDB successfully!")
	})
	return clientInstance
}
