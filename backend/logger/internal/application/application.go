package application

import (
	"log/slog"

	"github.com/davg/logger/internal/server"
	"github.com/davg/logger/internal/service"
	"github.com/davg/logger/internal/storage"
)

type Application struct {
	server *server.Server
}

func New(log *slog.Logger) *Application {
	collection := connectToMongoDB()
	storage := storage.NewStorage(collection)

	service := service.New(storage, *log)
	
	server := server.New(service)

	return &Application{
		server: server,
	}
}

func (a *Application) Start() error {
	return a.server.Start()
}

func (a *Application) GracefulStop() error {
	return a.server.GracefulStop()
}
