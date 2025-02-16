package storage

import (
	"context"

	"github.com/davg/drafts/internal/domain"
)

func (s *Storage) DraftDocument(ctx context.Context, id string) (*domain.OfferDocumentUpdate, error) {
	var offer domain.OfferDocumentUpdate

	if err := s.db.Where("id = ?", id).First(&offer).Error; err != nil {
		return nil, err
	}

	return &offer, nil
}

func (s *Storage) DraftRecord(ctx context.Context, id string) (*domain.OfferRecordUpdate, error) {
	var offer domain.OfferRecordUpdate

	if err := s.db.Where("id = ?", id).First(&offer).Error; err != nil {
		return nil, err
	}

	return &offer, nil
}

func (s *Storage) DraftsDocuments(ctx context.Context) (*[]domain.OfferDocumentUpdate, error) {
	var offers []domain.OfferDocumentUpdate

	if err := s.db.Find(&offers).Error; err != nil {
		return nil, err
	}

	return &offers, nil
}

func (s *Storage) DraftsRecords(ctx context.Context) (*[]domain.OfferRecordUpdate, error) {
	var offers []domain.OfferRecordUpdate

	if err := s.db.Find(&offers).Error; err != nil {
		return nil, err
	}

	return &offers, nil
}

func (s *Storage) CreateDraftDocument(ctx context.Context, draft *domain.OfferDocumentUpdate) error {
	return s.db.Create(draft).Error
}

func (s *Storage) CreateDraftRecord(ctx context.Context, draft *domain.OfferRecordUpdate) error {
	return s.db.Create(draft).Error
}

func (s *Storage) DeleteDraftDocument(ctx context.Context, id string) error {
	return s.db.Where("id = ?", id).Delete(&domain.OfferDocumentUpdate{}).Error
}

func (s *Storage) DeleteDraftRecord(ctx context.Context, id string) error {
	return s.db.Where("id = ?", id).Delete(&domain.OfferRecordUpdate{}).Error
}
