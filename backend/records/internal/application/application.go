package application

import (
	"log/slog"

	"github.com/davg/records/internal/server"
	"github.com/davg/records/internal/service/conflict"
	"github.com/davg/records/internal/service/record"
	"github.com/davg/records/internal/storage"
)

type Application struct {
	server *server.Server
}

func New(log *slog.Logger) *Application {
	storage := storage.New()

	recordService := record.NewRecordService(log, storage)
	conflictService := conflict.NewConflictService(log, storage)

	server := server.New(conflictService, recordService)

	return &Application{
		server: server,
	}
}

func (a *Application) Start() {
	a.server.Start()
}

func (a *Application) GracefulStop() {
	a.server.GracefulStop()
}
