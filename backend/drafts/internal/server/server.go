package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/davg/drafts/internal/config"
	"github.com/davg/drafts/internal/domain"
	"github.com/davg/drafts/internal/domain/requests"
	drafts_documents "github.com/davg/drafts/internal/server/drafts/documents"
	"github.com/davg/drafts/internal/server/drafts/drafts_record"
	"github.com/davg/drafts/internal/server/utils"
	"github.com/gin-gonic/gin"
)

type Service interface {
	DraftDocument(ctx context.Context, id string) (*domain.OfferDocumentUpdate, error)
	DraftRecord(ctx context.Context, id string) (*domain.OfferRecordUpdate, error)
	DraftsDocuments(ctx context.Context) (*[]domain.OfferDocumentUpdate, error)
	DraftsRecords(ctx context.Context) (*[]domain.OfferRecordUpdate, error)
	DeleteDraftDocument(ctx context.Context, id string) error
	DeleteDraftRecord(ctx context.Context, id string) error
	CreateDraftDocument(ctx context.Context, draft *requests.UpdateDocumentPost) (string, error)
	CreateDraftRecord(ctx context.Context, draft *requests.UpdateRecordPost) (string, error)
	UpdateRecordDraftOnServer(ctx context.Context, id string) error
}

type Server struct {
	server *http.Server
	engine *gin.Engine
}

func New(service Service) *Server {
	cfg := config.Config().Server
	engine := gin.Default()
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: engine,
	}

	group := engine.Group("/api")

	key := utils.GetKey("./keys/public.pem")

	drafts_documents.Register(group, service, key)
	drafts_record.Register(group, service, key)

	return &Server{
		server: httpServer,
		engine: engine,
	}
}

func (s *Server) Run() {
	if err := s.server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func (s *Server) GracefulStop() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		panic(err)
	}
}
