package mongodb

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func NewMongoDBConnection() (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client().
		ApplyURI(os.Getenv("MONGODB_URL")).
		SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client, nil
}
