package conflict

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/davg/records/internal/customerrors"
	"github.com/davg/records/internal/domain/models"
)

type ConflictStorage interface {
	Conflicts(ctx context.Context) (*[]models.ConflictModel, error)
	Conflict(ctx context.Context, id string) (*models.ConflictModel, error)
}

type ConflictService struct {
	log     *slog.Logger
	storage ConflictStorage
}

func NewConflictService(log *slog.Logger, storage ConflictStorage) *ConflictService {
	log = log.With("service", "conflict")
	return &ConflictService{
		log:     log,
		storage: storage,
	}
}

func (s *ConflictService) Conflicts(ctx context.Context) (*[]models.ConflictModel, error) {
	const op = "conflictService.Conflicts"
	log := s.log.With("op", op)

	log.Info("Get all conflicts")

	conflicts, err := s.storage.Conflicts(ctx)
	if err != nil {
		log.Error("failed to get conflicts", "error", err)
		return nil, fmt.Errorf("%w: %s", customerrors.ErrInternal, "failed to get conflicts")
	}

	return conflicts, nil
}

func (s *ConflictService) Conflict(ctx context.Context, id string) (*models.ConflictModel, error) {
	const op = "conflictService.Conflict"
	log := s.log.With("op", op)

	log.Info("Get conflict by id", "id", id)

	conflict, err := s.storage.Conflict(ctx, id)
	if err != nil {
		log.Error("failed to get conflict", "error", err)
		return nil, fmt.Errorf("%w: %s", customerrors.ErrNotFound, "conflict by id not found")
	}

	return conflict, nil
}
