package storage

import (
	"context"

	"github.com/davg/records/internal/domain/models"
)

func (s *Storage) AddDocument(ctx context.Context, document *models.DocumentModel) error {
	return s.db.WithContext(ctx).Create(document).Error
}
