package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/davg/logger/internal/config"
	"github.com/davg/logger/internal/domain"
	"github.com/davg/logger/internal/domain/requests"
	"github.com/davg/logger/internal/server/log"
	"github.com/gin-gonic/gin"
)

type Service interface {
	Log(ctx context.Context, id string) ([]domain.LogModel, error)
	Logs(ctx context.Context) ([]domain.LogModel, error)
	CreateLog(ctx context.Context, log requests.LogPOST) (string, error)
}

type Server struct {
	server *http.Server
	engine *gin.Engine
}

func New(service Service) *Server {
	cfg := config.Config()
	engine := gin.Default()

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Server.Port),
		Handler: engine,
	}

	group := engine.Group("/api")

	log.Register(group, service)

	return &Server{
		server: httpServer,
		engine: engine,
	}
}

func (s *Server) Start() error {
	if err := s.server.ListenAndServe(); err != nil {
		panic(err)
	}

	return nil
}

func (s *Server) GracefulStop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
