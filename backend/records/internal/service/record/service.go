package record

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"strconv"

	"github.com/davg/records/internal/config"
	"github.com/davg/records/internal/customerrors"
	"github.com/davg/records/internal/domain/models"
	"github.com/davg/records/internal/domain/requests"
	"github.com/davg/records/internal/service"
	authorization "github.com/davg/records/pkg/middlewares"
	"github.com/davg/records/pkg/nextgis"
	"github.com/davg/records/pkg/s3storage"
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
	UpdateRecordMapID(ctx context.Context, recordID string, mapID int) error

	AddDocument(ctx context.Context, document *models.DocumentModel) error
	Conflict(ctx context.Context, id string) (*models.ConflictModel, error)
}

type RecordService struct {
	log         *slog.Logger
	storage     RecordStorage
	client      *nextgis.Client
	minioBucket *s3storage.MinioStorage
}

func NewRecordService(log *slog.Logger, storage RecordStorage) *RecordService {
	cfg := config.Config().Client
	client := nextgis.NewClient(cfg.LayerID, cfg.Username, cfg.Password)
	minioBucket := s3storage.Connect()
	log = log.With("service", "record")
	return &RecordService{
		log:         log,
		storage:     storage,
		client:      client,
		minioBucket: minioBucket,
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

	log.Info("Get records by params", "name", name, "middleName", middleName, "lastName", lastName)

	records, totalCount, err := s.storage.RecordsByParams(ctx, name, middleName, lastName, limitInt, offsetInt, published, userID)
	if err != nil {
		log.Error("failed to get records", "error", err)
		return nil, 0, fmt.Errorf("%w: %s", customerrors.ErrInternal, "failed to get records")
	}

	return records, totalCount, nil
}

func (s *RecordService) CreateRecord(ctx context.Context, record *requests.RecordRequest, tokenPayload authorization.TokenPayload) (string, error) {
	const op = "recordService.CreateRecord"
	log := s.log.With("op", op)

	log.Info("Create record", "record", record)

	recordID := uuid.New().String()
	var documents []models.DocumentModel
	for _, doc := range record.Documents {
		documents = append(documents, models.DocumentModel{
			ID:        uuid.New().String(),
			Type:      doc.Type,
			ObjectKey: doc.ObjectKey,
			RecordID:  recordID,
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
		UserID:       tokenPayload.ID,
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

	if tokenPayload.Role != "superadmin" {
		return fmt.Errorf("%w: %s", customerrors.ErrForbidden, "only superadmin can publish record")
	}

	if err := s.storage.PublishRecord(ctx, recordID); err != nil {
		log.Error("failed to publish record", "error", err)
		return fmt.Errorf("%w: %s", customerrors.ErrBadRequest, "failed to publish record")
	}

	// Loading record to the map
	record, err := s.storage.Record(ctx, recordID)
	if err != nil {
		log.Error("failed to get record", "error", err)
		return fmt.Errorf("%w: %s", customerrors.ErrBadRequest, "failed to get record")
	}

	var conflictNames string
	for _, conflict := range record.Conflicts {
		conflictNames += conflict.Title + ", "
	}
	if len(record.Conflicts) > 0 {
		conflictNames = conflictNames[:len(conflictNames)-2]
	}

	var awardNames string
	for _, award := range record.Awards {
		awardNames += award + ", "
	}
	if len(record.Awards) > 0 {
		awardNames = awardNames[:len(awardNames)-2]
	}

	latitude, longitude, err := service.GetCoordinates("Оренбургская область, " + record.BirthPlace)
	if err != nil {
		log.Error("failed to get coordinates", "error", err)
		return fmt.Errorf("%w: %s", customerrors.ErrBadRequest, "failed to get coordinates")
	}

	recordData := map[string]interface{}{
		"fields": map[string]interface{}{
			"num":      1,
			"n_raion":  record.BirthPlace,
			"fio":      fmt.Sprintf("%s %s %s", record.LastName, record.Name, record.MiddleName),
			"years":    fmt.Sprintf("%s - %s", record.BirthDate, record.DeathDate),
			"info":     record.Bio,
			"kontrakt": conflictNames,
			"nagrads":  awardNames,
		},
		"geom": fmt.Sprintf("POINT(%f %f)", latitude, longitude),
	}

	id, err := s.client.AddFeature(recordData)
	if err != nil {
		log.Error("failed to add feature", "error", err)
		return fmt.Errorf("%w: %s", customerrors.ErrBadRequest, "failed to add feature")
	}

	if err := s.storage.UpdateRecordMapID(ctx, recordID, id); err != nil {
		log.Error("failed to update mapID", "error", err)
		return fmt.Errorf("%w: %s", customerrors.ErrBadRequest, "failed to update mapID")
	}

	// Loading documents to the map
	for _, doc := range record.Documents {
		object, err := s.minioBucket.GetFile(ctx, doc.ObjectKey)
		if err != nil {
			log.Error("failed to get file", "error", err)
			return fmt.Errorf("%w: %s", customerrors.ErrBadRequest, "failed to get file")
		}
		defer object.Close()

		fileBuffer := new(bytes.Buffer)
		_, err = io.Copy(fileBuffer, object)
		if err != nil {
			log.Error("failed to copy file", "error", err)
			return fmt.Errorf("%w: %s", customerrors.ErrInternal, "failed to copy file")
		}

		objectMetadata, err := s.client.UploadFile(fileBuffer, doc.ObjectKey)
		if err != nil {
			log.Error("failed to upload file", "error", err)
			return fmt.Errorf("%w: %s", customerrors.ErrInternal, "failed to upload file")
		}

		if _, err := s.client.AttachFile(id, objectMetadata.UploadMeta[0].ID, objectMetadata.UploadMeta[0].Name, objectMetadata.UploadMeta[0].MimeType, objectMetadata.UploadMeta[0].Size); err != nil {
			log.Error("failed to attach file", "error", err)
			return fmt.Errorf("%w: %s", customerrors.ErrInternal, "failed to attach file")
		}

	}

	return nil
}

func (s *RecordService) UpdateRecord(ctx context.Context, id string, record *models.RecordModel) error {
	const op = "recordService.UpdateRecord"
	log := s.log.With("op", op)

	log.Info("Update record", "record", record)

	recordModel, err := s.storage.Record(ctx, id)
	if err != nil {
		log.Error("failed to get record", "error", err)
		return fmt.Errorf("%w: %s", customerrors.ErrNotFound, "record by id not found")
	}

	record.ID = id
	record.UserID = recordModel.UserID
	record.MapID = recordModel.MapID
	record.Published = recordModel.Published

	if err := s.storage.UpdateRecord(ctx, record); err != nil {
		log.Error("failed to update record", "error", err)
		return fmt.Errorf("%w: %s", customerrors.ErrBadRequest, "failed to update record")
	}

	record, err = s.storage.Record(ctx, id)
	if err != nil {
		log.Error("failed to get record", "error", err)
		return fmt.Errorf("%w: %s", customerrors.ErrBadRequest, "failed to get record")
	}

	var conflictNames string
	for _, conflict := range record.Conflicts {
		conflictNames += conflict.Title + ", "
	}
	if len(record.Conflicts) > 0 {
		conflictNames = conflictNames[:len(conflictNames)-2]
	}

	var awardNames string
	for _, award := range record.Awards {
		awardNames += award + ", "
	}
	if len(record.Awards) > 0 {
		awardNames = awardNames[:len(awardNames)-2]
	}

	latitude, longitude, err := service.GetCoordinates("Оренбургская область, " + record.BirthPlace)
	if err != nil {
		log.Error("failed to get coordinates", "error", err)
		return fmt.Errorf("%w: %s", customerrors.ErrBadRequest, "failed to get coordinates")
	}

	recordData := map[string]interface{}{
		"fields": map[string]interface{}{
			"num":      1,
			"n_raion":  record.BirthPlace,
			"fio":      fmt.Sprintf("%s %s %s", record.LastName, record.Name, record.MiddleName),
			"years":    fmt.Sprintf("%s - %s", record.BirthDate, record.DeathDate),
			"info":     record.Bio,
			"kontrakt": conflictNames,
			"nagrads":  awardNames,
		},
		"geom": fmt.Sprintf("POINT(%f %f)", latitude, longitude),
	}

	if err := s.client.UpdateFeature(record.MapID, recordData); err != nil {
		log.Error("failed to update feature", "error", err)
		return fmt.Errorf("%w: %s", customerrors.ErrBadRequest, "failed to update feature")
	}

	return nil
}

func (s *RecordService) AddDocument(ctx context.Context, recordID string, document *requests.DocumentRequest) error {
	const op = "recordService.AddDocument"
	log := s.log.With("op", op)

	log.Info("Add document", "document", document)

	documentModel := &models.DocumentModel{
		ID:        uuid.New().String(),
		Type:      document.Type,
		ObjectKey: document.ObjectKey,
		RecordID:  recordID,
	}

	if err := s.storage.AddDocument(ctx, documentModel); err != nil {
		log.Error("failed to add document", "error", err)
		return fmt.Errorf("%w: %s", customerrors.ErrBadRequest, "failed to add document")
	}

	record, err := s.storage.Record(ctx, recordID)
	if err != nil {
		log.Error("failed to get record", "error", err)
		return fmt.Errorf("%w: %s", customerrors.ErrBadRequest, "failed to get record")
	}

	object, err := s.minioBucket.GetFile(ctx, document.ObjectKey)
	if err != nil {
		log.Error("failed to get file", "error", err)
		return fmt.Errorf("%w: %s", customerrors.ErrBadRequest, "failed to get file")
	}
	defer object.Close()

	fileBuffer := new(bytes.Buffer)
	_, err = io.Copy(fileBuffer, object)
	if err != nil {
		log.Error("failed to copy file", "error", err)
		return fmt.Errorf("%w: %s", customerrors.ErrInternal, "failed to copy file")
	}

	objectMetadata, err := s.client.UploadFile(fileBuffer, document.ObjectKey)
	if err != nil {
		log.Error("failed to upload file", "error", err)
		return fmt.Errorf("%w: %s", customerrors.ErrInternal, "failed to upload file")
	}

	if _, err := s.client.AttachFile(record.MapID, objectMetadata.UploadMeta[0].ID, objectMetadata.UploadMeta[0].Name, objectMetadata.UploadMeta[0].MimeType, objectMetadata.UploadMeta[0].Size); err != nil {
		log.Error("failed to attach file", "error", err)
		return fmt.Errorf("%w: %s", customerrors.ErrInternal, "failed to attach file")
	}

	return nil
}
