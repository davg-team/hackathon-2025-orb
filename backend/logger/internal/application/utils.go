package application

import (
	"context"
	"time"

	"github.com/davg/logger/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToMongoDB() *mongo.Collection {
	cfg := config.Config().Storage
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.URI))
	if err != nil {
		panic(err)
	}
	db := client.Database(cfg.DBName)
	logsCollection := db.Collection(cfg.Collection)
	return logsCollection
}
