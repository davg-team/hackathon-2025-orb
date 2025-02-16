package service

import (
	"context"
	"log/slog"

	"github.com/davg/drafts/internal/domain"
	"github.com/davg/drafts/internal/domain/requests"
	"github.com/davg/drafts/pkg/client"
	"github.com/google/uuid"
)

type Storage interface {
	DraftDocument(ctx context.Context, id string) (*domain.OfferDocumentUpdate, error)
	DraftRecord(ctx context.Context, id string) (*domain.OfferRecordUpdate, error)
	DraftsDocuments(ctx context.Context) (*[]domain.OfferDocumentUpdate, error)
	DraftsRecords(ctx context.Context) (*[]domain.OfferRecordUpdate, error)

	CreateDraftDocument(ctx context.Context, draft *domain.OfferDocumentUpdate) error
	CreateDraftRecord(ctx context.Context, draft *domain.OfferRecordUpdate) error
	DeleteDraftDocument(ctx context.Context, id string) error
	DeleteDraftRecord(ctx context.Context, id string) error
}

type Service struct {
	storage Storage
	log     *slog.Logger
	client  *client.Client
}

func New(storage Storage, log *slog.Logger) *Service {
	log = log.With(slog.String("service", "draft"))
	client := client.NewClient("http://records:8080")

	return &Service{
		storage: storage,
		client:  client,
		log:     log,
	}
}

func (s *Service) DraftDocument(ctx context.Context, id string) (*domain.OfferDocumentUpdate, error) {
	const op = "service.DraftDocument"

	logger := s.log.With(slog.String("op", op))
	logger.Info("draft document")

	draft, err := s.storage.DraftDocument(ctx, id)
	if err != nil {
		logger.Error("failed to get draft document", err.Error())
		return nil, err
	}

	logger.Info("draft document received")

	return draft, nil
}

func (s *Service) DraftRecord(ctx context.Context, id string) (*domain.OfferRecordUpdate, error) {
	const op = "service.DraftRecord"

	logger := s.log.With(slog.String("op", op))
	logger.Info("draft record")

	draft, err := s.storage.DraftRecord(ctx, id)
	if err != nil {
		logger.Error("failed to get draft record", err.Error())
		return nil, err
	}

	logger.Info("draft record received")

	return draft, nil
}

func (s *Service) DraftsDocuments(ctx context.Context) (*[]domain.OfferDocumentUpdate, error) {
	const op = "service.DraftsDocuments"

	logger := s.log.With(slog.String("op", op))
	logger.Info("drafts documents")

	drafts, err := s.storage.DraftsDocuments(ctx)
	if err != nil {
		logger.Error("failed to get drafts documents", err.Error())
		return nil, err
	}

	logger.Info("drafts documents received")

	return drafts, nil
}

func (s *Service) DraftsRecords(ctx context.Context) (*[]domain.OfferRecordUpdate, error) {
	const op = "service.DraftsRecords"

	logger := s.log.With(slog.String("op", op))
	logger.Info("drafts records")

	drafts, err := s.storage.DraftsRecords(ctx)
	if err != nil {
		logger.Error("failed to get drafts records", err.Error())
		return nil, err
	}

	logger.Info("drafts records received")

	return drafts, nil
}

func (s *Service) CreateDraftDocument(ctx context.Context, draft *requests.UpdateDocumentPost) (string, error) {
	const op = "service.CreateDraftDocument"

	logger := s.log.With(slog.String("op", op))
	logger.Info("create draft document")

	id := uuid.New().String()

	draft_model := &domain.OfferDocumentUpdate{
		ID:         id,
		DocumentID: uuid.New().String(),
		RecordID:   draft.RecordID,
		Type:       draft.Type,
		URL:        draft.URL,
	}

	if err := s.storage.CreateDraftDocument(ctx, draft_model); err != nil {
		logger.Error("failed to create draft document", err.Error())
		return "", err
	}

	logger.Info("draft document created")

	return id, nil
}

func (s *Service) CreateDraftRecord(ctx context.Context, draft *requests.UpdateRecordPost) (string, error) {
	const op = "service.CreateDraftRecord"

	logger := s.log.With(slog.String("op", op))
	logger.Info("create draft record")

	id := uuid.New().String()

	draft_model := &domain.OfferRecordUpdate{
		ID:           id,
		RecordID:     draft.RecordID,
		Name:         draft.Name,
		MiddleName:   draft.MiddleName,
		LastName:     draft.LastName,
		BirthDate:    draft.BirthDate,
		BirthPlace:   draft.BirthPlace,
		MilitaryRank: draft.MilitaryRank,
		Commissariat: draft.Commissariat,
		Awards:       draft.Awards,
		DeathDate:    draft.DeathDate,
		BurialPlace:  draft.BurialPlace,
		Bio:          draft.Bio,
	}

	if err := s.storage.CreateDraftRecord(ctx, draft_model); err != nil {
		logger.Error("failed to create draft record", err.Error())
		return "", err
	}

	logger.Info("draft record created")

	return id, nil
}

func (s *Service) DeleteDraftDocument(ctx context.Context, id string) error {
	const op = "service.DeleteDraftDocument"

	logger := s.log.With(slog.String("op", op))
	logger.Info("delete draft document")

	if err := s.storage.DeleteDraftDocument(ctx, id); err != nil {
		logger.Error("failed to delete draft document", err.Error())
		return err
	}

	logger.Info("draft document deleted")

	return nil
}

func (s *Service) DeleteDraftRecord(ctx context.Context, id string) error {
	const op = "service.DeleteDraftRecord"

	logger := s.log.With(slog.String("op", op))
	logger.Info("delete draft record")

	if err := s.storage.DeleteDraftRecord(ctx, id); err != nil {
		logger.Error("failed to delete draft record", err.Error())
		return err
	}

	logger.Info("draft record deleted")

	return nil
}

func (s *Service) UpdateRecordDraftOnServer(ctx context.Context, id string) error {
	const op = "service.UpdateRecordDraftOnServer"

	logger := s.log.With(slog.String("op", op))
	logger.Info("update record draft on server")

	record, err := s.DraftRecord(ctx, id)
	if err != nil {
		return err
	}

	logger.Info("record draft updated on server")

	return s.client.UpdateRecord(record.RecordID, record)
}
