package service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/davg/logger/internal/domain"
	"github.com/davg/logger/internal/domain/requests"
	"github.com/davg/logger/internal/errors"
	"github.com/google/uuid"
)

type Storage interface {
	Log(ctx context.Context, id string) ([]domain.LogModel, error)
	Logs(ctx context.Context) ([]domain.LogModel, error)
	CreateLog(ctx context.Context, log domain.LogModel) (string, error)
}

type Service struct {
	storage Storage
	log     slog.Logger
}

func New(storage Storage, log slog.Logger) *Service {
	return &Service{
		storage: storage,
		log:     log,
	}
}

func (s *Service) Log(ctx context.Context, id string) ([]domain.LogModel, error) {
	const op = "service.Log"

	logger := s.log.With(slog.String("op", op))
	logger.Info("Getting log")

	log, err := s.storage.Log(ctx, id)
	if err != nil {
		s.log.Error("error getting log", slog.String("error", err.Error()))
		return nil, fmt.Errorf("%w: %s", errors.InternalServerError, "error getting log")
	}

	logger.Info("Log found")
	return log, nil
}

func (s *Service) Logs(ctx context.Context) ([]domain.LogModel, error) {
	const op = "service.Logs"

	logger := s.log.With(slog.String("op", op))
	logger.Info("Getting logs")

	logs, err := s.storage.Logs(ctx)
	if err != nil {
		s.log.Error("error getting logs", slog.String("error", err.Error()))
		return nil, fmt.Errorf("%w: %s", errors.InternalServerError, "error getting logs")
	}

	logger.Info("Logs found")
	return logs, nil
}

func (s *Service) CreateLog(ctx context.Context, log requests.LogPOST) (string, error) {
	const op = "service.CreateLog"

	logger := s.log.With(slog.String("op", op))
	logger.Info("Creating log")

	log_model := domain.LogModel{
		ID:     uuid.New().String(),
		UserID: log.UserID,
		Action: log.Action,
		Time:   log.Time,
		Info:   log.Info,
	}

	id, err := s.storage.CreateLog(ctx, log_model)
	if err != nil {
		s.log.Error("error creating log", slog.String("error", err.Error()))
		return "", fmt.Errorf("%w: %s", errors.InternalServerError, "error creating log")
	}

	logger.Info("Log created")
	return id, nil
}
