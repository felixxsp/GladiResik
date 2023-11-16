package Orders

import (
	entity "GladiResik/Entity"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	Router  *gin.Engine
	Usecase *OrderUsecase
}

func OHD_Init(router *gin.Engine, usecase *OrderUsecase) *OrderHandler {
	return &OrderHandler{
		Router:  router,
		Usecase: usecase,
	}
}

func (router *OrderHandler) Standby(ctx context.Context) {
	router.Router.POST("/orders", func(ctx *gin.Context) {
		var IncomingOrder []entity.Incoming
		ctx.BindJSON(&IncomingOrder)
		ctx.IndentedJSON(http.StatusOK, IncomingOrder)
		router.Usecase.Create(ctx, IncomingOrder)
	})

	router.Router.GET("/orders", func(ctx *gin.Context) {
		item := router.Usecase.ViewAll(ctx)
		ctx.IndentedJSON(http.StatusOK, item)
	})
}
