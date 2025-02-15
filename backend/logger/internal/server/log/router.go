package log

import (
	"context"

	"github.com/davg/logger/internal/domain"
	"github.com/davg/logger/internal/domain/requests"
	"github.com/gin-gonic/gin"
)

type Service interface {
	Log(ctx context.Context, id string) ([]domain.LogModel, error)
	Logs(ctx context.Context) ([]domain.LogModel, error)
	CreateLog(ctx context.Context, log requests.LogPOST) (string, error)
}

type Router struct {
	router  *gin.RouterGroup
	service Service
}

func Register(router *gin.RouterGroup, service Service) {
	r := &Router{
		router:  router,
		service: service,
	}
	r.init()
}

func (r *Router) getLog(ctx *gin.Context) {
	id := ctx.Param("id")
	logs, err := r.service.Log(ctx, id)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, logs)
}

func (r *Router) getLogs(ctx *gin.Context) {
	logs, err := r.service.Logs(ctx)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, logs)
}

func (r *Router) createLog(ctx *gin.Context) {
	var log requests.LogPOST
	if err := ctx.ShouldBindJSON(&log); err != nil {
		ctx.JSON(400, err)
		return
	}
	id, err := r.service.CreateLog(ctx, log)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, id)
}

func (r *Router) init() {
	group := r.router.Group("/logs")
	group.GET("/:id", r.getLog)
	group.GET("/", r.getLogs)
	group.POST("/", r.createLog)
}
