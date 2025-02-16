package drafts_record

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
	DraftRecord(ctx context.Context, id string) (*domain.OfferRecordUpdate, error)
	DraftsRecords(ctx context.Context) (*[]domain.OfferRecordUpdate, error)
	DeleteDraftRecord(ctx context.Context, id string) error
	CreateDraftRecord(ctx context.Context, draft *requests.UpdateRecordPost) (string, error)
	UpdateRecordDraftOnServer(ctx context.Context, id string) error
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

func (r *Router) DraftsRecords(ctx *gin.Context) {
	_, err := authorization.FromContext(ctx)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	drafts, err := r.service.DraftsRecords(ctx)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	ctx.JSON(200, drafts)
}

func (r *Router) DraftRecord(ctx *gin.Context) {
	_, err := authorization.FromContext(ctx)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	id := ctx.Param("id")
	draft, err := r.service.DraftRecord(ctx, id)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	ctx.JSON(200, draft)
}

func (r *Router) CreateDraftRecord(ctx *gin.Context) {
	_, err := authorization.FromContext(ctx)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	var draft requests.UpdateRecordPost
	if err := ctx.ShouldBindJSON(&draft); err != nil {
		utils.HandleError(ctx, err)
		return
	}

	id, err := r.service.CreateDraftRecord(ctx, &draft)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{"id": id})
}

func (r *Router) DeleteDraftRecord(ctx *gin.Context) {
	_, err := authorization.FromContext(ctx)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	id := ctx.Param("id")
	if err := r.service.DeleteDraftRecord(ctx, id); err != nil {
		utils.HandleError(ctx, err)
		return
	}

	ctx.JSON(200, nil)
}

func (r *Router) UpdateRecordDraftOnServer(ctx *gin.Context) {
	payload, err := authorization.FromContext(ctx)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	if payload.Role != "superadmin" && payload.Role != "root" {
		utils.HandleError(ctx, err)
		return
	}

	id := ctx.Param("id")
	if err := r.service.UpdateRecordDraftOnServer(ctx, id); err != nil {
		utils.HandleError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{"status": "ok"})
}

func (r *Router) init(key *rsa.PublicKey) {
	draftRouter := r.router.Group("/drafts_records")
	draftRouter.GET("/", authorization.MiddlwareJWT(key), r.DraftsRecords)
	draftRouter.GET("/:id", authorization.MiddlwareJWT(key), r.DraftRecord)
	draftRouter.POST("/", authorization.MiddlwareJWT(key), r.CreateDraftRecord)
	draftRouter.DELETE("/:id", authorization.MiddlwareJWT(key), r.DeleteDraftRecord)
	draftRouter.PATCH("/:id", authorization.MiddlwareJWT(key), r.UpdateRecordDraftOnServer)
}
