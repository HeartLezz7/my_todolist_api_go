package database

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type databaseUrl struct {
	url string
}

func MongoDB() *mongo.Database {
	database := databaseUrl{url: os.Getenv("MONGO_URL")}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	option := options.Client().ApplyURI(database.url)

	client, err := mongo.Connect(ctx, option)
	if err != nil {
		err = client.Disconnect(ctx)
		panic(err)
	}

	return client.Database(os.Getenv("DATABASE_NAME"))
}
