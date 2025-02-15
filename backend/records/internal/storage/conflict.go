package storage

import (
	"context"

	"github.com/davg/records/internal/domain/models"
)

func (s *Storage) Conflicts(ctx context.Context) (*[]models.ConflictModel, error) {
	var conflict []models.ConflictModel
	if err := s.db.WithContext(ctx).Preload("Records", "published = ?", true).Find(&conflict).Error; err != nil {
		return nil, err
	}
	return &conflict, nil
}

// TODO: think about pagination there
func (s *Storage) Conflict(ctx context.Context, id string) (*models.ConflictModel, error) {
	var conflict models.ConflictModel
	if err := s.db.WithContext(ctx).Preload("Records", "published = ?", true).Where("id = ?", id).First(&conflict).Error; err != nil {
		return nil, err
	}
	return &conflict, nil
}
