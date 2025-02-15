package conflict

import (
	"context"

	"github.com/davg/records/internal/domain/models"
	"github.com/davg/records/internal/server/utils"
	"github.com/gin-gonic/gin"
)

type ConflictService interface {
	Conflicts(ctx context.Context) (*[]models.ConflictModel, error)
	Conflict(ctx context.Context, id string) (*models.ConflictModel, error)
}

type ConflictRouter struct {
	service ConflictService
	router  *gin.RouterGroup
}

func Register(router *gin.RouterGroup, service ConflictService) {
	conflictRouter := &ConflictRouter{
		router:  router,
		service: service,
	}

	conflictRouter.init()
}

func (r *ConflictRouter) Conflicts(ctx *gin.Context) {
	conflicts, err := r.service.Conflicts(ctx)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}
	ctx.JSON(200, conflicts)
}

func (r *ConflictRouter) Conflict(ctx *gin.Context) {
	id := ctx.Param("id")
	conflict, err := r.service.Conflict(ctx, id)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}
	ctx.JSON(200, conflict)
}

func (r *ConflictRouter) init() {
	conflictRouter := r.router.Group("/conflicts")

	conflictRouter.GET("/", r.Conflicts)
	conflictRouter.GET("/:id", r.Conflict)
}
