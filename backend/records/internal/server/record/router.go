package record

import (
	"context"
	"crypto/rsa"
	"fmt"

	"github.com/davg/records/internal/domain/models"
	"github.com/davg/records/internal/domain/requests"
	"github.com/davg/records/internal/server/utils"
	authorization "github.com/davg/records/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

type RecordService interface {
	Record(ctx context.Context, id string) (*models.RecordModel, error)
	Records(ctx context.Context, limit, offset string) (*[]models.RecordModel, int, error)
	RecordsByParams(
		ctx context.Context,
		name string,
		middleName string,
		lastName string,
		limit, offset string,
		published, userID string,
	) (*[]models.RecordModel, int, error)
	CreateRecord(ctx context.Context, record *requests.RecordRequest, tokenPayload authorization.TokenPayload) (string, error)
	PublishRecord(ctx context.Context, recordID string, tokenPayload authorization.TokenPayload) error
	UpdateRecord(ctx context.Context, id string, record *models.RecordModel) error

	AddDocument(ctx context.Context, recordID string, document *requests.DocumentRequest) error
}

type RecordRouter struct {
	router  *gin.RouterGroup
	service RecordService
}

func Register(router *gin.RouterGroup, service RecordService, key *rsa.PublicKey) {
	recordRouter := &RecordRouter{
		router:  router,
		service: service,
	}

	recordRouter.init(key)
}

func (r *RecordRouter) GetRecord(ctx *gin.Context) {
	id := ctx.Param("id")
	record, err := r.service.Record(ctx, id)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}
	ctx.JSON(200, record)
}

func (r *RecordRouter) GetRecordsByParams(ctx *gin.Context) {
	name := ctx.Query("name")
	middleName := ctx.Query("middleName")
	lastName := ctx.Query("lastName")
	limit := ctx.Query("limit")
	offset := ctx.Query("offset")
	published := ctx.Query("published")
	userID := ctx.Query("userID")
	records, totalCount, err := r.service.RecordsByParams(
		ctx,
		name,
		middleName,
		lastName,
		limit, offset,
		published, userID,
	)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	ctx.Header("X-Total-Count", fmt.Sprintf("%d", totalCount))
	ctx.JSON(200, records)
}

func (r *RecordRouter) CreateRecord(ctx *gin.Context) {
	payload, err := authorization.FromContext(ctx)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}
	var record requests.RecordRequest
	if err := ctx.ShouldBindJSON(&record); err != nil {
		utils.HandleError(ctx, err)
		return
	}

	recordID, err := r.service.CreateRecord(ctx, &record, payload)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}
	ctx.JSON(201, gin.H{"id": recordID})
}

func (r *RecordRouter) PublishRecord(ctx *gin.Context) {
	payload, err := authorization.FromContext(ctx)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}
	recordID := ctx.Param("id")
	if err := r.service.PublishRecord(ctx, recordID, payload); err != nil {
		utils.HandleError(ctx, err)
		return
	}
	ctx.JSON(200, gin.H{"message": "record published"})
}

func (r *RecordRouter) UpdateRecord(ctx *gin.Context) {
	id := ctx.Param("id")
	var record models.RecordModel
	if err := ctx.ShouldBindJSON(&record); err != nil {
		utils.HandleError(ctx, err)
		return
	}
	if err := r.service.UpdateRecord(ctx, id, &record); err != nil {
		utils.HandleError(ctx, err)
		return
	}
	ctx.JSON(200, record)
}

func (r *RecordRouter) AddDocument(ctx *gin.Context) {
	_, err := authorization.FromContext(ctx)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	recordID := ctx.Param("id")
	var document requests.DocumentRequest
	if err := ctx.ShouldBindJSON(&document); err != nil {
		utils.HandleError(ctx, err)
		return
	}
	if err := r.service.AddDocument(ctx, recordID, &document); err != nil {
		utils.HandleError(ctx, err)
		return
	}
	ctx.JSON(200, gin.H{"message": "document added"})
}

func (r *RecordRouter) init(key *rsa.PublicKey) {
	recordRouter := r.router.Group("/records")

	recordRouter.GET("/:id", r.GetRecord)
	recordRouter.GET("/", r.GetRecordsByParams)
	recordRouter.PUT("/:id", r.UpdateRecord)

	recordRouter.POST("/", authorization.MiddlwareJWT(key), r.CreateRecord)
	recordRouter.PATCH("/:id/publish", authorization.MiddlwareJWT(key), r.PublishRecord)

	recordRouter.POST("/:id/documents", authorization.MiddlwareJWT(key), r.AddDocument)
}
