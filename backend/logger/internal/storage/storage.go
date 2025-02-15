package storage

import "go.mongodb.org/mongo-driver/mongo"

type Storage struct {
	collection *mongo.Collection
}

func NewStorage(collection *mongo.Collection) *Storage {
	return &Storage{
		collection: collection,
	}
}
