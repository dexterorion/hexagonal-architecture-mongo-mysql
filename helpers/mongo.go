package helpers

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func StartMongoDb() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_CONN")))

	defer cancel()

	if err != nil {
		panic(err)
	}

	database := client.Database(os.Getenv("MONGO_DB"))
	if database == nil {
		panic(fmt.Errorf("database does not exist"))
	}

	return database
}
