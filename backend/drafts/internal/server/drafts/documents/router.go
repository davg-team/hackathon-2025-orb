package drafts_documents

import (
	"context"
	"crypto/rsa"

	"github.com/davg/drafts/internal/domain"
	"github.com/davg/drafts/internal/domain/requests"
	"github.com/davg/drafts/internal/server/utils"
	"github.com/davg/drafts/pkg/middlewares/authorization"
	"github.com/gin-gonic/gin"
)

type Service interface {
	DraftDocument(ctx context.Context, id string) (*domain.OfferDocumentUpdate, error)
	DraftsDocuments(ctx context.Context) (*[]domain.OfferDocumentUpdate, error)
	DeleteDraftDocument(ctx context.Context, id string) error
	CreateDraftDocument(ctx context.Context, draft *requests.UpdateDocumentPost) (string, error)
}

type Router struct {
	router  *gin.RouterGroup
	service Service
}

func Register(router *gin.RouterGroup, service Service, key *rsa.PublicKey) {
	r := &Router{
		router:  router,
		service: service,
	}

	r.init(key)
}

func (r *Router) DraftsDocuments(ctx *gin.Context) {
	_, err := authorization.FromContext(ctx)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	drafts, err := r.service.DraftsDocuments(ctx)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	ctx.JSON(200, drafts)
}

func (r *Router) DraftDocument(ctx *gin.Context) {
	_, err := authorization.FromContext(ctx)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	id := ctx.Param("id")
	draft, err := r.service.DraftDocument(ctx, id)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	ctx.JSON(200, draft)
}

func (r *Router) CreateDraftDocument(ctx *gin.Context) {
	_, err := authorization.FromContext(ctx)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	var draft requests.UpdateDocumentPost
	if err := ctx.ShouldBindJSON(&draft); err != nil {
		utils.HandleError(ctx, err)
		return
	}

	id, err := r.service.CreateDraftDocument(ctx, &draft)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{"id": id})
}

func (r *Router) DeleteDraftDocument(ctx *gin.Context) {
	_, err := authorization.FromContext(ctx)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	id := ctx.Param("id")
	if err := r.service.DeleteDraftDocument(ctx, id); err != nil {
		utils.HandleError(ctx, err)
		return
	}

	ctx.JSON(200, nil)
}

func (r *Router) init(key *rsa.PublicKey) {
	draftRouter := r.router.Group("/drafts_documents")
	draftRouter.GET("/", authorization.MiddlwareJWT(key), r.DraftsDocuments)
	draftRouter.GET("/:id", authorization.MiddlwareJWT(key), r.DraftDocument)
	draftRouter.POST("/", authorization.MiddlwareJWT(key), r.CreateDraftDocument)
	draftRouter.DELETE("/:id", authorization.MiddlwareJWT(key), r.DeleteDraftDocument)
}
