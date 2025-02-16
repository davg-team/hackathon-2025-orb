package aplication

import (
	"log/slog"

	"github.com/davg/drafts/internal/server"
	"github.com/davg/drafts/internal/service"
	"github.com/davg/drafts/internal/storage"
)

type Aplication struct {
	server *server.Server
}

func New(log *slog.Logger) *Aplication {
	storage := storage.New()

	service := service.New(storage, log)

	server := server.New(service)

	return &Aplication{
		server: server,
	}
}

func (a *Aplication) Start() {
	a.server.Run()
}

func (a *Aplication) GracefulStop() {
	a.server.GracefulStop()
}
