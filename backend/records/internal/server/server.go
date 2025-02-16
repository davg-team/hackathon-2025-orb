package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/davg/records/internal/config"
	"github.com/davg/records/internal/domain/models"
	"github.com/davg/records/internal/domain/requests"
	conflictRouter "github.com/davg/records/internal/server/conflict"
	recordRouter "github.com/davg/records/internal/server/record"
	"github.com/davg/records/internal/server/utils"
	authorization "github.com/davg/records/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

type ConflictService interface {
	Conflicts(ctx context.Context) (*[]models.ConflictModel, error)
	Conflict(ctx context.Context, id string) (*models.ConflictModel, error)
}

type RecordService interface {
	Record(ctx context.Context, id string) (*models.RecordModel, error)
	Records(ctx context.Context, limit, offset string) (*[]models.RecordModel, int, error)
	RecordsByParams(ctx context.Context, name string, middleName string, lastName string, limit, offset string, published, userID string) (*[]models.RecordModel, int, error)
	CreateRecord(ctx context.Context, record *requests.RecordRequest, tokenPayload authorization.TokenPayload) (string, error)
	PublishRecord(ctx context.Context, recordID string, tokenPayload authorization.TokenPayload) error
	UpdateRecord(ctx context.Context, id string, record *models.RecordModel) error

	AddDocument(ctx context.Context, recordID string, document *requests.DocumentRequest) error
}

type Server struct {
	server *http.Server
	engine *gin.Engine
}

func New(
	conflictService ConflictService,
	recordService RecordService,
) *Server {
	cfg := config.Config().Server
	engine := gin.Default()

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: engine,
	}

	group := engine.Group("/api")
	group.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	key := utils.GetKey("./keys/public.pem")

	conflictRouter.Register(group, conflictService)
	recordRouter.Register(group, recordService, key)

	return &Server{
		server: httpServer,
		engine: engine,
	}
}

func (s *Server) Start() {
	if err := s.server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func (s *Server) GracefulStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		panic(err)
	}
}
