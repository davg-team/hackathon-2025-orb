package record

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/davg/records/internal/customerrors"
	"github.com/davg/records/internal/domain/models"
	"github.com/davg/records/internal/domain/requests"
	authorization "github.com/davg/records/pkg/middlewares"
	"github.com/google/uuid"
)

type RecordStorage interface {
	Record(ctx context.Context, id string) (*models.RecordModel, error)
	Records(ctx context.Context, limit, offset int) (*[]models.RecordModel, error)
	RecordsByParams(
		ctx context.Context,
		name string,
		middleName string,
		lastName string,
		limit, offset int,
		published, userID string,
	) (*[]models.RecordModel, int, error)
	CreateRecord(ctx context.Context, record *models.RecordModel) error
	UpdateRecord(ctx context.Context, record *models.RecordModel) error
	PublishRecord(ctx context.Context, recordID string) error

	AddDocument(ctx context.Context, document *models.DocumentModel) error
	Conflict(ctx context.Context, id string) (*models.ConflictModel, error)
}

type RecordService struct {
	log     *slog.Logger
	storage RecordStorage
}

func NewRecordService(log *slog.Logger, storage RecordStorage) *RecordService {
	log = log.With("service", "record")
	return &RecordService{
		log:     log,
		storage: storage,
	}
}

func (s *RecordService) Record(ctx context.Context, id string) (*models.RecordModel, error) {
	const op = "recordService.Record"
	log := s.log.With("op", op)

	log.Info("Get record by id", "id", id)

	record, err := s.storage.Record(ctx, id)
	if err != nil {
		log.Error("failed to get record", "error", err)
		return nil, fmt.Errorf("%w: %s", customerrors.ErrNotFound, "record by id not found")
	}

	return record, nil
}

func (s *RecordService) Records(ctx context.Context, limit, offset string) (*[]models.RecordModel, int, error) {
	const op = "recordService.Records"
	log := s.log.With("op", op)

	log.Info("Get all records")

	if limit == "" {
		limit = "-1"
	}
	if offset == "" {
		offset = "-1"
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		log.Error("failed to convert limit to int", "error", err)
		return nil, 0, fmt.Errorf("%w: %s", customerrors.ErrBadRequest, "failed to convert limit to int")
	}
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		log.Error("failed to convert offset to int", "error", err)
		return nil, 0, fmt.Errorf("%w: %s", customerrors.ErrBadRequest, "failed to convert offset to int")
	}

	records, err := s.storage.Records(ctx, limitInt, offsetInt)
	if err != nil {
		log.Error("failed to get records", "error", err)
		return nil, 0, fmt.Errorf("%w: %s", customerrors.ErrInternal, "failed to get records")
	}

	return records, len(*records), nil
}

func (s *RecordService) RecordsByParams(
	ctx context.Context,
	name string,
	middleName string,
	lastName string,
	limit, offset string,
	published, userID string,
) (*[]models.RecordModel, int, error) {
	const op = "recordService.RecordsByParams"
	log := s.log.With("op", op)

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		log.Error("failed to convert limit to int", "error", err)
		return nil, 0, fmt.Errorf("%w: %s", customerrors.ErrBadRequest, "failed to convert limit to int")
	}
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		log.Error("failed to convert offset to int", "error", err)
		return nil, 0, fmt.Errorf("%w: %s", customerrors.ErrBadRequest, "failed to convert offset to int")
	}

	log.Info("Get records by params", "name", name, "middleName", middleName, "lastName", lastName)

	records, totalCount, err := s.storage.RecordsByParams(ctx, name, middleName, lastName, limitInt, offsetInt, published, userID)
	if err != nil {
		log.Error("failed to get records", "error", err)
		return nil, 0, fmt.Errorf("%w: %s", customerrors.ErrInternal, "failed to get records")
	}

	return records, totalCount, nil
}

func (s *RecordService) CreateRecord(ctx context.Context, record *requests.RecordRequest) (string, error) {
	const op = "recordService.CreateRecord"
	log := s.log.With("op", op)

	log.Info("Create record", "record", record)

	recordID := uuid.New().String()
	var documents []models.DocumentModel
	for _, doc := range record.Documents {
		documents = append(documents, models.DocumentModel{
			ID:       uuid.New().String(),
			Type:     doc.Type,
			URL:      doc.URL,
			RecordID: recordID,
		})
	}

	var conflicts []models.ConflictModel
	for _, conflictID := range record.Conflicts {
		conflict, err := s.storage.Conflict(ctx, conflictID)
		if err != nil {
			log.Error("failed to get conflict", "error", err)
			return "", fmt.Errorf("%w: %s", customerrors.ErrBadRequest, "failed to get conflict")
		}
		conflicts = append(conflicts, *conflict)
	}

	recordModel := &models.RecordModel{
		ID:           recordID,
		Name:         record.Name,
		MiddleName:   record.MiddleName,
		LastName:     record.LastName,
		BirthDate:    record.BirthDate,
		BirthPlace:   record.BirthPlace,
		MilitaryRank: record.MilitaryRank,
		Commissariat: record.Commissariat,
		Awards:       record.Awards,
		DeathDate:    record.DeathDate,
		BurialPlace:  record.BurialPlace,
		Bio:          record.Bio,
		Conflicts:    conflicts,
		Published:    false,
		Documents:    documents,
		// MapID:        record.MapID, TODO: IMPLEMENT!!!
	}

	if err := s.storage.CreateRecord(ctx, recordModel); err != nil {
		log.Error("failed to create record", "error", err)
		return "", fmt.Errorf("%w: %s", customerrors.ErrBadRequest, "failed to create record")
	}

	return recordID, nil
}

// TODO: check role
func (s *RecordService) PublishRecord(ctx context.Context, recordID string, tokenPayload authorization.TokenPayload) error {
	const op = "recordService.PublishRecord"
	log := s.log.With("op", op)

	log.Info("Publish record", "recordID", recordID)

	if tokenPayload.Role != "admin" {
		return fmt.Errorf("%w: %s", customerrors.ErrForbidden, "only admin can publish record")
	}

	if err := s.storage.PublishRecord(ctx, recordID); err != nil {
		log.Error("failed to publish record", "error", err)
		return fmt.Errorf("%w: %s", customerrors.ErrBadRequest, "failed to publish record")
	}

	return nil
}

func (s *RecordService) UpdateRecord(ctx context.Context, record *models.RecordModel) error {
	const op = "recordService.UpdateRecord"
	log := s.log.With("op", op)

	log.Info("Update record", "record", record)

	if err := s.storage.UpdateRecord(ctx, record); err != nil {
		log.Error("failed to update record", "error", err)
		return fmt.Errorf("%w: %s", customerrors.ErrBadRequest, "failed to update record")
	}

	return nil
}

func (s *RecordService) AddDocument(ctx context.Context, recordID string, document *requests.DocumentRequest) error {
	const op = "recordService.AddDocument"
	log := s.log.With("op", op)

	log.Info("Add document", "document", document)

	documentModel := &models.DocumentModel{
		ID:       uuid.New().String(),
		Type:     document.Type,
		URL:      document.URL,
		RecordID: recordID,
	}

	if err := s.storage.AddDocument(ctx, documentModel); err != nil {
		log.Error("failed to add document", "error", err)
		return fmt.Errorf("%w: %s", customerrors.ErrBadRequest, "failed to add document")
	}

	return nil
}
