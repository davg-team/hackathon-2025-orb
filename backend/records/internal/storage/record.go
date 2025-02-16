package storage

import (
	"context"
	"strconv"

	"github.com/davg/records/internal/domain/models"
)

func (s *Storage) Record(ctx context.Context, id string) (*models.RecordModel, error) {
	var record models.RecordModel
	if err := s.db.WithContext(ctx).Preload("Documents").Preload("Conflicts").Where("id = ?", id).First(&record).Error; err != nil {
		return nil, err
	}
	return &record, nil
}

func (s *Storage) RecordsByParams(
	ctx context.Context,
	name string,
	middleName string,
	lastName string,
	limit, offset int,
	published string,
	userID string,
) (*[]models.RecordModel, int, error) {
	var record []models.RecordModel
	query := s.db.WithContext(ctx).Preload("Documents").Preload("Conflicts")

	if limit != -1 {
		query = query.Limit(limit)
	}
	if offset != -1 {
		query = query.Offset(offset)
	}

	if name != "" {
		query = query.Where("name = ?", name)
	}
	if middleName != "" {
		query = query.Where("middle_name = ?", middleName)
	}
	if lastName != "" {
		query = query.Where("last_name = ?", lastName)
	}

	if published != "" {
		publishedBool, _ := strconv.ParseBool(published)
		query = query.Where("published = ?", publishedBool)
	}

	if userID != "" {
		query = query.Where("user_id = ?", userID)
	}

	if err := query.Find(&record).Error; err != nil {
		return nil, 0, err
	}

	var count int64
	if err := s.db.WithContext(ctx).Model(&models.RecordModel{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return &record, int(count), nil
}

func (s *Storage) Records(ctx context.Context, limit, offset int) (*[]models.RecordModel, error) {
	var record []models.RecordModel
	if err := s.db.WithContext(ctx).Preload("Documents").Limit(limit).Offset(offset).Preload("Conflicts").Find(&record).Error; err != nil {
		return nil, err
	}
	return &record, nil
}

func (s *Storage) CreateRecord(ctx context.Context, record *models.RecordModel) error {
	return s.db.WithContext(ctx).Create(record).Error
}

func (s *Storage) UpdateRecord(ctx context.Context, record *models.RecordModel) error {
	return s.db.WithContext(ctx).Save(record).Error
}

func (s *Storage) PublishRecord(ctx context.Context, recordID string) error {
	return s.db.WithContext(ctx).Model(&models.RecordModel{}).Where("id = ?", recordID).Update("published", true).Error
}

func (s *Storage) UpdateRecordMapID(ctx context.Context, recordID string, mapID int) error {
	return s.db.WithContext(ctx).Model(&models.RecordModel{}).Where("id = ?", recordID).Update("map_id", mapID).Error
}
