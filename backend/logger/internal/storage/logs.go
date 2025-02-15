package storage

import (
	"context"

	"github.com/davg/logger/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Storage) Log(ctx context.Context, id string) ([]domain.LogModel, error) {
	var logs []domain.LogModel
	cursor, err := s.collection.Find(ctx, bson.M{"user_id": id})
	if err != nil {
		return logs, err
	}
	if err = cursor.All(ctx, &logs); err != nil {
		return logs, err
	}
	return logs, nil
}

func (s *Storage) Logs(ctx context.Context) ([]domain.LogModel, error) {
	var logs []domain.LogModel
	cursor, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return logs, err
	}
	if err = cursor.All(ctx, &logs); err != nil {
		return logs, err
	}
	return logs, nil
}

func (s *Storage) CreateLog(ctx context.Context, log domain.LogModel) (string, error) {
	_, err := s.collection.InsertOne(ctx, log)
	if err != nil {
		return "", err
	}
	return log.ID, nil
}
